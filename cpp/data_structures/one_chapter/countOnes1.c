#include <iostream>

/*
 * count number of one in integer(binary)
*/

using namespace std;

int countOnes1(unsigned int n) {
	int ones = 0;
	while ( 0 < n ) {
		ones++;
		n &= n - 1;
	}
	
	return ones;
}

int main(int argc, char** argv)
{
	unsigned int n;
	while( cout << "input number be counted:" ,cin >> n)
	cout << "the integer(binary) consist " << countOnes1(n) << " 1" << endl;

	return 0;
}
