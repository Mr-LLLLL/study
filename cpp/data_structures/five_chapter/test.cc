#include <iostream>
#include <vector>
#include "BinTree.h"
#include "Trie.h"

using namespace std;

int main()
{
	Trie t;
	t.insert("hello");
	t.insert("hella");
	t.insert("baby");
	cout << t.search("hello") << endl;
	cout << t.searchPrefix("hel") << endl;
	cout << t.searchPrefix("he") << endl;
	cout << t.searchPrefix("hello") << endl;
	cout << t.search("baby") << endl;


	return 0;
}
