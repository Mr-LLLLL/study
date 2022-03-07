#include "PQ_LeftHeap.h"

template <typename T>
static BinNodePosi(T) merge_R(BinNodePosi(T) a, BinNodePosi(T) b) {
	if (nullptr == a)
		return b;
	if (nullptr == b)
		return a;
	if (a->data < b->data) {
		BinNodePosi(T) temp = a;
		a = b;
		b = temp;
	}
	a->rc = merge(a->rc, b);
	a->rc->parent = a;
	if (nullptr == a->lc || a->lc->npl < a->rc->npl) {
		BinNodePosi(T) temp = a->lc;
		a->lc = a->rc;
		a->rc = temp;
	}
	a->npl = a->rc ? a->rc->npl + 1 : 1;
	return a;
}

template <typename T>
static BinNodePosi(T) merge(BinNodePosi(T) a, BinNodePosi(T) b) {
	if (nullptr == a)
		return b;
	if (nullptr == b)
		return a;
	BinNodePosi(T) root = a;
	if (a->data < b->data)
		root = b;
	
	while (true) {
		if (a->data < b->data) {
			BinNodePosi(T) temp = a;
			a = b;
			b = temp;
			a->parent = b->parent;
			a->parent->rc = b;
		}
		if (nullptr == a->rc)
			break;
		a = a->rc;
	} 
	a->rc = b;
	b->parent = a;
	while (true) {
		if (a->rc != nullptr && (nullptr == a->lc || a->lc->npl < a->rc->npl)) {
			b = a->lc;
			a->lc = a->rc;
			a->rc = b;
		}
		a->npl = a->rc ? a->rc->npl + 1 : 1;
		if (root == a)
			break;
		a = a->parent;
	}

	return root;
}

	
template <typename T>
T PQ_LeftHeap<T>::delMax() {
	BinNodePosi(T) lHeap = this->_root->lc;
	BinNodePosi(T) rHeap = this->_root->rc;
	T e = this->_root->data;
	delete this->_root;
	--this->_size;
	this->_root = merge(lHeap, rHeap);
	if (this->_root)
		this->_root->parent = nullptr;
	return e;
}

template <typename T>
void PQ_LeftHeap<T>::insert(T e) {
	BinNodePosi(T) v = new BinNode<T>(e);
	this->_root = merge(this->_root, v);
	this->_root->parent = nullptr;
	++this->_size;
}

template <typename T>
T PQ_LeftHeap<T>::getMax() {
	return this->_root->data;
}
