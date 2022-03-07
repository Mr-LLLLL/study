#include "Quadlist.h"

template <typename T>
void Quadlist<T>::init() {
	header = new QuadlistNode<T>;
	trailer = new QuadlistNode<T>;
	header->succ = trailer;
	trailer->pred = header;
	header->pred = trailer->succ = nullptr;
	header->below = trailer->above = nullptr;
	header->above = trailer->below = nullptr;
	this->_size = 0;
}

template <typename T>
QlistNodePosi<T> Quadlist<T>::insertAfterAbove(T const& e, QlistNodePosi<T> p, 
		QlistNodePosi<T> b) {
	this->_size++;
	return p->insertAsSuccAbove(e, b);
}

template <typename T>
T Quadlist<T>::remove(QlistNodePosi<T> p) {
	p->pred->succ = p->succ;
	p->succ->pred = p->pred;
	--this->_size;
	T e = p->entry;
	delete p;
	return e;
}

template <typename T>
int Quadlist<T>::clear() {
	int oldSize = this->_size;
	while (0 < this->_size)
		remove(header->succ);
	return oldSize;
}
