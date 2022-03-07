#ifndef QUADLIST_H
#define QUADLIST_H

#include "QuadlistNode.h"

template <typename T>
class Quadlist {
private:
	int _size;
	QlistNodePosi<T> header;
	QlistNodePosi<T> trailer;
protected:
	void init();
	int clear();
public:
	Quadlist() { init(); }
	~Quadlist() { 
		clear();
		delete header;
		delete trailer;
	}
	int size() const { return _size; }
	bool empty() const { return _size <= 0; }
	QlistNodePosi<T> first() const { return header->succ; }
	QlistNodePosi<T> last() const { return trailer->pred; }
	bool valid(QlistNodePosi<T> p) {
		return p && (trailer != p) && (header != p);
	}

	T remove(QlistNodePosi<T> p);
	QlistNodePosi<T> insertAfterAbove(T const& e, QlistNodePosi<T> p, QlistNodePosi<T> b = nullptr);

	template <typename VST>
		void traverse(VST&);
};
#include "Quadlist.c"
#endif
