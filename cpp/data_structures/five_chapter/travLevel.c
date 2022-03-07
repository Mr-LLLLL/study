#include "BinNode.h"
#include <queue>

using namespace std;

template <typename T>
template <typename VST>
void BinNode<T>::travLevel (VST& visit) {
	queue<BinNodePosi(T)> Q;
	Q.enqueue (this);
	while (!Q.empty()) {
		BinNodePosi(T) x = Q.dequeue();
		visit(x->data);
		if (HasLChild (*x))
			Q.enqueue(x->lc);
		if (HasRChild (*x))
			Q.enqueue(x->rc);
	}
}
