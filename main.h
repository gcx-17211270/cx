#ifndef __MAIN_H__
#define __MAIN_H__

#include <sys/epoll.h>

#ifdef __UNIX_IPC__
#define UNIX_IPC_ADDRESS "/root/code/cx/cx.ipc"
#else
#define SOCKET_ADDRESS 8080
#endif
#define CLI_IPC_PATH "/var/tmp/"
#define STACK_SIZE (1024 * 1024)
#define BUFF_SIZE 100
#define MAX_EVENTS 10

// 当收到新的客户端请求，建立连接
int addNewConn(int listen_sock, struct epoll_event* ev, int epollfd);

// 建立新的线程处理连接上的请求，这个函数是新线程里的执行逻辑
void* start_routing (void* arg);

// 用来解析请求、创建子进程的进程
#define RUNCX_EXE "./runcx.exe"

#endif
