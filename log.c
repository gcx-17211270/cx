//
// Created by 32353_ixl9bt0 on 2024/3/10.
//

#include "log.h"
#include <stdio.h>
#include <time.h>

void logtime() {
    // 获取当前时间
    time_t current_time = time(NULL);

    // 转换时间为本地时间
    struct tm* local_time = localtime(&current_time);

    // 格式化时间字符串
    char time_str[50];
    strftime(time_str, sizeof(time_str), "%Y-%m-%d %H:%M:%S", local_time);

    // 打印日志消息和时间
    printf("[%s] ", time_str);
}