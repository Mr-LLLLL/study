#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>

const int BUF_SIZE = 1024;

int main(int argc, char* argv[])
{
	int sock;
	char message[BUF_SIZE];
	int str_len, recv_len, recv_cnt;
	sockaddr_in serv_adr;
	
	if (argc != 3) {
		std::cout << "Usage : " << argv[0] << " <IP> <port>" << std::endl;
		exit(1);
	}

	sock = socket(PF_INET, SOCK_STREAM, 0);
	if (sock == -1) {
		std::cerr << "socket() error" << std::endl;
		exit(1);
	}

	memset(&serv_adr, 0, sizeof(serv_adr));
	serv_adr.sin_family = AF_INET;
	serv_adr.sin_addr.s_addr = inet_addr(argv[1]);
	serv_adr.sin_port = htons(atoi(argv[2]));

	if (connect(sock, (sockaddr*)&serv_adr, sizeof(serv_adr)) == -1) {
		std::cerr << "connect() error!" << std::endl;
		exit(1);
	} else
		std::cout << "Connected........" << std::endl;


	while (1) {
		std::cout << "Input message(Q to quit): " << std::endl;
		fgets(message, BUF_SIZE, stdin);

		if (!strcmp(message, "q\n") || !strcmp(message, "Q\n"))
			break;

		str_len = write(sock, message, strlen(message));
		recv_len = 0;
		while (recv_len < str_len) {
			recv_cnt = read(sock, message, BUF_SIZE - 1);
			if (recv_cnt == -1) {
				std::cerr << "read() error" << std::endl;
				exit(1);
			}
			recv_len += recv_cnt;
		}
		message[recv_len] = 0;
		std::cout << "Message from server: " << message << std::endl;
	}
	close(sock);
	return 0;
}
