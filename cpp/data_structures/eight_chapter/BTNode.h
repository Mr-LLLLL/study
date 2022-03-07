#ifndef BTNODE_H
#define BTNODE_H
#include "../two_chapter/Vector.h"
#define BTNodePosi(T) BTNode<T>*
template <typename T>
struct BTNode {
	BTNodePosi(T) parent;
	Vector<T> key;
	Vector<BTNodePosi(T)> child;
	BTNode() : parent(nullptr) {
		child.insert(nullptr);
	}
	BTNode(T e, BTNodePosi(T) lc = nullptr, BTNodePosi(T) rc = nullptr) {
		parent = nullptr;
		key.insert(e);
		child.insert(lc);
		child.insert(rc);
		if (lc)
			lc->parent = this;
		if (rc)
			rc->parent = this;
	}
};


#endif
