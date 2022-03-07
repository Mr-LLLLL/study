#include "QuadlistNode.h"

template <typename T>
QlistNodePosi<T> QuadlistNode<T>::insertAsSuccAbove(T const& e, QlistNodePosi<T> b) {
	QlistNodePosi<T> x = new QuadlistNode<T> (e, this, succ, nullptr, b);
	succ->pred = x;
	succ = x;
	if (b)
		b->above = x;
	return x;
}
