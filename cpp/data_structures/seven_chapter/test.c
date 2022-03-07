#include <iostream>
#include <vector>

using namespace std;

class A {
private:
	void print() {
		cout << "a.print" << endl;
	}
public:
	void printd() {
		cout << "d.print" << endl;
		print();
	}
};

class B : public A {
private:
	void printc() {
		cout << "c.print" << endl;
	}
public:
	void printb() {
		printc();
		cout << "b.print" << endl;
	}
};

int main()
{
	vector<vector<int>> vv;
	vv.push_back(vector<int>());
	return 0;
}
