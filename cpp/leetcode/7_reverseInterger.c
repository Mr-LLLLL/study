#include <iostream>

using namespace std;

int reverse(int x) {
	int rev = 0;
	while (x != 0) {
		int pop = x % 10;
		x /= 10;
		if (rev > 0x7fffffff / 10 || (rev == 0x7fffffff / 10 && pop > 7))
			return 0;
		if (rev < (int)0x80000000 / 10 || (rev == (int)0x80000000 / 10 && pop < -8))
			return 0;
		rev = rev * 10 + pop;
	}
	return rev;
}

int main(int argc, char** argv)
{
	int n;
	while (cout << "please input number:", cin >> n)
		cout << reverse(n) << endl;

	return 0;
}
