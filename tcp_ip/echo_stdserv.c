#include <iostream>
#include <cstdio>
#include <cstdlib>
#include <cstring>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>

const int BUF_SIZE = 1024;

int main(int argc, char* argv[])
{
	int serv_sock, clnt_sock;
	char message[BUF_SIZE];
	int str_len;

	sockaddr_in serv_addr, clnt_addr;
	socklen_t clnt_addr_sz;

	FILE* readfp;
	FILE* writefp;

	if (argc != 2) {
		std::cout << "Usage : " << argv[0] << " <port>" << std::endl;
		exit(1);
	}

	serv_sock = socket(PF_INET, SOCK_STREAM, 0);
	if (serv_sock == -1) {
		std::cerr << "socket() error" << std::endl;	
		exit(1);
	}

	memset(&serv_addr, 0, sizeof(serv_addr));
	serv_addr.sin_family = AF_INET;
	serv_addr.sin_addr.s_addr = htonl(INADDR_ANY);
	serv_addr.sin_port = htons(atoi(argv[1]));

	if (bind(serv_sock, (sockaddr*)&serv_addr, sizeof(serv_addr)) == -1) {
		std::cerr << "bind() error" << std::endl;
		exit(1);
	}

	if (listen(serv_sock, 5) == -1) {
		std::cerr << "listen() error";
		exit(1);
	}
	clnt_addr_sz = sizeof(clnt_addr);

	for (int i = 0; i < 5; ++i) {
		clnt_sock = accept(serv_sock, (sockaddr*)&clnt_addr, &clnt_addr_sz);
		if (clnt_sock == -1) {
			std::cerr << "accept() error" << std::endl;
			exit(1);
		} else
			std::cout << "Connected client " << i << std::endl;

		readfp = fdopen(clnt_sock, "r");
		writefp = fdopen(dup(clnt_sock), "w");

		while (!feof(readfp)) {
			fgets(message, BUF_SIZE, readfp);
			fputs(message, writefp);
			fflush(writefp);
		}
		shutdown(fileno(writefp), SHUT_WR);
		fclose(writefp);
	}
	fclose(readfp);
	return 0;
}


