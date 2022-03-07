#include "GraphMatrix.h"
#include "../ten_chapter/PQ_ComplHeap.h"

template <typename Tv, typename Te>
template <typename PU>
void Graph<Tv, Te>::pfs(int s, PU prioUpdater) {
	reset();
	int v = s;
	do 
		if (UNDISCOVERED == status(v))
			PFS(v, prioUpdater);
	while (s != (v = (++v % n)));
}

template <typename Tv, typename Te>
template <typename PU>
void Graph<Tv, Te>::PFS(int s, PU prioUpdater) {
	priority(s) = 0;
	status(s) = VISITED;
	parent(s) = -1;
	while (1) {
		for (int w = firstNbr(s); -1 < w; w = nextNbr(s, w))
			prioUpdater(this, s, w);
		for (int shortest = INT_MAX, w = 0; w < n; ++w)
			if (UNDISCOVERED == status(w))
				if (shortest > priority(w)) {
					shortest = priority(w);
					s = w;
				}
		if (VISITED == status(s))
			break;
		status(s) = VISITED;
		type(parent(s), s) = TREE;
	}
}

// priority_queue of Prim method
template <typename Tv, typename Te>
void Graph<Tv, Te>::prim(int s) {
	PQ_ComplHeap<pair<int, int>> pq;
	pq.insert(pair<int, int>(INT_MAX, s));
	parent(s) = -1;
	status(s) = VISITED;
	while (!pq.empty()) {
		for (int w = firstNbr(s); -1 < w; w = nextNbr(s, w)) {
			if (w == UNDISCOVERED) {
				this->parent(w) = s;
				pq.insert(pair<int, int>(-1 - this->weight(s, w), w));
			}
		}
		do {
			s = pq.delMax().second;
		} while (status(s) == UNDISCOVERED);
		status(s) = VISITED;
		type(parent(s), s) = TREE;
	}
}


// priority_queue of Dijkstra method
template <typename Tv, typename Te>
void Graph<Tv, Te>::dijkstra(int s) {
	PQ_ComplHeap<pair<int, int>> pq;
	pq.insert(pair<int, int>(INT_MAX, s));
	parent(s) = -1;
	status(s) = VISITED;
	while (!pq.empty()) {
		for (int w = firstNbr(s); -1 < w; w = nextNbr(s, w)) {
			if (w == UNDISCOVERED) {
				this->parent(w) = s;
				pq.insert(pair<int, int>(-1 - this->weight(s, w) + this->weight(parent(s), s), w));
			}
		}
		do {
			s = pq.delMax().second;
		} while (status(s) == UNDISCOVERED);
		status(s) = VISITED;
		type(parent(s), s) = TREE;
	}
}


