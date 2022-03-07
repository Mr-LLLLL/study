#include <stdio.h>
#include <fcntl.h>
#include <errno.h>
#include <string.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>
#include <sys/epoll.h>
#include <sys/time.h>
#include <pthread.h>


int main(int argc, char* argv[])
{
	int status;
	pid_t pid = fork();
	if (pid == 0) 
		return 3;
	else {
		printf("Child PID: %d \n", pid);
		pid = fork();
		if (pid == 0)
			exit(7);
		else {
			sleep(1);
			printf("Child PID: %d \n", pid);
			wait(&status);
			if (WIFEXITED(status))
				printf("Child send one: %d \n", WEXITSTATUS(status));
			 
			wait(&status);
			if (WIFEXITED(status))
				printf("Child send one: %d \n", WEXITSTATUS(status));
			sleep(30);
		}
	}

	return 0;
}
