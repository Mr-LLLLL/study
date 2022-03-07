#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <sys/time.h>
#include <sys/select.h>

const int BUF_SIZE = 30;
void error_handling(char *buf);

int main(int argc, char *argv[])
{
	int serv_sock, clnt_sock;
	struct sockaddr_in serv_adr, clnt_adr;
	struct timeval timeout;
	fd_set reads, cpy_reads;

	socklen_t adr_sz;

	int fd_max, str_len, fd_num;
	char buf[BUF_SIZE];
	if (argc != 2) {
		printf("usage : %s <port>\n", argv[0]);
		exit(1);
	}

	serv_sock = socket(PF_INET, SOCK_STREAM, 0);
	memset(&serv_adr, 0, sizeof(serv_adr));
	serv_adr.sin_family = AF_INET;
	serv_adr.sin_addr.s_addr = htonl(INADDR_ANY);
	serv_adr.sin_port = htons(atoi(argv[1]));

	if (bind(serv_sock, (struct sockaddr*)&serv_adr, sizeof(serv_adr)) == -1)
		error_handling("bind() error");
	if (listen(serv_sock, 5) == -1)
		error_handling("listen() error");

	FD_ZERO(&reads);
	FD_SET(serv_sock, &reads);
	fd_max = serv_sock;

	while (1) {
		cpy_reads = reads;
		timeout.tv_sec = 0;
		timeout.tv_usec = 3000000;

		if ((fd_num = select(fd_max + 1, &cpy_reads, 0, 0, &timeout)) == -1)
			break;
		if (fd_num == 0)
			printf("ksdjfk\n");

		int i;
		for (i = 0; i < fd_max + 1; ++i) {
			if (FD_ISSET(i, &cpy_reads)) {
				if (i == serv_sock) {
					clnt_sock = accept(serv_sock, (struct sockaddr*)&clnt_adr, &adr_sz);
					FD_SET(clnt_sock, &reads);
					if (clnt_sock > fd_max)
						fd_max = clnt_sock;
					printf("connect client: %d \n", clnt_sock);
				} else {
					str_len = read(i, buf, BUF_SIZE);
					if (str_len == 0) {
						FD_CLR(i, &reads);
						close(i);
						printf("closed client: %d \n", i);
					} else 
						write(i, buf, str_len);
				}
			}
		}
	}
	close(serv_sock);
	return 0;
}

void error_handling(char *buf) {
	fputs(buf, stderr);
	fputc('\n', stderr);
	exit(1);
}



