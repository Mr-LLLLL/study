#include <arpa/inet.h>
#include <errno.h>
#include <fcntl.h>
#include <limits.h>
#include <netdb.h>
#include <pthread.h>
#include <semaphore.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ioctl.h>
#include <sys/msg.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <sys/utsname.h>
#include <sys/wait.h>
#include <time.h>
#include <unistd.h>

static void pr_winsize(int fd);

static void sig_winch(int signo) {
  printf("SIGWINCH received\n");
  pr_winsize(0);
}
static void pr_winsize(int fd) {}

int main(int argc, char **argv) {
  if (isatty(0) == 0)
    exit(1);

  struct winsize size;

  if (ioctl(0, TIOCGWINSZ, (char *)&size) < 0)
    fputs("TIOCGWINSZ error", stderr);

  printf("%d ros, %d columns\n", size.ws_row, size.ws_col);

  size.ws_row += 20;

  if (ioctl(0, TIOCSWINSZ, size) < 0)
    fputs("TIOCSWINSZ error", stderr);

  return 0;
}
/*
int main()
{
        char c[512];
        char *p = "hello";
        strcpy(c, p);
        int n = strlen(c);
        printf("%d\n", n);
        puts(c);
        return 0;
}
*/
