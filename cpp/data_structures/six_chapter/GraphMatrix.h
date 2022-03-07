#ifndef GRAPHMATRIX_H
#define GRAPHMATRIX_H
#include <vector>
#include "Graph.h"

using namespace std;

template <typename Tv>
struct Vertex {
	Tv data;
	int inDegree, outDegree;
	VStatus status;
	int dTime, fTime;
	int parent;
	int priority;
	Vertex (Tv const& d = (Tv)0) : data(0), inDegree(0), outDegree(0), status(UNDISCOVERED),
	dTime(-1), fTime(-1), parent(-1), priority(INT_MAX) {}
};

template <typename Te>
struct Edge {
	Te data;
	int weight;
	EType type;
	Edge (Te const& d, int w) : data(d), weight(w), type(UNDETERMINED){}
};

template <typename Tv, typename Te>
class GraphMatrix : public Graph<Tv, Te> {
private:
	vector< Vertex< Tv>> V;
	vector< vector< Edge< Te>*>> E;
public:
	GraphMatrix() { this->n = this->e = 0; }
	virtual ~GraphMatrix() {
		for (int j = 0; j < this->n; ++j)
			for (int k = 0; k < this->n; ++k)
				delete E[j][k];
	}

	virtual Tv& vertex(int i) override { return V[i].data; }
	virtual int inDegree(int i) override { return V[i].inDegree; }
	virtual int outDegree(int i) override { return V[i].outDegree; }
	virtual int firstNbr(int i) override { return nextNbr(i, this->n); }
	virtual int nextNbr(int i, int j) override {
		while ((-1 < j) && (!exists(i, --j)));
		return j;
	}
	virtual VStatus& status(int i) override { return V[i].dTime; }
	virtual int& dTime(int i) override { return V[i].dTime; }
	virtual int& fTime(int i) override { return V[i].fTime; }
	virtual int& paretn(int i) override { return V[i].parent; }
	virtual int& priority(int i) override { return V[i].priority; }

	virtual int insert(Tv const& vertex) override {
		for (int j = 0; j < this->n; ++j) 
			E[j].insert(nullptr);
		this->n++;
		E.insert(vector< Edge< Te>*>(this->n, this->n, (Edge<Te>*)nullptr));
		return V.insert(Vertex<Tv>(vertex));
	}
	virtual Tv remove(int i) override {
		for (int j = 0; j < this->n; ++j)
			if (exists(i, j)) {
				delete E[i][j];
				V[j].indegree--;
			}
		E.remove(i);
		this->n--;
		Tv vBak = vertex(i);
		V.remove(i);
		for (int j = 0; j < this->n; ++j)
			if (Edge<Te>* e = E[j].remove(i)) {
				delete e;
				V[j].outDegree--;
			}
		return vBak;
	}
	virtual bool exists(int i, int j) override {
		return (0 <= i) && (i < this->n) && (0 <= j) && (j < this->n) && E[i][j] != nullptr;
	}
	virtual EType& type(int i, int j) override { return E[i][j]->type; }
	virtual Te& edge (int i, int j) override { return E[i][j]->data; }
	virtual int& weight(int i, int j) override { return E[i][j]->weight; }

	virtual void insert(Te const& edge, int w, int i, int j) override {
		if (exists(i, j))
			return;
		E[i][j] = new Edge<Te>(edge, w);
		++this->e;
		++V[i].outDegree;
		++V[j].inDegree;
	}
	virtual Te remove(int i, int j) override {
		Te eBak = edge(i, j);
		delete E[i][j];
		E[i][j] = nullptr;
		--this->e;
		--V[i].outDegree;
		--V[i].inDegree;
		return eBak;
	}
};

#endif
