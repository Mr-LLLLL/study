#ifndef PQ_LEFTHEAP_H
#define PQ_LEFTHEAP_H

#include "PQ.h"
#include "../five_chapter/BinTree.h"

template <typename T>
class PQ_LeftHeap : public PQ<T>, public BinTree<T> {
public:
	PQ_LeftHeap() {}
	PQ_LeftHeap(T* E, int n) {
		for (int i = 0; i < n; ++i)
			insert(E[i]);
	}
	void insert(T) override;
	T getMax() override;
	T delMax() override;
};

#include "PQ_LeftHeap.c"

#endif
