#include <iostream>
#include "Vector.h"
#include "test.h"

using namespace std;

int main()
{
	Vector<int> v;
	for (int i = 4; i >= 0; i--)
		v.insert(i);
	v.sort(0, 5);
	for (int i = 0; i < 5; i++)
		cout << v[i] << ' ';
	


	return 0;
}
