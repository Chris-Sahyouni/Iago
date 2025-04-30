#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void main(int argc, char* argv[]) {

    if (argc != 2) {
        printf("Usage: <Buffer contents> <pointer value>");
    }

    vuln(argv[1], atoi(argv[2]));
}

void vuln(char* in, int ptr_val) {
    char* controllable_ptr;
    char buf[64];
    strcpy(buf, in);
    *controllable_ptr = ptr_val;
}