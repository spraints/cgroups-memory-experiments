#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char **argv) {
  if (argc != 3) {
    printf("Usage: ./child CHILDNUM BYTESPERSECOND\n");
    return 1;
  }
  char *childnum = argv[1];
  int bytesPerSecond = atoi(argv[2]);
  long bytesAllocated = 0;
  void *last = NULL;
  for (;;) {
    void *p = malloc(bytesPerSecond);
    if (!p) {
      printf("child %s: could not allocate\n", childnum);
      return 2;
    }
    if (last) {
      memcpy(p, last, bytesPerSecond);
    }
    last = p;

    bytesAllocated += bytesPerSecond;
    printf("child %s: %ld MB allocated\n", childnum, bytesAllocated / 1024 / 1024);

    sleep(1);
  }
  return 0;
}
