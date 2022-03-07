#include <iostream>

using namespace std;


/*
 * count "1" in binary of integer
*/

#define POW(c) (1 << c) //2^c
#define MASK(c) (((unsigned long) -1) / (POW(POW(c)) + 1)) //
#define ROUND(n, c) (((n) & MASK(c)) + ((n) >> POW(c) & MASK(c)))

int countOnes2 (unsigned int n) {
	n = ROUND ( n, 0 );
	n = ROUND ( n, 1 );
	n = ROUND ( n, 2 );
	n = ROUND ( n, 3 );
	n = ROUND ( n, 4 );
	return n;
}

int main (int argc, char** argv)
{
	unsigned int n;
	while (cout << "input interger:", cin >> n)
		cout << "the binary of integer include  " << countOnes2(n) << " \"1\"" << endl;

	return 0;
}
