#include <stdio.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <stdlib.h>
#include <stddef.h>
#include <unistd.h>
#include <string.h>
#include "main.h"
#include "serialize.h"

int main(int argc, char** argv){    
    struct sockaddr_un un, sun;
    int fd = socket(AF_UNIX, SOCK_STREAM, 0);
    if (fd < 0) {
        perror("socket error");
        exit(EXIT_FAILURE);
    }
    memset(&un, 0, sizeof(un));
    un.sun_family = AF_UNIX;
    sprintf(un.sun_path, "%s%05ld", CLI_IPC_PATH, (long)getpid());
    int size = offsetof(struct sockaddr_un, sun_path) + strlen(un.sun_path);
    if(bind(fd, (struct sockaddr*)&un, size) < 0) {
        perror("bind failed");
        exit(1);
    }
    
    memset(&sun, 0, sizeof(sun));
    sun.sun_family = AF_UNIX;
    strcpy(sun.sun_path, UNIX_IPC_ADDRESS);
    int len = offsetof(struct sockaddr_un, sun_path) + strlen(sun.sun_path);
    if(connect(fd, (struct sockaddr*)&sun, len) < 0) {
        perror("connect failed");
        exit(1);
    }
    
    int *sz = (int*)malloc(sizeof(int));
    char* s = Serialize1(argc, argv, sz);
    printf("size = %d\n", *sz);
    write(fd, s, *sz);
    close(fd);
    free(sz);
    free(s);
    unlink(un.sun_path);
    return 0;
}

// 用完连接后，删除/var/tmp/pid下的临时文件
