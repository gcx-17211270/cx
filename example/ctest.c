#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <syslog.h>

#define sleep "/usr/bin/sleep"

int main(int argc, char** argv) {
    // use of syslog
    openlog(__func__, LOG_CONS| LOG_ODELAY|LOG_NOWAIT|LOG_PERROR|LOG_PID,
        LOG_AUTH|LOG_MAIL);
    syslog(LOG_INFO, "hello world");
    closelog();


    // test for execve
    char** s = (char**)malloc(sizeof(char*) * 3);
    s[0] = sleep;
    char* s1 = malloc(sizeof(char) * 1);
    s1[0] = '3';
    // s1[1] = '0'; 
    s[1] = s1;
    
    
    s[2] = NULL;
    char *envp[] = {NULL};
    execve(sleep, s, envp);
}
