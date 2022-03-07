#include <iostream>
#include <vector>

using namespace std;

bool isPalindrome(int x) {
	if (x < 0 || (x % 10 == 0 && x != 0))
		return false;
	int revertedNumber = 0;
	while (x > revertedNumber) {
		revertedNumber = revertedNumber * 10 + x % 10;
		x /= 10;
	}
	return revertedNumber == x || x == revertedNumber / 10;
}

int main()
{
	int x;
	while (cout << "please input number:", cin >> x)
		cout << isPalindrome(x) << endl;

	return 0;
}
