#include <iostream>
#include <forward_list>
#include <memory>

using namespace std;

int main()
{
	forward_list<int> l(5);
	for (int i : l)
		cout << i << endl;
	return 0;
}
