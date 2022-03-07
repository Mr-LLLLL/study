#include "BinNode.h"
#include <stack>

using namespace std;

template <typename T>
static void gotoHLVFL (stack<BinNodePosi(T)>& S) {
	while (BinNodePosi(T) x = S.top())
		if (HasLChild (*x)) {
			if (HasRChild (*x))
				S.push(x->rc);
			S.push(x->lc);
		} else 
			S.push(x->rc);
	S.pop();
}

template <typename T, typename VST>
void travPost_I (BinNodePosi(T) x, VST& visit) {
	stack<BinNodePosi(T)> S;
	if (x)
		S.push(x);
	while (!S.empty()) {
		if (S.top() != x->parent)
			gotoHLVFL (S);
		x = S.top();
		S.pop();
		visit(x->data);
	}
}

template <typename T, typename UST>
void travPost_I2 (BinNodePosi(T) x, VST& visit) {
	stack<BinNodePosi(T) S, S1;
	if (x) S.push (x);
	while (!S.empty()) {
		x = S.pop();
		S1.push (x);
		if (HasLchild (*x))
			S.push (x->lc);
		if (HasRchild (*x))
			S.push (x->rc);
	}
	while (!S1.empty()) {
		visit(s1.top()->data);
		s1.pop();
	}
}


