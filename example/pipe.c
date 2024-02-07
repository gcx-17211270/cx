#include <unistd.h>
#include <stdio.h>
#include <sys/types.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <string.h>

int main(int argc, char** argv){
    int pipeFd[2];
    if (pipe(pipeFd) == -1) {
        perror("pipe");
        exit(EXIT_FAILURE);
    }

    char buf[100];
    pid_t pid = fork();
    if (pid == -1) {
        perror("fork");
        exit(EXIT_FAILURE);
    } else if (pid == 0) {
        close(pipeFd[0]);
        sprintf(buf, "fork subprocess success, pid = %d\n\0", getpid());
        printf("child:%s sizeof(buf)=%d\n", buf, strlen(buf));
        write(pipeFd[1], buf, strlen(buf));
        sleep(200);
    } else {
        close(pipeFd[1]);
        int r = read(pipeFd[0], buf, 100);
        while (r != EOF && r != 0) {
            printf("parent: %d-%s\n", r, buf);
            r = read(pipeFd[0], buf, 100);
        }
    }
    
    return 0;
}

