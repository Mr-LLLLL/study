#include <iostream>
#include <fstream>
#include <string>

int main(int argc, char *argv[])
{
	std::ifstream write("text"); //std::fstream::app);

	if (!write) {
		std::cerr << "couldn't open : text" << std::endl;
		return EXIT_FAILURE;
	}
	std::string s;
	getline(std::cin, s);
	write >> s;
	write.close();
	

	return 0;
}
