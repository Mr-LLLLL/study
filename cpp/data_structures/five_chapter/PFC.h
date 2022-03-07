#ifndef PFC_H
#define PFC_H
#include "BinTree.h"
#include <vector>
#include "bitmap.h"
#include "skiplist.h"

using namespace std;

typedef BinTree<char> PFCTree;
typedef vector<PFCTree*> PFCForest;
typedef skiplist<char, char*> PFCTable;

#define N_CHAR (0x80 - 0x20) //consider printable char

PFCForest* initForest();
PFCTree* generateTree (PFCForest* );
void generateCT (Bitmap*, int, PFCTable*, BinNodePosi(char));
PFCTable* generateTable(PFCTree*);
int encode(PFCTable*, Bitmap&, char*);
void decode(PFCTree*, Bitmap&, int);


#endif
