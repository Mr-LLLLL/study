#include <iostream>

using namespace std;

int fib(int n) {
	return (n < 2) ? n : fib(n - 1) + fib(n - 2);
}

int main (int argc, char** argv)
{
	int n = 4;
	int k = 5;
	while(k -= 1) {
		int j = 0;
		while(n -= 1)
			cout << j << endl;
	}
	
	return 0;
}
