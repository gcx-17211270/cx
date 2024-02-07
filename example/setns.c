#define _GNU_SOURCE

#define errExit(msg) do {perror(msg); exit(EXIT_FAILURE);\
                    } while (0);

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>


int main(int argc, char** argv){
    int fd;
    
    if (argc < 3) {
        fprintf(stderr, "%s /proc/PID/ns/FILE cmd args...\n", argv[0]);
        exit(EXIT_FAILURE);
    }
    
    fd = open(argv[1], O_RDONLY);
    if (fd == -1)
        errExit("open");
    if (setns(fd, 0) == -1)
        errExit("setns");
    execvp(argv[2], &argv[2]);
    errExit("execvp");         
    
    return 0;
}

