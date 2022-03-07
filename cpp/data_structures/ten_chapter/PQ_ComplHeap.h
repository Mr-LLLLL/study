#ifndef PQ_COMPLHEAP_H
#define PQ_COMPLHEAP_H

#include "../two_chapter/Vector.h"
#include "PQ.h"

template<typename T>
class Vector;

template <typename T>
class PQ_ComplHeap : public Vector<T>, public PQ<T> {
private:
	bool InHeap(int n, int i) {
		return -1 < i && i < n;
	}
	int Parent(int i) {
		return (i - 1) >> 1;
	}
	int LastInternal(int n) {
		return Parent(n - 1);
	}
	int LChild(int i) {
		return 1 + (i << 1);
	}
	int RChild(int i) {
		return (1 + i) << 1;
	}
	bool ParentValid(int i) {
		return 0 < i;
	}
	bool LChildValid(int n, int i) {
		return InHeap(n, LChild(i));
	}
	bool RChildValid(int n, int i) {
		return InHeap(n, RChild(i));
	}
	int Bigger(T* PQ, int i, int j) {
		return PQ[i] < PQ[j] ? j : i;
	}
	int Bigger(T* PQ, int i, int j, int temp) {
		return temp < PQ[j] ? j : i;
	}
	int ProperParent(T* PQ, int n, int i, int temp) {
		return RChildValid(n, i) ? Bigger(PQ, i, Bigger(PQ, RChild(i), LChild(i)), temp) :
			LChildValid(n, i) ? Bigger(PQ, i, LChild(i), temp) : i;
	}
protected:
	int percolateDown(int n, int i);
	int percolateUp(int i);
	void heapify(int n);
public:
	PQ_ComplHeap() {}
	PQ_ComplHeap(T* A, int n) {
		this->copyFrom(A, 0, n);
		heapify(n);
	}
	void insert(T) override;
	T getMax() override;
	T delMax() override;
};

#include "PQ_ComplHeap.c"
#endif
