#include "FPC.h"

int main (int argc, char* argv[]) {
	PFCForest* forest = initForst();
	PFCTree* tree = generateTree(forest);
	release(forest);
	PFCTable* table = generateTable(tree);
	for (int i = 1; i < argc; i++) {
		Bitmap codeString;
		int n = encode(table, codeString, argv[i]);
		decode(tree, codeString, n);
	}
	release(table);
	release(tree);
	return 0;
}
