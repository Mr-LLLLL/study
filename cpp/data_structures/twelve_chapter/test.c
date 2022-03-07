#include <iostream>
#include <algorithm>
#include <string>
#include <cstdlib>

using namespace std;

int middle(int x, int y, int z);

int main()
{
	string s("hello");
	s.begin() = s.begin() + 2;
	cout << s << endl;
}

int middle(int x, int y, int z) {
	if (x <= y)
		if (x <= z)
			if (y <= z)
				return y;
			else
				return z;
		else
			return x;
	else
		if (y <= z)
			if (x <= z)
				return x;	
			else
				return z;
		else
			return y;
}



			
