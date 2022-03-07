#include <iostream>

using std::cin;
using std::cout;
using std::endl;

int fib(int n ) {
	int f = 1, g = 0;	//initialization:fib(1), fib(0)
	while ( 0 < n-- ) {
		g += f; f = g - f;
	}
	return g;
}

int main(int argc, char** argv)
{
	int i;
	while(cout << " input any integer: ", cin >> i) 
	cout << "the fib(" << i << ") = " << fib(i) << endl;	

	return 0;
}
