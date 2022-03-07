#include "PFC.h"
#include <stdlib.h>

PFCForest* initForest() {
	PFCForest* forest = new PFCForest;
	for (int i = 0; i < N_CHAR; ++i) {
		forest->insert(i, new PFCTree());
		(*forest)[i]->insertAsRoot(0x20 + i);
	}
	return forest;
}

PFCTree* generateTree (PFCForest* forest) {
	srand ( (unsigned int ) time ( nullptr ) );
	while ( 1 < forest->size() ) {
		PFCTree* s = new PFCTree;
		s->insertAsRoot ( '^' );
		size_t r1 = rand() % forest->size();
		s->attachAsLC ( s->root(), ( *forest ) [r1] );
		forest->remove ( r1 );
		size_t r2 = rand() % forest->size();
		s->attachAsRC ( s->root(), ( *forest ) [r2] );
		forest->remove ( r2 );
		forest->insert ( forest->size(), s );
	}
	return ( *forest ) [0];
}

void generateCT ( Bitmap* code, int length, PFCTable* table, BinNodePosi ( char ) v ) {
	if (IsLeaf ( *v) ) {
		table->put(v->data, code->bits2string(length));
		return;
	}
	if (HasLChild (*v)) {
		code->clear (length);
		generateCT (code, length + 1, table, v->lc);
	}
	if (HasRChild (*v)) {
		code->set(length);
		generateCT(code, length + 1, table, v->rc);
	}
}

int encode (PFCTable* table, Bitmap& codeString, char* s) {
	int n = 0;
	for (size_t m = strlen(s), i = 0; i < m; ++i) {
		char** pCharCode = table->get(s[i]);
		if (!pCharCode) pCharCode = table->get(s[i] + 'A' - 'a');
		if (!pCharCode) pCharCode = table->get(' ');
		printf ("%s", *pCharCode);
		for (size_t m = strlen(*pCharCode), j = 0; j < m; ++j)
			'1' == *(*pCharCode + j) ? codeString.set(n++) : codeString.clear(n++);
	}
	return n;
}

void decode (PFCTree* tree, Bitmap& code, int n) {
	BinNodePosi(char) x = tree->root();
	for (int i = 0; i < n; ++i) {
		x = code.test(i) ? x->rc : x->lc;
		if (IsLeaf(*x)) {
			printf( "%c", x->data); 
			x = tree->root();
		}
	}
}
