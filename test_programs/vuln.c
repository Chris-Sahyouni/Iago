#include <stdio.h>
#include <stdlib.h>

void main(int argc, char* argv[]) {
    // needs gadgets

    if (argc != 2) {
        printf("Usage: <Buffer contents> <pointer value>");
    }

    vuln(argv[0], atoi(argv[1]));
}

void vuln(char* in, int ptr_val) {
    char* controllable_ptr;
    char buf[64];
    strcpy(buf, in);
    *controllable_ptr = ptr_val;
}