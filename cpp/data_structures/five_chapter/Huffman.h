#ifndef HUFFMAN_H
#define HUFFMAN_H
#define N_CHAR (0x80 - 0x20)	//printable char
struct HuffChar {
	char ch;
	int weight;
	HuffChar (char c = '^', int w = 0) : ch(c), weight(w) { }
	bool operator< (HuffChar const& hc) {
		return weight > hc.weight;
	}
	bool operator== (HuffChar const& hc) {
		return weight == hc.weight;
	}
};

#define HuffTree BinTree<HuffChar>
typedef List<HuffTree*> HuffForest;
#include "Bitmap.h"
typedef Bitmap HuffCode;
#include "Hashtable.h"
typedef Hashtable<char, char*> HuffTable;



#endif
