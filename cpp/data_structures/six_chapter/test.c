#include <iostream>
#include <bitset>
#include "../ten_chapter/PQ_ComplHeap.h"
#include <set>

#include <stack>
#include <string>
#include <queue>
using namespace std;

int next(int n) {
	int res = 0;
	while (n) {
		int t = n % 10;
		res += t * t;
		n /= 10;
	}
	return res;
}

bool isHappy(int n) {
	int i1 = n;
	int i2 = next(n);

	while (i2 != i1) {
		i1 = next(i1);
		i2 = next(next(i2));
		cout << i1 << endl;
	}
	return i1 == 1;
}

int main()
{
	cout << sizeof(bool) << endl;
	return 0;
}
