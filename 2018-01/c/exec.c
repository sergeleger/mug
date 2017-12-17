#include <stdio.h>
#include <unistd.h>

void main() {
    int result;
    
    printf("Executing ls\n");

    result = execlp("ls", "ls", "-l", (char *)0);

    printf("This is never seen unless execlp failed: %d\n", result);
}
