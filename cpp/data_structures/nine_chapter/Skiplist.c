#include "Skiplist.h"
#include <cstdlib>

template <typename K, typename V>
void Skiplist<K, V>::insertAsFirst(Quadlist<Entry<K, V>>* p) {
	this->push_front(p);
}

template <typename K, typename V>
V* Skiplist<K, V>::get(K k) {
	if (this->empty())
		return nullptr;
	typename list<Quadlist<Entry<K, V>>*>::iterator qlist = this->begin();
	QlistNodePosi<Entry<K, V>> p = (*qlist)->first();
	return skipSearch(qlist, p, k) ? &(p->entry.value) : nullptr;
}

template <typename K, typename V>
bool Skiplist<K, V>::skipSearch(typename list<Quadlist<Entry<K, V>>*>::iterator &qlist,
		QlistNodePosi<Entry<K, V>> &p, K& k) {
	while (true) {
		while (p->succ && (p->entry.key <= k))
			p = p->succ;
		p = p->pred;
		if (p->pred && (k == p->entry.key))
			return true;
		qlist++;
		if (qlist == this->end())
			return false;
		p = (p->pred) ? p->below : (*qlist)->first();
	}
}

template <typename K, typename V>
bool Skiplist<K, V>::put(K k, V v) {
	Entry<K, V> e = Entry<K, V> (k, v);
	if (this->empty())
		insertAsFirst(new Quadlist<Entry<K, V>>);
	typename list<Quadlist<Entry<K, V>>*>::iterator qlist = this->begin();
	QlistNodePosi<Entry<K, V>> p = (*qlist)->first();
	if (skipSearch(qlist, p, k))
		while (p->below)
			p = p->below;
	qlist = --this->end();
	QlistNodePosi<Entry<K, V>> b = (*qlist)->insertAfterAbove(e, p);
	while (rand() & 1) {
		while ((*qlist)->valid(p) && !p->above)
			p = p->pred;
		if (!((*qlist)->valid(p))) {
			if (qlist == this->begin()) {
				insertAsFirst (new Quadlist<Entry<K, V>>);
				qlist = ++this->begin();
			}
			typename list<Quadlist<Entry<K, V>>*>::iterator temp = --qlist;
			p = (*temp)->first()->pred;
			++qlist;
		} else
			p = p->above;
		--qlist;
		b = (*qlist)->insertAfterAbove(e, p, b);
	}
	return true;
}

template <typename K, typename V>
bool Skiplist<K, V>::remove(K k) {
	if (this->empty())
		return false;
	typename list<Quadlist<Entry<K, V>>*>::iterator qlist = this->begin();
	QlistNodePosi<Entry<K, V>> p = (*qlist)->first();
	if (!skipSearch(qlist, p, k))
		return false;
	do {
		QlistNodePosi<Entry<K, V>> lower = p->below;
		(*qlist)->remove(p);
		p = lower;
		qlist++;
	} while (qlist != this->end());
	while (!this->empty() && this->front()->empty())
		this->pop_front();
	return true;
}

