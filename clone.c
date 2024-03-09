/**
创建子进程，执行指定命令
*/
#define _GNU_SOURCE
#include <sched.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <signal.h>
#include <stdlib.h>
#include <sys/utsname.h>
#include "clone.h"
#include "main.h"
#include "serialize.h"
#include "log.h"

#define errExit(msg)    do { perror(msg); exit(EXIT_FAILURE); \
                               } while (0)

static int              /* Start function for cloned child */
childFunc(void *arg)
{
    struct utsname uts;

    /* Change hostname in UTS namespace of child */

    if (sethostname(arg, strlen(arg)) == -1)
        errExit("sethostname");

    /* Retrieve and display hostname */

    if (uname(&uts) == -1)
        errExit("uname");
    printf("uts.nodename in child:  %s\n", uts.nodename);

    /* Keep the namespace open for a while, by sleeping.
       This allows some experimentation--for example, another
       process might join the namespace. */

    sleep(200);

    return 0;           /* Child terminates now */
}

int doInit(void* arg){
    logtime();
    printf("hello doInit");
    char *envp[] = {NULL};
    struct Info* info = (struct Info*)arg;
    // 多的2个参数，一个是第一个的应用程序名称，一个是最后的NULL
    char **args = (char**)malloc(sizeof(char*) * (info->argc + 2));
    args[0] = RUNCX_EXE;
    for (int i = 0; i < info->argc; i++) {
        struct Pack *pack = info->argv + i;
        args[i+1] = (char*)malloc(pack->len);
        memcpy(args[i+1], pack->arg, pack->len);
        logtime();
        printf("doInit: %d %s %p\n", pack->len, pack->arg, pack);
        free(pack->arg);
    }
    args[info->argc + 2 - 1] = NULL;
    free(info->argv);
    free(info);
    execve(RUNCX_EXE, args, envp);
    return 0;
}

//我需要的是一个类似K8s的kubelet的工具，它能够解析yaml文件，创建对应的子进程
// 1.同一个container下的进程，在相同的ns下
// 2.在这个进程下，进行cgroup的限制
int cloneSubProcess(int argc, char** argv) {
    void* child_stack_down = malloc(STACK_SIZE);
    if (child_stack_down == NULL) {
        perror("malloc error occur");
        exit(1);
    }
    struct Info* info = ArgSerialize(argc, argv);
    logtime();
    printf("cloneSubProcess: build info:");
    printInfo(info);
    int clone_status = clone(childFunc, child_stack_down + STACK_SIZE,
        // CLONE_NEWCGROUP|         // 我的centos7.6上没有这个
                             CLONE_NEWUTS
        , argv[1]);

    if (clone_status == -1) {
        logtime();
        perror("cloneSubProcess: clone subprocess error occur");
        return 1;
    }
    logtime();
    printf("clone subprocess success, pid = %d\n", clone_status);
    sleep(10);
    return 0;
}


