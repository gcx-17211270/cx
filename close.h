#ifndef __CLOSE_H__
#define __CLOSE_H__

// 接受到关闭信号时，执行清理资源的动作
void closeHandler(int);

// 定义的清理资源的函数
struct CloseFunction {
    int(*Function)(void*);
    void* arg;
    struct CloseFunction* pre;
    struct CloseFunction* next;
};

// 两种清理动作
int closeFd(void*);
int unlinkFd(void*);

// 调用这个函数时，执行所有的清理资源函数
void closeClear(void* arg, int(*f)(void*));

#endif

