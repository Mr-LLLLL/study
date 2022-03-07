#include <iostream>

using namespace std;

template <typename T> class Hailstone {		//make a Hailstone object
public:
		virtual void operator() (T &e) {		//presume T can do arithmetic operation
		int step = 0;						//conversion steps
		while (1 != e) {
			e & 1 ? e = 3 * e + 1 : e >>= 1;
			++step;
		}

		e = step;							//conversion steps
	}
};


int main(int argc, char** argv)
{
	Hailstone<int> hail;
	int n;
	while(cout << "please input one number: ", cin >> n) {
		hail(n);
		cout << "the Hailstone length is: " << n << endl;
	}

	return 0;
}
