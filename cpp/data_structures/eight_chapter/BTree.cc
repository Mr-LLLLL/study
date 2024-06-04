#include "BTree.h"


template <typename T>
BTNodePosi(T) BTree<T>::search(const T& e) {
	BTNodePosi(T) v = this->_root;
	this->_hot = nullptr;
	while (v) {
		int r = v->key.search(e);
		if ((0 <= r) && (e == v->key[r]))
			return v;
		this->_hot = v;
		v = v->child[r + 1];
	}
	return nullptr;
}

template <typename T>
bool BTree<T>::insert(const T& e) {
	BTNodePosi(T) v = search(e);
	if (v)
		return false;
	int r = this->_hot->key.search(e);
	this->_hot->key.insert(r + 1, e);
	this->_hot->child.insert(r + 2, nullptr);
	this->_size++;
	solveOverflow(this->_hot);
	return true;
}

template <typename T>
void BTree<T>::solveOverflow(BTNodePosi(T) v) {
	if (this->_order >= v->child.size())
		return;
	int s = this->_order / 2;
	BTNodePosi(T) u = new BTNode<T>();
	u->key = Vector<T>(v->key, s + 1, this->_order);
	u->child = Vector<T>(v->child, s + 1, this->_order + 1);
	v->key.remove(s + 1, this->_order);
	v->child.remove(s + 1, this->_order + 1);
	if (u->child[0])
		for (int j = 0; j < this->_order - s; j++)
			u->child[j]->parent = u;
	BTNodePosi(T) p = v->parent;
	if (!p) {
		this->_root = p = new BTNode<T>();
		p->child[0] = v;
		v->parent = p;
	}
	int r = 1 + p->key.search(v->key[0]);
	p->key.insert(r, v->key.remove(s));
	p->child.insert(r + 1, u);
	u->parent = p;
	solveOverflow(p);
}

template <typename T>
bool BTree<T>::remove(const T& e) {
	BTNodePosi(T) v = search(e);
	if (!v)
		return false;
	int r = v->key.search(e);
	if (v->child[0]) {
		BTNodePosi(T) u = v->child[r + 1];
		while (u->child[0])
			u = u->child[0];
		v->key[r] =u->key[0];
		v = u;
		r = 0;
	}
	v->key.remove(r);
	v->child.remove(r + 1);
	this->_size--;
	solveUnderflow(v);
	return true;
}

template <typename T>
void BTree<T>::solveUnderflow(BTNodePosi(T) v) {
	if ((_order + 1) / 2 <= v->child.size())
		return;
	BTNodePosi(T) p = v->parent;
	if (!p) {
		if (!v->key.size() && v->child[0]) {
			this->_root = v->child[0];
			this->_root->parent = nullptr;
			v->child[0] = nullptr;
			release(v);
		}
		return;
	}
	int r = 0;
	while (p->child[r] != v)
		r++;
	if (0 < r) {
		BTNodePosi(T) ls = p->child[r - 1];
		if ((_order + 1) / 2 < ls->child.size()) {
			v->key.insert(0, p->key[r - 1]);
			p->key[r - 1] = ls->key.remove(ls->key.size() - 1);
			v->child.insert(0, ls->child.remove(ls->child.size() - 1));
			if (v->child[0])
				v->child[0]->parent = v;
			return;
		}
	}

	if (p->child.size() - 1 > r) {
		BTNodePosi(T) rs = p->child[r + 1];
		if ((this->_order + 1) / 2 < rs->child.size()) {
			v->key.insert(v->key.size(), p->key[r]);
			p->key[r] = rs->key.remove(0);
			v->child.insert(v->child.size(), rs->child.remove(0));
			if (v->child[v->child.size() - 1])
				v->child[v->child.size() - 1]->parent = v;
			return;
		}
	}

	if (0 < r) {
		BTNodePosi(T) ls = p->child[r -1];
		ls->key.insert(ls->key.size(), p->key.remove(r - 1));
		p->child.remove(r);
		ls->child.insert(ls->child.size(), v->child.remove(0));
		if (ls->child[ls->child.size() - 1])
			ls->child[ls->child.size() - 1]->parent = ls;
		while (!v->key.empty()) {
			ls->key.insert(ls->key.size(), v->key.remove(0));
			ls->child.insert(ls->child.size(),v->child.remove(0));
			if (ls->child[ls->child.size() - 1])
				ls->child[ls->child.size() - 1]->parent = ls;
		}
		release(v);
	} else {
		BTNodePosi(T) rs = p->child[r + 1];
		rs->key.insert(0, p->key.remove(r));
		p->child.remove(r);
		rs->child.insert(0, v->child.remove(v->child.size() - 1));
		if (rs->child[0])
			rs->child[0]->parent = rs;
		while (!v->key.empty()) {
			rs->key.insert(0, v->key.remove(v->key.size() - 1));
			rs->child.insert(0, v->child.remove(v->child.size() - 1));
			if (rs->child[0])
				rs->child[0]->parent = rs;
		}
		release(v);
	}
	solveUnderflow(p);
	return;
}
