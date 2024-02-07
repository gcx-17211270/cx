#ifndef __CLONE_H__
#define __CLONE_H__

int cloneSubProcess(int argc, char**argv);

int doInit(void* arg);

typedef struct ResourceLimit {
    int cpu;        // 单位m 1000m代表1核cpu的算力
    int memory;     // 单位MB
    int blkio;      // 对io的速度限制，不清楚该怎么处理呢
} RL, *rl;

typedef struct SubProcessRequest {
    const char* exe;
    struct ResourceLimit rl;
} SPR, *spr;

#endif
