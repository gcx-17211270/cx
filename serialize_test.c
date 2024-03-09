//
// Created by 32353_ixl9bt0 on 2024/3/9.
//
#include "serialize.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
    char* buf = (char*)malloc(sizeof(char) * 100);
    char* cur = buf;
    int len = 0;
    char* s[2];
    s[0] = "Hello";
    s[1] = "World";
    for (int i = 0; i < 2; i++) {
        int l = strlen(s[i]);
        memcpy(cur, s[i], l);
        cur += l;
        *cur = '\n';
        cur++;
        len += l + 1;
    }
    int *argc = (int*)malloc(sizeof(int));
    char **argv = Deserialize2(buf, len, argc);
    if (argv == NULL) {
        perror("Deserialize2 error");
        exit(1);
    }
    for (int i = 0; i < *argc; i++) {
        printf("%s\n", argv[i]);
    }
    return 0;
}