#ifndef GRAPH_H
#define GRAPH_H
#include <limits.h>
#include <stack>

using namespace std;

enum VStatus {
	UNDISCOVERED, DISCOVERED, VISITED
};
enum EType {
	UNDETERMINED, TREE, CROSS, FORWARD, BACKWARD
};

template <typename Tv, typename Te>
class Graph {
private:
	void reset () {
		for (int i = 0; i < n; ++i) {
		status (i) = UNDISCOVERED;
		dTime (i) = fTime (i) = -1;
		parent (i) = -1;
		priority (i) = INT_MAX;
		for (int j = 0; i < n; ++j) 
			if (exist (i, j))
				type (i, j) = UNDETERMINED;
		}
	}
	void BFS (int, int&);	//broad-first search
	void DFS (int, int&);	//Depth-First Search
	void BCC (int, int&, stack<int>&);	//bi-connected component
	bool TSort (int, int&, stack<Tv>*);	//topological sorting
	template <typename PU>
	void PFS (int, PU);	//Priority-First Search
public:
	int n;
	virtual int insert (Tv const&) = 0;
	virtual Tv remove (int) = 0;
	virtual Tv& vertex (int) = 0;
	virtual int inDegree (int) = 0;
	virtual int outDegree (int) = 0;
	virtual int firstNbr (int) = 0;
	virtual int nextNbr (int, int) = 0;
	virtual VStatus& status (int) = 0;
	virtual int& dTime (int) = 0;
	virtual int& fTime (int) = 0;
	virtual int& parent (int) = 0;
	virtual int& priority (int) =0;
	
	int e;
	virtual bool exist (int, int) = 0;
	virtual void insert (Te const&, int, int, int) = 0;
	virtual Te remove (int, int) = 0;
	virtual EType& type (int, int) = 0;
	virtual Te& edge (int, int) = 0;
	virtual int& weight (int, int) = 0;

	void bfs (int);
	void dfs (int);
	void bcc (int);
	stack<Tv>* tSort (int);
	void prim (int);
	void dijkstra (int);
	template <typename PU>
	void pfs (int, PU);
};







#endif
