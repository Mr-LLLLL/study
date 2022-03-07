#include "GraphMatrix.h"
#include <stack>

using namespace std;

template <typename Tv, typename Te>
stack<Tv>* Graph<Tv, Te>::tSort(int s) {
	reset();
	int v = s;
	stack<Tv>* S = new stack<Tv>;
	do {
		if (UNDISCOVERED == status(v))
			if (!tSort(v, clock, S)) {
					while (!S->empty())
					S->pop();
					break;
			}
	} while (s != (v = (++v % n)));
}

template <typename Tv, typename Te>
bool Graph<Tv, Te>::TSort(int v, stack<Tv>* S) {
	status(v) = DISCOVERED;
	for (int u = firstNbr(v); -1 < u; u = nextNbr(v, u))
		switch (status(u)) {
			case UNDISCOVERED:
				if (!TSort(u, clock, S))
					return false;
				break;
			case DISCOVERD:
				return false;
				break;
		}
	status(v) = VISITED;
	S->push(vertex(v));
	return true;
}
					
