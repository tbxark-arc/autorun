#include <stdio.h>

int main(int argc, char *argv[]) {
    char buf[1024];
    FILE *file;
    size_t nread;    
    file = fopen("sub/name.text", "r");
    if (file) {
        while ((nread = fread(buf, 1, sizeof buf, file)) > 0);
        fclose(file);
    }
    printf("Hello %s", buf);
}