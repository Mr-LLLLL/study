#include "Splay.h"

template <typename NodePosi> inline
void attachAsLChild(NodePosi p, NodePosi lc) {
	p->lc = lc;
	if (lc)
		lc->parent = p;
}

template <typename NodePosi> inline
void attachAsRChild(NodePosi p, NodePosi rc) {
	p->rc = rc;
	if (rc)
		rc->parent = p;
}

template <typename T>
BinNodePosi(T) Splay<T>::splay(BinNodePosi(T) v) {
	if (!v)
		return nullptr;	
	BinNodePosi(T) p, g;
	while ((p = v->parent) && (g = p->parent)) {
		BinNodePosi(T) gg = g->parent;
		if (IsLChild(*v))
			if (IsLChild(*p)) {
				attachAsLChild(g, p->rc);
				attachAsLChild(p, v->rc);
				attachAsRChild(p, g);
				attachAsRChild(v, p);
			} else {
				attachAsRChild(g, v->lc);
				attachAsLChild(p, v->rc);
				attachAsLChild(v, g);
				attachAsRChild(v, p);
			}
		else if (IsRChild(*p)) {
			attachAsRChild(g, p->lc);
			attachAsLChild(p, g);
			attachAsRChild(p, v->lc);
			attachAsLChild(v, g);
		} else {
			attachAsRChild(v, g);
			attachAsLChild(v, p);
			attachAsRChild(p, v->lc);
			attachAsLChild(g, v->rc);
		}
		if (!gg)
			v->parent = nullptr;
		else
			(g == gg->lc) ? attachAsLRchild(gg, v) : attachAsRChild(gg, v);
		updateHeightaAbove(g);
	}
	if (p = v->parent) {
		if (IsLChild(*v)) {
			attachAsRChild(v, p);
			attachAsLChild(p, v->rc);
		} else {
			attachAsLChild(v, p);
			attachAsRChild(g, v->lc);
		}
		updateHeight(p);
		updateHeight(v);
	}
	v->parent = nullptr;
	return v;
}
			
template <typename T>
BinNodePosi(T) &Splay<T>::search(const T& e) {
	BinNodePosi(T) p = searchIn(this->_root, e, this->_hot = nullptr);
	this->_root = splay(p ? p : this->_hot);
	return this->_root;
}

template <typename T>
BinNodePosi(T) Splay<T>::insert(const T& e) {
	if (!this->_root) {
		this->_size++;
		return this->_root = new BinNode<T>(e);
	}
	if (e == search(e)->data)
		return this->_root;
	this->_size++;
	BinNodePosi(T) t = this->_root;   
	if (this->_root->data < e) {
		t->parent = this->_root = new BinNode<T>(e, nullptr, t, t->rc);
		if (HasRChild(*t)) {
			t->rc->parent = this->_root;
			t->rc = nullptr;
		}
	} else {
			t->parent = this->_root = new BinNode<T>(e, nullptr, t->lc, t);
			if (HasLChild(*t)) {
				t->lc->parent = this->_root;
				t->lc = nullptr;
			}
	}
	updateHeightAbove(t);
	return this->_root;
}

template <typename T>
bool Splay<T>::remove(const T& e) {
	if (!this->_root || (e != search(e)->data))
		return false;
	BinNodePosi(T) w = this->_root;
	if (!HasLChild(*this->_root)) {
		this->_root = this->_root->rc;
		if (this->_root)
			this->_root->parent = nullptr;
	} else if (!HasRChld(*this->_root)) {
		this->_root = this->_root->lc;
		if (this->_root)
			this->_root->parent = nullptr;
	} else {
		BinNodePosi(T) lTree = this->_root->lc;
		lTree->parent = nullptr;
		this->_root->lc = nullptr;
		search(w->data);
		this->_root->lc = lTree;
		lTree->parent = this->_root;
	}
	release(w->data);
	release(w);
	this->_size--;
	if (this->_root)
		updateHeight(this->_root);
	return true;
}


