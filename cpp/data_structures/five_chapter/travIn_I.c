#include "BinNode.h"
#include <stack>

using namespace std;

template <typename T>
static void goAlongLeftBranch (BinNodePosi(T) x, stack<BinNodePosi(T)>& S) {
	while (x) {
		S.push(x);
		x = x->lc;
	}
}

template <typename T, typename VST>
void travIn_I1 (BinNodePosi(T) x, VST& visit) {
	stack<BinNodePosi(T)> S;
	while (true) {
		goAlongLeftBranch(x, S);
		if (S.empty())
			break;
		x = S.top();
		S.pop();
		visit(x->data);
		x = x->rc;
	}
}

template <typename T, typename VST>
void travIn_I2 (BinNodePosi(T) x, VST& visit) {
	stack<BinNodePosi(T)> S;
	while (true)
		if (x) {
			S.push(x);
			x = x->lc;
		} else if (!S.empty()) {
			x = S.top();
			S.pop();
			visit (x->data);
			x = x->rc;
		} else
			break;
}

template <typename T> 
BinNodePosi(T) BinNode<T>::succ() {
	BinNodePosi(T) s = this;
	if (rc) {
		s = rc;
		while (HasLChild (*s))
			s = s->lc;
	} else {
		while (IsRChild (*s))
			s = s->parent;
		s = s->parent;
	}
	return s;
}

template <typename T> 
BinNodePosi(T) BinNode<T>::pre() {
	BinNodePosi(T) s = this;
	if (lc) {
		s = lc;
		while (HasRChild (*s))
			s = s->rc;
	} else {
		while (IsLChild (*s))
			s = s->parent;
		s = s->parent;
	}
	return s;
}

template <typename T, typename VST>
void travIn_I3 (BinNodePosi(T) x, VST& visit) {
	bool backtrack = false;
	while (true)
		if (!backtrack && HasLChild (*x))
			x = x->lc;
		else {
			visit (x->data);
			if (HasRChild (*x)) {
				x = x->rc;
				backtrack = false;
			} else {
				if (!(x = x->succ()))
					break;
				backtrack = true;
			}
		}
}

template <typename T, typename VST>
void travIn_I4 (BinNodePosi(T) x, VST& visit) {
	while (true) 
		if (HasLchild (*x))
			x = x->lc;
		else {
			visit (x->data);
			while (!HasRChild (*x))
				if (!(x = x->succ()))
					return;
				else
					visit (x->data);
			x = x->rc;
		}
}
