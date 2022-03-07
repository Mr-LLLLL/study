#include "PQ_ComplHeap.h"

template <typename T>
T PQ_ComplHeap<T>::getMax() {
	return this->_elem[0];
}

template <typename T>
void PQ_ComplHeap<T>::insert(T e) {
	Vector<T>::insert(e);
	percolateUp(this->_size - 1);
}

template <typename T>
int PQ_ComplHeap<T>::percolateUp(int i) {
	T temp = this->_elem[i];
	while (ParentValid(i)) {
		int j = Parent(i);
		if (temp <= this->_elem[j])
			break;
		this->_elem[i] = this->_elem[j];
		i = j;
	}
	this->_elem[i] = temp;
	return i;
}

template <typename T>
T PQ_ComplHeap<T>::delMax() {
	T maxElem = this->_elem[0];
	this->_elem[0] = this->_elem[--this->_size];
	percolateDown(this->_size, 0);
	return maxElem;
}

template <typename T>
int PQ_ComplHeap<T>::percolateDown(int n, int i) {
	int j;
	int temp = this->_elem[i];
	while (i != (j = ProperParent(this->_elem, n, i, temp))) {
		this->_elem[i] = this->_elem[j];
		i = j;
	}
	this->_elem[i] = temp;
	return i;
}

template <typename T>
void PQ_ComplHeap<T>::heapify(int n) {
	for (int i = LastInternal(n); InHeap(n, i); --i)
		percolateDown(n, i);
}
