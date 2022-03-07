#include <iostream>
#include <queue>
#include "PQ_ComplHeap.h"
#include <map>

using namespace std;

int main()
{
	Vector<int> v;
	v.insert(1);
	v.insert(5);
	v.insert(3);
	int* a;
	a = &v[0];
	PQ_ComplHeap<int> pq(a, 3);
	cout << pq.delMax() << endl;
	cout << pq.delMax() << endl;
	cout << pq.delMax() << endl;
	
		
	return 0;
}
