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
#include "clone.h"
#include "main.h"
#include "serialize.h"

int doInit(void* arg){
  char *envp[] = {NULL};
  struct Info* info = (struct Info*)arg;  
  // 多的2个参数，一个是第一个的应用程序名称，一个是最后的NULL
  char **args = (char**)malloc(sizeof(char*) * (info->argc + 2));
  args[0] = RUNCX_EXE;
  for (int i = 0; i < info->argc; i++) {
      struct Pack *pack = info->argv + i;
      args[i+1] = (char*)malloc(pack->len);
      memcpy(args[i+1], pack->arg, pack->len);
     printf("doInit: %d %s %p\n", pack->len, pack->arg, pack);
      free(pack->arg);
  }
  args[info->argc + 2 - 1] = NULL;
  printf("hello doInit");
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
    printf("%d - %s\n", argc, argv[0]);
   struct Info* info = ArgSerialize(argc, argv);
    printf("before: %p %d\n", info, *(int*)info);
    printf("before clone : %d\n", info->argc);
    int clone_status = clone(doInit, child_stack_down + STACK_SIZE,
        CLONE_NEWCGROUP|CLONE_NEWNET|CLONE_NEWIPC
        |CLONE_NEWNS|CLONE_NEWPID|CLONE_NEWUSER|CLONE_NEWUTS
        |SIGCHLD
        , info);

    if (clone_status == -1) {
        perror("clone subprocess error occur");
        return 1;
    }
    printf("clone subprocess success, pid = %d\n", clone_status);
    return 0;
}


