#ifndef __MAIN_H__
#define __MAIN_H__

#define UNIX_IPC_ADDRESS "/root/gcx/cx/cx.ipc"
#define CLI_IPC_PATH "/var/tmp/"
#define STACK_SIZE (8 * 1024)
#define BUFF_SIZE 100
#define MAX_EVENTS 10

// 当收到新的客户端请求，建立连接
int addNewConn(int listen_sock, struct epoll_event* ev, int epollfd);

// 建立新的线程处理连接上的请求，这个函数是新线程里的执行逻辑
void* start_routing (void* arg);

#endif
