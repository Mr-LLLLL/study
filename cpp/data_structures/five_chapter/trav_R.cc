#include "BinNode.h"

template <typename T, typename VST>
void travPre_R(BinNodePosi(T) x, VST& visit) {
	if (!x)
		return;
	visit (x->data);
	travPre_R(x->lc, visit);
	travPre_R(x->rc, visit);
}

template <typename T, typename VST>
void travPost_R(BinNodePosi(T) x, VST& visit) {
	if (!x)
		return;
	travPost_R(x->lc, visit);
	travPost_R(x->rc, visit);
	visit (x->data);
}

template <typename T, typename VST>
void travIn_R (BinNodePosi(T) x, VST& visit) {
	if (!x)
		return;
	travIn_R(x->lc, visit);
	visit (x->data);
	travIn_R(x->rc, visit);
}
