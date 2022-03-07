#include "BinNode.h"
#include <stack>

using namespace std;

template <typename T, typename VST>
void travPre_I1 (BinNodePosi(T) x, VST& visit) {
	stack <BinNodePosi(T)> s;
	if (x)
		s.push(x);
	while (!s.empty()) {
		x = s.pop();
		visit (x->data);
		if (HasRChild (*x))
			s.push(x->rChild);
		if (HasLChild (*x))
			s.push(x->lChild);
	}
}

template <typename T, typename VST>
static void visitAlongLeftBreanch (BinNodePosi(T) x, VST& visit, stack<BinNodePosi(T)>& S) {
	while (x) {
		visit (x->data);
		if (x->rc)
			S.push(x->rc);
		x = x->lc;
	}
}

template <typename T, typename VST>
void travPre_I2 (BinNodePosi(T) x, VST& visit) {
	stack<BinNodePosi(T)> S;
	while (true) {
		visitAlongLeftBranch (x, visit, S);
		if (S.empty())
			break;
		x = S.top();
		S.pop();
	}
}


