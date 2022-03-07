#include <iostream>

using namespace std;

double sqrt(const double &d, const double precision) {
	double k = d;
	while (1) {
		if (k * k > d - precision && k * k < d + precision)
			break;
		k = 0.5 * (k + d / k);	//newton algorithm
	}
	return k;
}

int main(int argc, char** argv)
{
	double d;
	while (cout << "please input a number:", cin >> d)
		cout << endl << "the square root is: " << sqrt(d, 0.1) << endl;

	return 0;
}
