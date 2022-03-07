#include <stdio.h>
#include <unistd.h>
#include <cstring>
#include <unistd.h>
#include <stdlib.h>
#include <fstream>
#include <unistd.h>
#include <iostream>
#include <vector>
#include <set>
#include <map>
#include <functional>

#define NDEBUG

#include <cassert>



using std::cout;
using std::cin;
using std::endl;

void print();
class Test {
public:
	Test(int i = 2) : _i(i), _j(0) {
		cout << "Test()" << endl;
	}
	Test(int i, int j) : _i(i), _j(j) {
		cout << "test(int)" << endl;
	}
	Test(const Test&) = default;

	Test(const Test&& t) {
		_i = t._i;
		_j = t._j;
	}
	bool operator <(const Test& t) const {
		return _i < t._i;
	}
	bool operator >(const Test& t) const {
		return _i > t._i;
	}



	int _i;
	mutable int _j;
};

void test(int i)
{
	cout << i << endl;
}


int main(int argc, char **argv) 
{
	char *ch = "hello";
	std::string str(ch);
	cout << str << endl;

	std::string filename;
	filename = "hello.txt";
	FILE *p;
	p = fopen(filename.c_str(), "w");
	fputs("hello", p);
		
	return 0;
}



