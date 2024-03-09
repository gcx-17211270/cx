#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <stddef.h>
#include <sys/epoll.h>
#include <fcntl.h>
#include <signal.h>
#include <pthread.h>
#include <string.h>
#include <netinet/in.h>

#include "main.h"
#include "serialize.h"
#include "clone.h"
#include "close.h"
#include "log.h"

int* castInt2Point(int x) {
    int *y = (int*)malloc(sizeof(int));
    *y = x;
    return y;
}

int main(int argc, char** argv){
    logtime();
    printf("main: now pid is %d\n", getpid());
    signal(SIGINT, closeHandler);
    
    // 从默认的unix域，接受到创建一个子进程的信息，然后开始clone
    int fd, size;
    struct sockaddr* addr;
    int family;
#ifdef __UNIX_IPC__
    struct sockaddr_un un;
    un.sun_family = AF_UNIX;
    strcpy(un.sun_path, UNIX_IPC_ADDRESS);
    addr = (struct sockaddr*)&un;
    size = offsetof(struct sockaddr_un, sun_path) + strlen(un.sun_path);
    closeClear(NULL, unlinkFd);
    family = AF_UNIX;
#else
    struct sockaddr_in saddr;
    bzero((void*)&saddr, sizeof(saddr));
    saddr.sin_port = htons(8080);
    saddr.sin_addr.s_addr = INADDR_ANY;
    saddr.sin_family = PF_INET;
    addr = (struct sockaddr*)&saddr;
    size = sizeof(saddr);
    family = AF_INET;
#endif
    // 确定绑定地址长度的方法是，先计算sun_path成员在sockaddr_un结构中的偏移量，
    // 然后将结构与路径名长度相加（因为在sockaddr_un结构中与sun_path之前的成员与实现相关）
    if ((fd = socket(family, SOCK_STREAM, 0)) == -1) {
        perror("socket failed");
        exit(1);
    };
    closeClear((void*)castInt2Point(fd), closeFd);
    if(bind(fd, addr, size) < 0) {
        perror("bind failed");
        exit(1);
    }
    logtime();
    printf("Unix domain socket bound\n");
    if (listen(fd, 10) < 0) {
        perror("listen failed");
        exit(1);
    }
    
    // 使用epoll处理连接的建立的cs之间的io
    struct epoll_event ev, events[MAX_EVENTS];
    int epollfd, nfds;
    epollfd = epoll_create1(0);
    closeClear((void*)castInt2Point(epollfd), closeFd);
    // 连接请求的处理
    ev.events = EPOLLIN;
    ev.data.fd = fd;
    if (epoll_ctl(epollfd, EPOLL_CTL_ADD, fd, &ev) == -1) {
        perror("add accept epoll event failed");
        exit(1);
    }
    for (;;) {
        nfds = epoll_wait(epollfd, events, MAX_EVENTS, -1);
        if (nfds == -1) {
            perror("epoll_wait");
            exit(EXIT_FAILURE);
        }
        for (int n = 0; n < nfds; n++) {
            if (events[n].data.fd == fd) {
                if (addNewConn(fd, &ev, epollfd) == -1) {
                    perror("add new conn error");
                    exit(EXIT_FAILURE);
                }
            } else {
                pthread_t tid;
                int* iofd = (int*)malloc(sizeof(int));
                *iofd = events[n].data.fd;
                if (pthread_create(&tid, NULL, start_routing, iofd) != 0) {
                    logtime();
                    fprintf(stderr, "pthread_create error. %s", tid);
                }
            }
        }
    }
    
    return 0;
}

int addNewConn(int listen_sock, struct epoll_event* ev, int epollfd) {
    struct sockaddr addr;
    int addrlen = sizeof(struct sockaddr);
    int conn_sock = accept(listen_sock,
            (struct sockaddr *) &addr, &addrlen);
    if (conn_sock == -1) {
        perror("accept");
        return -1;
    }
    ev->events = EPOLLET;
    ev->data.fd = conn_sock;
    if (epoll_ctl(epollfd, EPOLL_CTL_ADD, conn_sock, ev) == -1) {
        perror("epoll_ctl: conn_sock");
        return -1;
    }
    logtime();
    printf("main:addNewConn:add new conn\n");
    return 0;
}

void* start_routing (void* arg) {
    char buf[BUFF_SIZE];
    int len = 0, r = 0;
    int* fd = (int*)arg;
    char* content = (char*)malloc(len * sizeof(char));
    while ((r = read(*fd, buf, BUFF_SIZE)) > 0) {
        content = (char*)realloc(content, (len + r + 1) * sizeof(char));
        memcpy(content + len, buf, r);
        len += r;
        content[len] = '\0';
    };
    logtime();
    printf("pthread conn:start_routing: read %d %s\n", len, buf);
    int *argc = (int*)malloc(sizeof(int));
    char** argv = Deserialize1(argc, content);
    if (argv == NULL) {
        goto free;
    }
    cloneSubProcess(*argc, argv);
    
free:    free(arg);
    free(content);
    if (argv != NULL) {
        for (int i = 0; i < *argc; i++) {
            free(argv[i]);
        }
        free(argv);
    }
    free(argc);
}

