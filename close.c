#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "close.h"
#include "main.h"

static struct CloseFunction cf = {.next = NULL};

int closeFd(void* arg) {
    int* fd = (int*)arg;
    printf("close fd :%d\n", *fd);
    if (close(*fd) == -1) {
        perror("close");
        return 1;
    }
    free(fd);
    return 0;
}

int unlinkFd(void* arg) {
    printf("%d unlink address\n", getpid());
#ifdef __UNIX_IPC__
    if (unlink(UNIX_IPC_ADDRESS) == -1) {
        perror("unlink error");
        return 1;
    };
#endif
    return 0;
}

void closeClear(void* arg, int(*f)(void*)) {
    struct CloseFunction* tm = (struct CloseFunction*)malloc(sizeof(struct CloseFunction));
    struct CloseFunction *temp = &cf;
    while (temp->next != NULL) {
        temp = temp->next;
    }
    temp->next = tm;
    tm->pre = temp;
    tm->arg = arg;
    tm->Function = f;
    tm->next = NULL;
}

void closeHandler(int signum) {
    int failFlag;
    struct CloseFunction* temp = cf.next;
    while (temp != NULL) {
        int(*f)(void*) = temp->Function;
        if (f(temp->arg) != 0){
            failFlag = 1;
        };
        temp = temp->next;
    }
    exit(failFlag);
}

