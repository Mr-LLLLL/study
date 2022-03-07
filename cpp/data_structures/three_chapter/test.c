#include <iostream>
#include <vector>
#include <algorithm>
#include <list>
using namespace std;


int main()
{
	list<int>* p = new list<int>{1,2,3};
	(*p).remove(++(*p).begin());
	
	
	for (auto iter : (*p))
		cout << iter << endl;
	
	return 0;
}
