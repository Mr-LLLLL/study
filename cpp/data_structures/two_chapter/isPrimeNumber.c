#include <iostream>

using namespace std;

class Bitmap {
private:
	int *F, N;
	int *T; int top;
protected:
	bool valid (int r) {
		return (0 <= r) && (r < top);
	}
	bool erased (int k) {
		return valid (F[k]) && !(T[ F[k] ] + 1 + k); //T[ F[k] ] = -1 - k;
	}
public:
	Bitmap (int n = 8) {
		N = n;
		F = new int[N];
		T = new int[N];
		top = 0;
	}
	~Bitmap() {
		delete [] F;
		delete [] T;
	}

	void set (int k) {
		if (test (k) )
			return;
		if (!erased (k))
			F[k] = top++;
		T[ F[k] ] = k;
	}
	void clear (int k) {
		if (test(k))
			T[ F[k] ] = -1 - k;
	}
	bool test (int k) {
		return valid (F[k]) && (k == T[ F[k] ]);
	}
};

double sqrt(const double &d, const double precision) {
	double k = d;
	while (1) {
		if (k * k > d - precision && k * k < d + precision)
			break;
		k = 0.5 * (k + d / k);	//newton algorithm
	}
	return k;
}

void eratosthenes(const int& n) {
	Bitmap B(n);
	B.set(0);
	B.set(1);
	for (int i = 2; i < sqrt(n, 1); ++i) 
		if (!B.test(i))
			for (int j = i * i; j < n; j += i) {
				B.set(j);
			}
	for (int i = 0; i < n; ++i)
		if (!B.test(i))
			cout << i << ", ";
	cout << endl;
}

void isPrimeNumber(const int& n) {
	eratosthenes(n);
}

int main(int argc, char** argv)
{
	int n;
	while(cout << "please input number:", cin >> n)
		isPrimeNumber(n);


	return 0;
}
