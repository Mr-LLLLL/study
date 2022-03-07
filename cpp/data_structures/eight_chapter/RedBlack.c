#include "RedBlack.h"

template <typename T>
int RedBlack<T>::updateHeight(BinNodePosi(T) x) {
	x->height = (stature(x->lc) > stature(x->rc) ? stature(x->lc) : stature(x->rc));
	return IsBlack(x) ? x->height++ : x->height;
}

template <typename T>
BinNodePosi(T) RedBlack<T>::insert(const T& e) {
	BinNodePosi(T) &x = search(e);
	if (x)
		return x;
	x = new BinNode<T>(e, this->_hot, nullptr, nullptr, -1);
	this->_size++;
	solveDoubleRed(x);
	return x ? x : this->_hot->parent;
}

template <typename T>
void RedBlack<T>::solveDoubleRed(BinNodePosi(T) x) {
	while (true) {
		if (IsRoot(*x)) {
			this->_root->color = RB_BLACK;
			this->_root->height++;
			return;
		}
		BinNodePosi(T) p = x->parent;
		if (IsBlack(p))
			return;
		BinNodePosi(T) g = p->parent;
		BinNodePosi(T) u = uncle(x);
		if (IsBlack(u)) {
			BinNodePosi(T) gg = g->parent;
			BinNodePosi(T) r = FromParentTo(*g) = rotateAt(x);
			r->parent = gg;
			r->color = RB_BLACK;
			r->rc->color = RB_RED;
			r->lc->color = RB_RED;
			return;
		} else {
			p->color = RB_BLACK;
			p->heigth++;
			u->color = RB_BLACK;
			u->height++;
			if (!IsRoot(*g))
				g->color = RB_RED;
			x = g;
		}
	}
}

template <typename T>
bool RedBlack<T>::remove(const T& e) {
	BinNodePosi(T) & x = search(e);
	if (!x)
		return false;
	BinNodePosi(T) r = removeAt(x, this->_hot);
	if (!(--this->_size))
		return true;
	if (!this->_hot) {
		this->_root->color = RB_BLACK;
		updateHeight(this->_root);
		return true;
	}

	if (BlackHeightUpdated(*this->_hot))
		return true;
	if (IsRed(r)) {
		r->color = RB_BLACK;
		r->height++;
		return true;
	}
	solveDoubleBlack(r);
	return true;
}

template <typename T>
void RedBlack<T>::solveDoubleBlack(BinNodePosi(T) r) {
	     BinNodePosi(T) p = r ? r->parent : this->_hot;
		 if (!p)
			 return;
		 BinNodePosi(T) s = (r == p->lc) ? p->rc : p->lc;
		 if (IsBlack(s)) {
			 BinNodePosi(T) t = nullptr;
			 if (IsRed(s->rc))
				 t = s->rc;
			 if (IsRed(s->lc))
				 t = s->lc;
			 if (t) {
				 RBColor oldColor = p->color;
				 BinNodePosi(T) b = FromParentTo(*p) = rotateAt(t);
				 if (HasLChild(*b)) {
					 b->lc->color = RB_BLACK;
					 updateHeight(b->lc);
				 }
				 if (HasRchild(*b)) {
					 b->rc->color = RB_BLACK;
					 updateHeight(b->rc);
				 }
			 } else {
				 s->color = RB_RED;
				 s->height--;
				 if (IsRed(p)) {
					 p->color = RB_BLACK;
				 } else {
					 p->height--;
					 solveDoubleBlack(p);
				 }
			 }
		 } else {
			 s->color = RB_BLACK;
			 p->color = RB_RED;
			 BinNodePosi(T) t = IsLChild(*s) ? s->lc : s->rc;
			 this->_hot = p;
			 FromParentTo(*p) = rotateAt(t);
			 solveDoubleBlack(r);
		 }
}
