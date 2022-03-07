#include <iostream>
#include <unistd.h>

int g_val = 10;

int main(int argc, char* argv[])
{
	pid_t pid;
	int lval = 20;
	g_val++, lval += 5;
	std::cout << "hello" << std::endl;	

	pid = fork();
	if (pid == 0) {
		std::cout << "I'm child process , " << g_val << ", " << lval << std::endl;
	}
	else {
		std::cout << "Child pro id : " << pid << std::endl;
		pid = fork();
		g_val++, lval++;
		if (pid == 0) {
			std::cout << "I'm second child , " << g_val << ", " << lval << std::endl;
		} else {
			std::cout << "Child pro id : " << pid << std::endl;
		}
	}
	return 0;
}

