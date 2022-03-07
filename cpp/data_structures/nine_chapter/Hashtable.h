#ifndef HASHTABLE_H
#define HASHTABLE_H

#include "Entry.h"
#include "Dictionary.h"
#include "Bitmap.h"

template <typename K, typename V>
class Hashtable : public Dictionary<K, V> {

private:
	Entry<K, V>** ht;
	int M;
	int N;
	Bitmap* lazyRemoval;
	bool lazilyRemoved(int i) {
		return lazyRemoval->test(i);
	}
	void markAsRemoved(int i) {
		lazyRemoval->set(i);
	}
protected:
	int probe4Hit(const K& k);
	int probe4Free(const K& k);
	void rehash();
public:
	Hashtable(int c = 5);
	~Hashtable();
	int size() const { return N; }
	bool put(K, V);
	V* get(K k);
	bool remove(K k);
};

static size_t hashCode(char c) { return (size_t)c; }
static size_t hashCode(int k) { return (size_t)k; }
static size_t hashCode(long long i) { return (size_t)((i >> 32) + (int)i); }
static size_t hashCode(char s[]) {
	int h = 0;
	for (size_t n = strlen(s), i = 0; i < n; ++i) {
		h = (h << 5) | (h >> 27);
		h += (int)s[i];
	}
	return (size_t)h;
}

#include "Hashtable.c"
#endif
