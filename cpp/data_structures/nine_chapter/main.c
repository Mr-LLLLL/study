#include <list>
#include <iostream>
#include "Skiplist.h"
#include "Hashtable.h"

using namespace std;

void Eratosthenes(int n, char* file) {
	Bitmap B(n);
	B.set(0);
	B.set(1);
	for (int i = 2; i < n; ++i)
		if (!B.test(i))
			for (int j = i * i; j < n; j += i)
				B.set(j);
	B.dump(file);
}

int main()
{
	Skiplist<int, int> skip;
	skip.put(1, 1);
	skip.put(1, 2);
	skip.put(1, 3);
	skip.put(1, 4);
	skip.put(1, 5);
	skip.put(1, 6);
	skip.put(1, 7);
	skip.put(1, 8);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	skip.put(1, 9);
	cout << *skip.get(1) << endl;


	return 0;
}

