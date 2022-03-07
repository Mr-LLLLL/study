#include <iostream>
#include <cstdlib>
#include <vector>
#include <list>

using namespace std;
template <typename T>
class iterator1;

template <typename T>
struct Vector {
	T* p;
	iterator1<T> iter;

	Vector() {
		p = new int(6);
	}

	iterator1<T> begin() {
		return iterator1<T>(p);
	}
};

template <typename T>
struct iterator1 {
	T* i;
	iterator1(T* p = 0) : i(p) {}
	T& operator* () const {
		return *i;
	}
};
int main()
{
	int gambleSum;
	cout << "input your gamble's money please: ";
	cin  >> gambleSum;
	double meanSum = 0;
	int n = 0;
	for (int i = 10000; i > 0; --i) {
		int sum = gambleSum;
		int temp = 2 * sum;
		int j = 1;
		while (sum > 0 && sum < temp) {
			if (rand() & 1) {
				sum += j;
				j = 1;
			} else {
				sum -= j;
				j <<= 1;
			}
		}
		if (sum > gambleSum)
			n++;
		meanSum += sum;
	}
	cout << "the gamble's money of final mean is: " << meanSum / 10000 << endl;
	cout << n << " mans have profit" << endl;
	
	return 0;
}
