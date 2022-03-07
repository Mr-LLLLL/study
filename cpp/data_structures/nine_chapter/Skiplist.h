#ifndef SKIPLIST_H
#define SKIPLIST_H

#include <list>
#include "Entry.h"
#include "Quadlist.h"
#include "Dictionary.h"

using std::list;

template <typename K, typename V>
class Skiplist : public Dictionary<K, V>, public list<Quadlist<Entry<K, V>>*> {
protected:
	bool skipSearch (
			typename list<Quadlist<Entry<K, V>>*>::iterator &qlist,
			QlistNodePosi<Entry<K, V>> &p,
			K& k);
	void insertAsFirst(Quadlist<Entry<K, V>>* p);
public:
	int size() const {
		return this->empty() ? 0 : this->back()->size();
	}
	int level() {
		return size();
	}
	bool put(K, V);
	V* get(K k);
	bool remove (K k);
};
#include "Skiplist.c"
#endif
