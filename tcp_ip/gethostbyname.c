#include <stdio.h>
#include <iostream> 
#include <stdlib.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <netdb.h>

int main(int argc, char* argv[])
{
	hostent *host;
	if (argc != 2) {
		std::cerr << "Usage : " << argv[0] << " <addr>" << std::endl;
		exit(1);
	}

	host = gethostbyname(argv[1]);
	if (!host) {
		std::cerr << "gethost... error" << std::endl;
		exit(1);
	}

	std::cout << "Official name: " << host->h_name << std::endl;
	for (int i = 0; host->h_aliases[i]; ++i)
		std::cout << "Aliases " << i + 1 <<  ": " << host->h_aliases[i] << std::endl;
	std::cout << "Address tyep: " << ((host->h_addrtype == AF_INET) ? "AF_INET" : "AF_INET6") << std::endl;
	for (int i = 0; host->h_addr_list[i]; ++i)
		std::cout << "IP addr " << i + 1 << ": " << inet_ntoa(*(in_addr*)host->h_addr_list[i]) << std::endl;
	return 0;

}

