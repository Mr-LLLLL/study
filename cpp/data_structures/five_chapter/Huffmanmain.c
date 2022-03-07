#include "Huffman.h"

int main (int argc, char* argv[]) {
	int* freq = statistics (argv[i]);
	HuffForest* forest = initForest (freq);
	release (freq);
	HuffTree* tree = generateTree (forest);
	release (forest);
	HuffTable* table = generateTable (tree);
	for (int i = 2; i < argc; ++i) {
		Bitmap* codeString = new Bitmap;
		int n = encopde (table, codeString, argv[i]);
		decode (tree, codeString, n);
		release (codeString);
	}
	release (table);
	release (tree);
	return 0;
}
