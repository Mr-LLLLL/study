#include "GraphMatrix.h"
#include <stack>

using namespace std;

template <typename Tv, typename Te>
void Graph<Tv, Te>::dfs(int s) {
	reset();
	int clock = 0;
	int v = s;
	do 
		if (UNDISCOVERED == status(v))
			DFS(v, clock);
	while (s != (v = (++v % n)));
}

template <typename Tv, typename Te>
void Graph<Tv, Te>::DFS(int v, int& clock) {
	dTime(v) = ++clock;
	status(v) = DISCOVERED;
	for (int u = firstNbr(v); -1 < u; u = nextNbr(v, u))
		switch (status(u)) {
			case UNDISCOVERED:
				type(v, u) = TREE;
				parent(u) = v;
				DFS(u, clock);
				break;
			case DISCOVERED:
				type(v, u) = BACKWARD;
				break;
			default:
				type(v, u) = (dTime(v) < dTime(u)) ? FORWARD : CROSS;
				break;
		}
	status(v) = VISITED;
	fTime(v) = ++clock;
}

template <type Tv, typename Te>
void DFS_I(int v, int& clock) {
	stack<int> s;
	status(v) = DISCOVERED;
	dTime(v) = ++clock;
	s.push(v);
	while (!s.empty()) {
		v = s.top();
		s.pop();
		for (int u = firstNbr(v); -1 < u; u = nextNbr(v, n))
			switch (status(u)) {
				case UNDISCOVERD:
					type(v, u) = TREE;
					parent(u) = v;
					status(u) = DISCOVERED;
					s.push(u);
					dTime(u) = ++clock;
					v = u;
					break;
				case DISCOVERD:
					type(v, u) = BACKWARD;
					break;
				default:
					if (TREE == type(v, y))
						break;
					type(v, u) = (dTime(v) < dTime(u)) ? FORWARD : CROSS;
					break;
			}
		status(v) = VISITED;
		fTime(v) = ++clock;
	}
}

//base on PFS frame
template <typename Tv, typename Te> struct DfsPU {
	virtual void operator() (Graph<Tv, Te>* g, int uk, int v) {
		if (g->status(v) == UNDISCOVERD)
			if (g->priority(v) > g->priority(uk) - 1) {
				g->priotiry(v) = g->priority(uk) - 1;
				g->parent(v) = uk;
				return;
			}
	}
};

