#include <iostream>

using namespace std;

template <typename T>
class Vector {
private:
	size_t _size, _capacity; T* _elem;
public:
	Vector (int c = 3, int s = 0, T v = 0) {
		_elem = new T[_capacity = c];
		for (_size = 0; _size < s; _elem[_size++] = v);
	}
	Vector (const Vector<T> &v) {
		copyFrom(v._elem, 0, v.size);
	}
	~Vector() {
		delete []_elem;
	}
	void copyFrom (T const *A, size_t lo, size_t hi) {
		_elem = new T[_capacity = 2 * (hi - lo)];
		_size = 0;
		while (lo < hi)
			_elem[_size++] = A[lo++];
	}
};

int main(int argc, char **argv)
{
	Vector<int> v;

	return 0;
}
