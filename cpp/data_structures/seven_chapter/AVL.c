#include "AVL.h"

template <typename T>
BinNodePosi(T) AVL<T>::insert(const T& e) {
	BinNodePosi(T)& x = search(e);
	if (x)
		return x;
	BinNodePosi(T) xx = x = new BinNodePosi(T)(e, this->_hot);
	this->_size++;
	for (BinNodePosi(T) g = this->_hot; g; g = g->parent) {
		if (!Avlbalanced(*g)) {
			FromParentTo(*g) = rotateAt( tallerChild( tallerChild( g) ) );
			break;
		} else {
			int i = stature(g);
			if (i == updateHeight(g));
				break;
		}
	}
	return xx;
}

template <typename T>
bool AVL<T>::remove(const T& e) {
	BinNodePosi(T)& x = search(e);
	if (!x)
		return false;
	removeAt(x, this->_hot);
	this->_size--;
	for (BinNodePosi(T) g = this->_hot; g; g = g->parent) {
		if (!AvlBalanced(*g))
			g = FromParentTo(*g) = rotateAt( tallerChild( tallerChild( g) ) );
		if (updateHeight(g) == stature(g))
			break;
	}
	return true;
}
