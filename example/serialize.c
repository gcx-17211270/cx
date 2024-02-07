#include <string.h>
#include <stdlib.h>
#include <stdio.h>

#include "serialize.h"

// 用来将参数序列化传递给服务器端 -- 这个失败了，想想为什么
char* Serialize(int argc, char** argv) {
    struct Info *info = (struct Info*)malloc(sizeof(struct Info));
    info->argc = argc - 1;
    struct Pack* pack = (struct Pack*)malloc(sizeof(struct Pack) * (argc - 1));
    info->argv = pack;
    for (int i = 0; i < argc - 1; i++) {
        int len = strlen(argv[i + 1]);
        char *s = (char*)malloc(sizeof(char) * len);
        memcpy(s, argv[i + 1], len);
        pack[i].len = len;
        pack[i].arg = s;
    }
    printf("%s\n", info->argv[0].arg);
}

// 用来将参数序列化传递给服务器端
// 第一个参数是程序名，所以这里argc是减一的状态，同理，反序列化时加一
char* Serialize2(int argc, char** argv) {
    int size = sizeof(int) * argc;    // 初始4个字节用来存放argc，后来每4个字节用来存放字符串长度
    for (int i = 1; i < argc; i++) {
        size += strlen(argv[i]);
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

char** Deserialize2(int* argc, char* s) {
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
    *argc += 1;
    return res;
}

int main(int argc, char** argv) {
    int i = 0, j = 0;
    Serialize(argc, argv);
}
