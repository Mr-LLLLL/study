#ifndef VECTOR_H
#define VECTOR_H

#include "../ten_chapter/PQ_ComplHeap.h"

template <typename T>
class PQ_ComplHeap;

const int DEFAULT_CAPACITY = 3;

template <typename T>
class Vector {
protected:
	int _size;
	int _capacity;
	T* _elem;
	void copyFrom(T const* A, int lo, int hi);
	void expand();
	void shrink();
	int bubbleMax(int lo, int hi);
	int bubbleMin(int lo, int hi);
	void bubbleSort(int lo, int hi);
	int max(int lo, int hi);
	void selectionSort(int lo, int hi);
	void merge(int lo, int mi, int hi);
	void mergeSort(int lo, int hi);
	int partition(int lo, int hi);
	void quickSort(int lo, int hi);
	void heapSort(int lo, int hi);
public:
	//construct function
	Vector (int c = DEFAULT_CAPACITY, int s = 0, T v = 0) {
		_elem = new T[_capacity = c];
		for (_size = 0; _size < s; _elem[_size++] = v);
	}
	Vector(T const* A, int n) {
		copyFrom(A, 0, n);
	}
	Vector(T const* A, int lo, int hi) {
		copyFrom(A, lo, hi);
	}
	Vector(Vector<T> const& V) {
		copyFrom(V._elem, 0, V._size);
	}
	Vector(Vector<T> const& V, int lo, int hi) {
		copyFrom(V._elem, lo, hi);
	}
	//destructed function
	~Vector() {
		delete [] _elem;
	}
	//readable
	int size() const { return _size; }
	bool empty() const { return !_size; }
	int disordered() const;
	int find(T const& e) const { return find(e, 0, _size); }
	int find(T const& e, int lo, int hi) const;
	int search(T const& e) const { return (0 >= _size) ? -1 : search(e, 0, _size); }
	int search(T const& e, int lo, int hi) const;
	//writable
	T& operator[] (int r) const;
	Vector<T> &operator= (Vector<T> const&);
	T remove(int r);
	int remove(int lo, int hi);
	int insert(int r, T const& e);
	int insert(T const& e) { return insert(_size, e); }
	void sort(int lo, int hi);
	void sort() { sort(0, _size); }
	void unsort(int lo, int hi);
	void unsort() { unsort(0, _size); }
	int deduplicate();//unsoted Vector
	int uniquify(); //sorted Vector
	//traverse
	void traverse(void(*)(T&));
	template <typename VST>	void traverse(VST&);
};

#include "Vector.c"
#endif
