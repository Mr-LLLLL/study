#include <iostream>
#include <stack>

using namespace std;


void convert(stack<char>& s, int n, int base) {
	static char digit[] = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'};
	while (n > 0) {
		s.push(digit[n % base]);
		n /= base;
	}
}

int main(int argc, char** argv)
{
	int n, base;
	static stack<char> s;
	while (cout << "please input number and base:", cin >> n >> base) {
		cout << "the number of " << base << " is: ";
		convert(s, n, base);
		while (!s.empty()) {
			cout << s.top();
			s.pop();
		}
		cout << endl;
	}
				 
	

	return 0;
}
