#include "GraphMatrix.h"

template <typename Tv, typename Te>
void Graph<Tv, Te>::bcc(int s) {
	reset();
	int clock = 0;
	int v = s;
	stack<int> S;
	do
		if (UNDISCOVERED == status(v)) {
			BCC(v, clock, S);
			S.pop();
		}
	while (s != (v = (++v % n)));
}

#define hca(x) (fTime(x))

template <typename Tv, typename Te>
void Graph<Tv, Te>::BCC(int v, int& clock, stack<int>& S) {
	hca(v) = dTime = ++clock;
	status(v) = DISCOVERED;
	S.push(v);
	for (int u = firstNbr(v); -1 < u; u = nextNbr(v, u))
		switch (status(u)) {
			case UNDISCOVERED:
				parent(u) = v;
				type(v, u) = TREE;
				BCC(u, clock, S);
				if (hca(u) < dTime(v))
					hca(v) = hca(v) > hca(u) ? hca(u) : hca(v);
				else {
					while (v != S.top())
						S.pop();
					S.pop();
					S.push(v);
				}
				break;
			case DISCOVERED:
				type(v, u) = BACKWARD;
				if (u != parent(v))
					hca(v) = hca(v) > dTime(u) ? dTime(u) : hca(v);
				break;
			default:
				type(v, u) = (dTime(v) < dTime(u)) ? FORWARD : CROSS;
				break;
		}
	status(v) = VISITED;
}
#undef hca
