#ifndef AVL_H
#define AVL_H
#include "BST.h"

template <typename T>
class AVL : public BST<T> {
public:
	virtual BinNodePosi(T) insert(const T& e) override;
	virtual bool remove(const T& e) override;
	bool Balanced(const BinNodePosi(T)& x) {
		return stature(x.lc) == stature(x.rc);
	}
	int BalFax(const BinNodePosi(T)& x) {
		return stature(x.lc) - stature(x.rc);
	}
	bool AvlBalanced(const BinNodePosi(T)& x) {
		return (-2 < BalFax(x) && 2 > BalFax(x));
	}
	BinNodePosi(T) tallerChild(const BinNodePosi(T)& x) {
		return stature(x->lc) > stature(x->rc) ? x->lc :
			(stature(x->lc) < stature(x->rc) ? x->rc :
			 (IsLChild(*x) ? x->lc : x->rc));
	}
				

};

#endif
