#include "serialize.h"
#include <string.h>
#include <stdlib.h>
#include <stdio.h>

// 用来将参数序列化传递给服务器端
struct Info* ArgSerialize(int argc, char** argv) {
    struct Info *info = (struct Info*)malloc(sizeof(struct Info));
    info->argc = argc;
    struct Pack* pack = (struct Pack*)malloc(sizeof(struct Pack) * argc);
    info->argv = pack;
    int size = sizeof(struct Info) + argc * sizeof(struct Pack);
    for (int i = 0; i < argc; i++) {
        int len = strlen(argv[i]);
        size += len;
        char *s = (char*)malloc(sizeof(char) * len);
        memcpy(s, argv[i], len);
        pack[i].len = len;
        pack[i].arg = s;
    }
    info->size = size;
    return info;
}

void printInfo(struct Info* info) {
    printf("\targc:%d len:%d argv=(", info->argc, info->size);
    for (int i = 0; i < info->argc; i++) {
        printf("%s ", info->argv[i].arg);
    }
    printf(")\n");
}

// 用来将参数序列化传递给服务器端
// 第一个参数是程序名，所以这里argc是减一的状态，同理，反序列化时加一
char* Serialize1(int argc, char** argv, int* size_r) {
    int size = sizeof(int) * argc;    // 初始4个字节用来存放argc，后来每4个字节用来存放字符串长度
    for (int i = 1; i < argc; i++) {
        size += strlen(argv[i]);
    }
    if (size_r != NULL) {
        *size_r = size;
    }
    char *s, *tmp;
    s = (char*)malloc(sizeof(char) * size);
    memset(s, 0, size);
    int* n = (int*)malloc(sizeof(int));
    *n = argc - 1;
    memcpy(s, n, sizeof(int));
    tmp = s + sizeof(int); 
    for (int i = 1; i < argc; i++) {
        *n = strlen(argv[i]);
        memcpy(tmp, n, sizeof(int));
        tmp += sizeof(int);
        memcpy(tmp, argv[i], *n);
        tmp += *n;
    }
    free(n);
    return s;
};

// 这个的序列化是按 四字节长度 然后一个字符串 这样的方式组织的
char** Deserialize1(int* argc, char* s) {
    *argc = *(int*)s;
    char** res = (char**)malloc(sizeof(char*) * (*argc));
    char *tmp = s + sizeof(int);
    for (int i = 0; i < *argc; i++) {
        int *len = (int*)tmp;
        char* str = (char*)malloc(sizeof(char) * (*len));
        tmp += sizeof(int);
        memcpy(str, tmp, *len);
        tmp += *len;
        res[i] = str;
    }
    return res;
}

// 这个的序列化是按 \n分割，然后len是返回的行数
char** Deserialize2(char* s, int len, int* argc) {
    char** res = NULL;
    *argc = 0;
    int start = 0;
    for (int i = 1; i < len; i++) {
        if (s[i] == '\n') {
            res = (char**)realloc(res, ((*argc) + 1) * sizeof(char*));
            if (res == NULL) {
                perror("realloc error occur");
                return NULL;
            }
            if ((i - start) > 0) {
                res[*argc] = (char*)malloc(sizeof(char) * (i - start));
                memcpy(res[*argc], s+start, (i - start));
                (*argc)++;
                start = i + 1;
            }
        }
    }
    if (s[len - 1] != '\n') {
        res = (char**)realloc(res, ((*argc) + 1) * sizeof(char*));
        if (res == NULL) {
            perror("realloc error occur");
            return NULL;
        }
        res[*argc] = (char*)malloc(sizeof(char) * (len - start));
        memcpy(res[*argc], s+start, (len - start));
        (*argc)++;
    }
    return res;
}

