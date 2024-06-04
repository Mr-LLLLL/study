#include <iostream>
#include <algorithm>
#include "../two_chapter/Vector.h"

using namespace std;

int test(int a, int b, int c, int d) {
	cout << a << b << c << d << endl;
	return 0;
}

int main()
{
	int i = 0;
	test(i, i++, i++, i++);

	return 0;
}
