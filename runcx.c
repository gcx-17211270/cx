#include <stdio.h>

int main(int argc, char** argv){
    printf("Hello world\n");
    for (int i = 0; i < argc; i++) {
        printf("%d.%s\n", i, argv[i]);
    }
    return 0;
}

