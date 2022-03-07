#ifndef SPLAY_H
#define SPLAY_H
#include "../seven_chapter/BST.h"
template <typename T>
class Splay : public BST<T> {
protected:
	BinNodePosi(T) splay(BinNodePosi(T) v);
public:
	BinNodePosi(T) &search(const T& e);
	BinNodePosi(T) insert(const T& e);
	bool remove(const T& e);
};
#endif
