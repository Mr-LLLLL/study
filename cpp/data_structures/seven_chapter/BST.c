#include "BST.h"
template <typename T>
static BinNodePosi(T)& searchIn(BinNodePosi(T)& v, const T& e, BinNodePosi(T)& hot){
	if (!v || (e == v->data))
		return v;
	hot = v;
	return searchIn(((e < v->data) ? v->lc : v->rc), e ,hot);
}

template <typename T>
static BinNodePosi(T)& searchIn_I(BinNodePosi(T) &v, const T& e, BinNodePosi(T)& hot) {
	if (!v || e == v->data)
		return v;
	hot = v;
	while (true) {
		BinNodePosi(T) & c = (e < hot->data) ? hot->lc : hot->rc;
		if (!c || e == c->data)
			return c;
		hot = c;
	}
}

template <typename T>
BinNodePosi(T)& BST<T>::search(const T& e){
	return searchIn(this->_root, e, _hot = nullptr);
}

template <typename T>
BinNodePosi(T) BST<T>::searchAll(const T& e) {
	BinNodePosi(T) res = searchIn_I(this->_root, e, _hot = nullptr);
	searchLEdge(res, e);
	SearchREdge(res, e);
	return res;
}

template <typename T>
static void searchLEdge(BinNodePosi(T) v, const T& e) {
	if (!v)
		return ;
	v = v->left;
	while (v) {
		if (e == v->data) {
			printRTree(v);
			v = v->left;
		} else 
			v = v->right;
	}
}

template <typename T>
static void searchREdge(BinNodePosi(T) v, const T& e) {
	if (!v)
		return ;
	v = v->right;
	while (v) {
		if (e == v->data) {
			printLTree(v);
			v = v->right;
		} else 
			v = v->left;
	}
}

template <typename T>
BinNodePosi(T) BST<T>::insert(const T& e) {
	BinNodePosi(T)& x = search(e);
	if (x)
		return x;
	x = new BinNodePosi(T)(e, _hot);
	this->_size++;
	updateHeightAbove(x);
	return x;
}

template <typename T>
bool BST<T>::remove(const T& e) {
	BinNodePosi(T)& x = search(e);
	if (!x)
		return false;
	removeAt(x, _hot);
	this->_size--;
	updateHeightAbove(_hot);
	return true;
}

template <typename T>
static BinNodePosi(T) removeAt(BinNodePosi(T)& x, BinNodePosi(T)& hot) {
	BinNodePosi(T) w = x;
	BinNodePosi(T) succ = nullptr;
	if (!HasLChild(*x))
		succ = x = x->rc;
	else if (!HasRChild(*x))
		succ = x = x->lc;
	else {
		w = w->succ();
		swap (x->data, w->dat);
		BinNodePosi(T) u = w->parent;
		((u == x) ? u->rc : u->lc) = succ = w->rc;
	}
	hot = w->parent;
	if (succ)
		succ->parent = hot;
	release(w->data);
	release(w);
	return succ;
}

template <typename T>
BinNodePosi(T) BST<T>::connect34(BinNodePosi(T) a, BinNodePosi(T) b, BinNodePosi(T) c, 
		BinNodePosi(T) T0, BinNodePosi(T) T1, BinNodePosi(T) T2, BinNodePosi(T) T3) {
	a->lc = T0;
	if (T0)
		T0->parent = a;
	a->rc = T1;
	if (T1)
		T1->parent = a;
	updateHeight(a);
	c->lc = T2;
	if (T2)
		T2->parent = c;
	c->rc = T3;
	if (T3)
		T3->parent = c;
	updateHeight(c);
	b->lc = a;
	a->parent = b;
	b->rc = c;
	c->parent = b;
	updateHeight(b);
	return b;
}

template <typename T>
BinNodePosi(T) BST<T>::rotateAt(BinNodePosi(T) v) {
	BinNodePosi(T) p = v->parent;
	BinNodePosi(T) g = p->parent;
	if (IsLChild(*p)) 
		if (IsLChild(*v)) {
			p->parent = g->parent;
			return connect34(v, p, g, v->lc, v->rc, p->rc, g->rc);
		} else {
			v->parent = g->parent;
			return connect34(p, v, g, p->lc, v->lc, v->rc, g->rc);
		}
	else 
		if (IsRChild(*v)){
			p->parent = g->parent;
			return connect34(g, p, v, g->lc, p->lc, v->lc, v->rc);
		} else {
			v->parent = g->parent;
			return connect34(g, v, p, g->lc, v->lc, v->rc, p->rc);
		}
}

template <typename T>
void zig(BinNodePosi(T) x) {
	FromParent(*x) = x->lc;
	x->lc->parent = x->parent;
	x->parent = x->lc;
	x->lc = x->parent->rc;
	x->parent->rc = x;
	if (x->lc)
		x->lc->parent = x;
}

template <typename T>
void zag(BinNodePosi(T) x) {
	FromParent(*x) = x->rc;
	x->rc->parent = x->parent;
	x->parent = x->rc;
	x->rc = x->parent->rc;
	x->parent->lc = x;
	if (x->rc)
		x->rc->parent = x;
}

template <typename T>
void strtchByZag(BinNodePosi(T)& x) {
	int h = 0;
	BinNodePosi(T) p = x;
	while (p->rc)
		p = p->rc;	//the final tree's root
	while (x->lc)
		x = x->lc;	//the final tree's leaf
	x->height = h++;
	for (; x != p; x = x->parent, x->height = h++) {
		while (x->rc)
			x->zag();
	}
}
