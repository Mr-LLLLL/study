#include <iostream>
#include <algorithm>
#include <vector>
#include <stack>

using namespace std;

class test {
public:
	test(int xx) : x(xx) {};
	bool operator==(test const& t) {
		return x == t.x;
	}
	bool operator!=(test const& t) {
		return !(*this == t);
	}
private:
	int x;
};


enum weekday{a, sun, mon, tue, thu, fri, sat};

int main()
{
	cout << (int)' '  << endl;
}
