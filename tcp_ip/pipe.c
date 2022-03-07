#include <stdio.h>
#include <unistd.h>

const int BUF_SIZE = 49;

int main(int argc, char* argv[])
{
	int fds1[2], fds2[2];
	char str1[] = "who are you?";
	char str2[] = "thank you for your massage";
	char buf[BUF_SIZE];
	pid_t pid;
	
	pipe(fds1);
	pipe(fds2);
	pid = fork();
	if (pid == 0) {
		write(fds1[1], str1, sizeof(str1));
		read(fds2[0], buf, BUF_SIZE);
		printf("Child proc output: %s \n", buf);
	} else {
		read(fds1[0], buf, BUF_SIZE);
		printf("Parent proc output: %s \n", buf);
		write(fds2[1], str2, sizeof(str2));
		sleep(1);
	}
	return 0;
}
