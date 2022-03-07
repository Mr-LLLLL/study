#include <iostream>

using namespace std;

class Solution {
public:
	/*
	 * Nweton's method
	 * f(x) = x^2 - x[0]
	 * g(x) = f(x[0]) + f'(x[0])(x - x[0])
	 * x = x[0] - f(x[0]) / f'(x[0])
	 * x[n+1] = x[n] - f(x[n]) / f'(x[n])
	 * x[n+1] = x[n] - (x[n]^2 - x[0]) / 2x[n]
	 * x[n+1] = x[n] / 2 + x[0] / 2x[n] 
	 * x[n+1] = (x[n] + x[0] / x[n]) / 2
	 */
	int sqrt(int x) {
		double r = x;
		while ((int)r > x / r) {
			r = (r + x / r) / 2;
		}
		return r;
	}
};

int main(int argc, char** argv)
{
	int x;
	while (cout << "input a number: " << endl, cin >> x)
		cout << "the number sqrt is : " << Solution().sqrt(x) << endl;


	return 0;
}
