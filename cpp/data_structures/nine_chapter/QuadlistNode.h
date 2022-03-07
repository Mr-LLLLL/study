#ifndef QUADLISTNODE_H
#define QUADLISTNODE_H

#include "Entry.h"
template <typename T>
class QuadlistNode {
public:
	T entry;
	QuadlistNode<T>* pred;
	QuadlistNode<T>* succ;
	QuadlistNode<T>* above;
	QuadlistNode<T>* below;
	QuadlistNode(T e = T(), QuadlistNode<T>* p = nullptr, QuadlistNode<T>* s = nullptr, QuadlistNode<T>* a = nullptr, QuadlistNode<T>* b = nullptr) : entry(e), pred(p), succ(s), above(a), below(b) {}
	QuadlistNode<T>* insertAsSuccAbove(T const& e, QuadlistNode<T>* b = nullptr);
};

template <typename T>
using QlistNodePosi = QuadlistNode<T>*;
#include "QuadlistNode.c"
#endif
