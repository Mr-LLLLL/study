#ifndef REDBLACK_H
#define REDBLACK_H
#include "../seven_chapter/BST.h"

template <typename T>
class RedBlack : public BST<T> {
protected:
	void solveDoubleRed(BinNodePosi(T) x);
	void solveDoubleBlack(BinNodePosi(T) x);
	int updateHeight(BinNodePosi(T) x);
public:
	BinNodePosi(T) insert(const T& e);
	bool remove(const T& e);
	bool IsBlack(BinNodePosi(T) p) {
		return !p || (RB_BLACK == p->color);
	}
	bool IsRed(BinNodePosi(T) p) {
		return !IsBlack(p);
	}
	bool BlackHeightUpdated(BinNode<T> x) {
		return (stature(x.lc) == stature(x.rc)) && (x.height == (IsRed(&x) ? stature(x.lc) : stature(x.lc) + 1));
	}
};
#endif
