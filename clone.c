#define _GNU_SOURCE
#include <sched.h>
#include <stdio.h>
#include <signal.h>
#include <unistd.h>
#include <stdlib.h>
#include "clone.h"
#include "main.h"


int doInit(void* arg){
  int* fd = (int*)arg; 
  char *args[] = {"/usr/bin/sleep", "300", NULL}; 
  char *envp[] = {NULL};
  execve("/usr/bin/sleep", args, envp);
  return 0;
}

int cloneSubProcess(int argc, char** argv) {
    void* child_stack_down = malloc(STACK_SIZE);
    if (child_stack_down == NULL) {
        perror("malloc error occur");
        exit(1);
    }
    int (*fn)(void*) = doInit;
    int clone_status = clone(fn, child_stack_down + STACK_SIZE,
        CLONE_NEWCGROUP|CLONE_NEWNET|CLONE_NEWIPC
        |CLONE_NEWNS|CLONE_NEWPID|CLONE_NEWUSER|CLONE_NEWUTS
        |SIGCHLD
        , NULL);

    if (clone_status == -1) {
        perror("clone subprocess error occur");
        exit(1);
    }
    printf("clone subprocess success, pid = %d\n", clone_status);
    return 0;
}


