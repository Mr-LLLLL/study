#include <iostream>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>

const int BUF_SIZE = 30;

int main(int argc, char* argv[])
{
	int sock;
	char message[BUF_SIZE];
	int str_len;
	socklen_t adr_sz;

	sockaddr_in serv_adr, from_adr;
	if (argc != 3) {
		std::cerr << "usage : " << argv[0] << " <IP> <port>" << std::endl;
		exit(1);
	}

	sock = socket(PF_INET, SOCK_DGRAM, 0);
	if (sock == -1) {
		std::cerr << "socket() error" << std::endl;
		exit(1);
	}

	memset(&serv_adr, 0, sizeof(serv_adr));
	serv_adr.sin_family = AF_INET;
	serv_adr.sin_addr.s_addr = inet_addr(argv[1]);
	serv_adr.sin_port = htons(atoi(argv[2]));

	connect(sock, (sockaddr*)&serv_adr, sizeof(serv_adr));

	while (1) {
		fputs("Insert message(q to quit): ", stdout);
		fgets(message, sizeof(message), stdin);
		if (!strcmp(message, "q\n") || !strcmp(message, "Q\n"))
			break;
		/*
		sendto(sock, message, strlen(message), 0, (sockaddr*)&serv_adr, sizeof(serv_adr));
		adr_sz = sizeof(from_adr);
		str_len = recvfrom(sock, message, BUF_SIZE, 0, (sockaddr*)&from_adr, &adr_sz);
		*/
		write(sock, message, strlen(message));
		str_len = read(sock, message, sizeof(message) - 1);
		message[str_len] = 0;
		printf("message from server: %s", message);
	}
	close(sock);
	return 0;
}

