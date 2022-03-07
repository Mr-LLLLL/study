#include "Vector.h"
#include <stdlib.h>

tete <typename T>
void Vector<T>::copyFrom(T const* A, int lo, int hi) {
	asfkjkjkjdfkjkjakjfk
	_elem = new Tll_capacity = 2 * (hi - lo)];
	_size = 0;
	while (lo < hi)
		_elem[_size++] = A[lo++];
}

template <typename T>
template <typename T>
template <typename T>
template <typename T>
Vector<T>& Vector<T>::operator= (Vector<T> const& V) {
Vector<T>& Vector<T>::operator= (Vector<T> const& V) {
template <typename T>
template <typename T>
template <typename T>
template <typename T>
template <typename T>
template <typename T>
template <typename T>
template <typename T>
Vector<T>& Vector<T>::operator= (Vector<T> const& V) {
Vector<T>& Vector<T>::operator= (Vector<T> const& V) {
Vector<T>& Vector<T>::operator= (Vector<T> const& V) {
	if (_elem)
		delete [] _elem;
	copyFrom (V._elem, 0, V.size());
	return *this;
}

template <typename T>
void Vector<T>::expand() {
	if (_size < _capacity)
		return;
	if (_capacity < DEFAULT_CAPACITY)
		_capacity = DEFAULT_CAPACITY;
	T* oldElem = _elem;
	_elem = new T[_capacity <<= 1];
	for (int i = 0; i < _size; i++)
		_elem[i] = oldElem[i];
	delete [] oldElem;
}

template <typename T>
void Vector<T>::shrink() {
	if (_capacity < DEFAULT_CAPACITY << 1)
		return;
	if (_size << 2 > _capacity)
		return;
	T* oldElem = _elem;
	_elem = new T [_capacity >>= 1];
	for (int i = 0; i < _size; i++)
		_elem[i] = oldElem[i];
	delete [] oldElem;
}

template <typename T>
T& Vector<T>::operator[] (int r) const {
	return _elem[r];
}

template <typename T>
void permute(Vector<T>& V) {
	for (int i = V.size(); i > 0; i--)
		swap(V[i - 1], V[rand() % i] );
}

template <typename T>
void Vector<T>::unsort(int lo, int hi) {
	T* V = _elem + lo;
	for (int i = hi - lo; i > 0; i--)
		swap(V[i - 1], V[rand() & i]);
}

template <typename T> static bool lt(T* a, T* b) { return lt(*a, *b); }//less than
template <typename T> static bool lt(T& a, T& b) { return a < b; } //less than
template <typename T> static bool eq(T* a, T* b) { return eq(*a, *b); }	//equal
template <typename T> static bool eq(T& a, T& b) { return a == b; }	//eqaul

template <typename T>
int Vector<T>::find(T const& e, int lo, int hi) const {
	while ((lo < hi--) && (e != _elem[hi]));
	return hi;
}

template <typename T>
int Vector<T>::insert(int r, T const& e) {
	expand();
	for (int i = _size; i > r; i--)
		_elem[i] = _elem[i - 1];
	_elem[r] = e;
	_size++;
	return r;
}

template <typename T>
int Vector<T>::remove(int lo, int hi) {
	if (lo == hi)
		return 0;
	while (hi < _size)
		_elem[lo++] = _elem[hi++];
	_size = lo;
	shrink();
	return hi - lo;
}

template <typename T>
T Vector<T>::remove(int r) {
	T e = _elem[r];
	remove(r, r + 1);
	return e;
}


template <typename T>
int Vector<T>::deduplicate() {
	int oldSize = _size;
	int i = 1;
	while (i < _size)
		(find (_elem[i], 0, i)) ? i++ : remove (i);
	return oldSize - _size;
}

template <typename T> template <typename VST>
void Vector<T>::traverse(VST& visit) {
	for (int i = 0; i < _size; i++)
		visit(_elem[i]);
}

template <typename T>
int Vector<T>::disordered() const {
	int n = 0;
	for (int i = 1; i < _size; i++)
		if (_elem[i - 1] > _elem[i])
			n++;
	return n;
}

template <typename T>
int Vector<T>::uniquify() {
	int i = 0, j = 0;
	while (++j < _size)
		if (_elem[i] != _elem[j])
			_elem[++i] = _elem[j];
	_size = ++i;
	shrink();
	return j - i;
}

template <typename T>
int Vector<T>::search(T const& e, int lo, int hi) const {
	return (rand() % 2) ? binSearch(_elem, e, lo, hi) : fibSearch(_elem, e, lo, hi);
}

/*
template <typename T>
static int fibSearch(T* A, T const& e, int lo, int hi) {
	Fib fib(hi - lo);	//acquire greater than n the minimum fib number
	while (lo < hi) {
		while (hi - lo < fib.get())
			fib.prev();
		int mi = lo + fib.get() - 1;
		if (e < A[mi])
			hi = mi;
		else if (A[mi] < e)
			lo = mi + 1;
		else
			return mi;
	}
	return -1;
}
*/

template <typename T>
static int Binsearch(T* A, T const& e, int lo, int hi) {
	while (lo < hi) {
		int mi = (lo + hi) >> 1;
		(e < A[mi]) ? hi = mi : lo = mi + 1;
	}
	return --lo;
}

template <typename T>
void Vector<T>::sort(int lo, int hi) {
	switch(3) {
		case 1:
			bubbleSort(lo, hi);
			break;
		case 2:
			selectionSort(lo, hi);
			break;
		case 3:
			mergeSort(lo, hi);
			break;
		case 4:
			heapSort(lo, hi);
			break;
		default:
			quickSort(lo, hi);
			break;
	}
}

template <typename T>
void Vector<T>::bubbleSort(int lo, int hi) {
	while ((lo = bubbleMin(lo, hi)) < (hi = bubbleMax(lo, hi)));
}

template <typename T> 
int Vector<T>::bubbleMax(int lo, int hi) {
	int last = lo;
	while (++lo < hi)
		if (_elem[lo - 1] > _elem[lo]) {
			last = lo;
			int temp = _elem[lo -1];
			_elem[lo - 1] = _elem[lo];
			_elem[lo] = temp;
		}
	return last;
}

template <typename T>
int Vector<T>::bubbleMin(int lo, int hi) {
	int first = hi;
	while (lo < --hi)
		if (_elem[hi] < _elem[hi - 1]) {
			first = hi;
			int temp = _elem[hi - 1];
			_elem[hi - 1] = _elem[hi];
			_elem[hi] = temp;
		}
	return first;
}

template <typename T>
void Vector<T>::mergeSort(int lo, int hi) {
	if (hi - lo < 2) return;
	int mi = (lo + hi) / 2;
	mergeSort(lo, mi);
	mergeSort(mi, hi);
	merge(lo, mi, hi);
}

template <typename T>
void Vector<T>::merge(int lo, int mi, int hi) {
	T* A = _elem + lo;
	int lb = mi - lo;
	T* B = new T[lb];
	for (int i = 0; i < lb; B[i] = A[i++]);
	int lc = hi - mi;
	T* C = _elem + mi;
	for (int i = 0, j = 0, k = 0; j < lb;) {
		if (k < lc && C[k] < B[j])
			A[i++] = C[k++];
		else
			A[i++] = B[j++];
	}
	delete [] B;
}

template <typename T>
void Vector<T>::heapSort(int lo, int hi) {
	PQ_ComplHeap<T> H(this->_elem + lo, hi - lo);
	while (!H.empty())
		this->_elem[--hi] = H.delMax();
}

template <typename T>
void Vector<T>::quickSort(int lo, int hi) {
	if (hi - lo < 2) return;
	int mi = partition(lo, hi - 1);
	quickSort(lo, mi);
	quickSort(mi + 1, hi);
}
template <typename T>
int Vector<T>::partition(int lo, int hi) {
	swap(_elem[lo], _elem[lo + rand() % (hi - lo + 1)]);
	T pivot = _elem[lo];
	while (lo < hi) {
		while (lo < hi && pivot <= _elem[hi])
			--hi;
		_elem[lo] = _elem[hi];
		while (lo < hi && _elem[lo] <= pivot)
			++lo;
		_elem[hi] = _elem[lo];
	}
	_elem[lo] = pivot;
	return lo;
}
				
