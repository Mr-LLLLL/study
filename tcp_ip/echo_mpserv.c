#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <signal.h>
#include <sys/wait.h>
#include <arpa/inet.h>
#include <sys/socket.h>

const int BUF_SIZE = 30;

void read_childproc(int sig);

int main(int argc, char *argv[])
{
	int serv_sock, clnt_sock;
	struct sockaddr_in serv_adr, clnt_adr;

	int fds[2];

	pid_t pid;
	struct sigaction act;
	socklen_t adr_sz;
	int str_len;
	char buf[BUF_SIZE];
	
	if (argc != 2) {
		std::cout << "usage : " << argv[0] << " <port>" << std::endl;
		exit(1);
	}

	act.sa_handler = read_childproc;
	sigemptyset(&act.sa_mask);
	act.sa_flags = 0;

	sigaction(SIGCHLD, &act, 0);

	serv_sock = socket(PF_INET, SOCK_STREAM, 0);
	
	memset(&serv_adr, 0, sizeof(serv_adr));
	serv_adr.sin_family = AF_INET;
	serv_adr.sin_addr.s_addr = htonl(INADDR_ANY);
	serv_adr.sin_port = htons(atoi(argv[1]));

	if (bind(serv_sock, (sockaddr*)&serv_adr, sizeof(serv_adr)) == -1) {
		std::cerr << "bind() error" << std::endl;
		exit(1);
	}

	if (listen(serv_sock, 5) == -1) {
		std::cerr << "listen() error" << std::endl;
		exit(1);
	}

	pipe(fds);
	pid = fork();
	if (pid == 0) {
		FILE *fp = fopen("echomsg.txt", "wt");
		char msgbuf[BUF_SIZE];
		for (int i = 0; i < 10; ++i) {
			int len = read(fds[0], msgbuf, BUF_SIZE);
			fwrite((void*)msgbuf, 1, len, fp);
		}
		fclose(fp);
	}
	while (1) {
		adr_sz = sizeof(clnt_adr);
		clnt_sock = accept(serv_sock, (sockaddr*)&clnt_adr, &adr_sz);
		if (clnt_sock == -1)
			continue;
		else
			std::cout << "new client connected..." << std::endl;

		pid = fork();
		if (pid == -1) {
			close(clnt_sock);
		} else if (pid == 0) {
			close(serv_sock);
			while ((str_len = read(clnt_sock, buf, BUF_SIZE)) != 0) {
				write(clnt_sock, buf, str_len);
				write(fds[1], buf, str_len);
			}

			close(clnt_sock);
			std::cout << "client disconnected..." << std::endl;
			return 0;
		} else
			close(clnt_sock);
	}
	close(serv_sock);
	return 0;
}

void read_childproc(int sig) {
	pid_t pid;
	int status;
	pid = waitpid(-1, &status, WNOHANG);
	std::cout << "removed proc id: " << pid << std::endl;
}




