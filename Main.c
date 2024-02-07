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

#include "main.h"
#include "serialize.h"
#include "clone.h"
#include "close.h"

int* castInt2Point(int x) {
    int *y = (int*)malloc(sizeof(int));
    *y = x;
    return y;
}

int main(int argc, char** argv){
    printf("%d\n", getpid());
    signal(SIGINT, closeHandler);
    
    // 从默认的unix域，接受到创建一个子进程的信息，然后开始clone
    int fd, size;
    struct sockaddr_un un;
    un.sun_family = AF_UNIX;
    strcpy(un.sun_path, UNIX_IPC_ADDRESS);
    // 确定绑定地址长度的方法是，先计算sun_path成员在sockaddr_un结构中的偏移量，
    // 然后将结构与路径名长度相加（因为在sockaddr_un结构中与sun_path之前的成员与实现相关）
    if ((fd = socket(AF_UNIX, SOCK_STREAM, 0)) == -1) {
        perror("socket failed");
        exit(1);
    };
    closeClear((void*)castInt2Point(fd), closeFd);
    closeClear(NULL, unlinkFd);
    
    size = offsetof(struct sockaddr_un, sun_path) + strlen(un.sun_path);
    if(bind(fd, (struct sockaddr*)&un, size) < 0) {
        perror("bind failed");
        exit(1);
    }
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
                printf("new conn established\n");
            } else {
                pthread_t tid;
                int* iofd = (int*)malloc(sizeof(int));
                *iofd = events[n].data.fd;
                if (pthread_create(&tid, NULL, start_routing, iofd) != 0) {
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
    printf("add new conn\n");
    return 0;
}

void* start_routing (void* arg) {
    char buf[BUFF_SIZE];
    int len = 0, r = 0;
    int* fd = (int*)arg;
    char* content = (char*)malloc(len);
    while ((r = read(*fd, buf, BUFF_SIZE)) > 0) {
        content = (char*)realloc(content, len + r);
        memcpy(content + len, buf, r);
        len += r;
    };
    printf("read %d\n", len);
    int *argc = (int*)malloc(sizeof(int));
    char** argv = Deserialize1(argc, content);
    printf("有新创建进程请求\n");
    for (int i = 0; i < *argc; i++) {
        printf("%d.%s\n", i+1, argv[i]);
    }
    printf("\n");
    cloneSubProcess(*argc, argv);
    
    free(arg);
    free(content);
    for (int i = 0; i < *argc; i++) {
        free(argv[i]);
    }
    free(argv);
    free(argc);
}

