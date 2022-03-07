#include <iostream>

using namespace std;

/*
 * recursion
inline int sqr(int a) {
	return a * a;
}

int power(int n, int a) {
	if (0 == n)
		return 1;
	return (n & 1) ? sqr(power(n >> 1, a)) * a : sqr(power(n >> 1, a));
}

int main(int argc, char** argv)
{
	int n, a;
	while (cout << "input base, power: ",cin >> a >> n)
		cout << "the " << n << " power of " << a << " is " << power(n, a) << endl;

	return 0;
}
*/

double power(double a, int n) {
	if (n < 0) {
		n = -n;
		a = 1 / a;
	}
	double res = 1;
	double aur_res = a;
	while (n != 0) {
		if (n & 1)
			res *= aur_res;
		aur_res *= aur_res;
		n >>= 1;
	}
	return res;
}

int main(int argc, char* argv[])
{
	double a;
	int n;
	while (cout << "input base, power: ",cin >> a >> n)
		cout << "the " << n << " power of " << a << " is " << power(a, n) << endl;

	return 0;
}

