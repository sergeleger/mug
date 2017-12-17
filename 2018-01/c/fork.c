#include <stdio.h>
#include <unistd.h>

void main() {
    pid_t pid = -1;

    printf("Just one process so far\n");
    printf("Calling fork ...\n");

    pid = fork();

    if(pid == 0 ) {
        printf("I'm the child!\n");
    } else if (pid > 0) {
        printf("I'm the parent, child pid: %d\n", pid);
    }
}
