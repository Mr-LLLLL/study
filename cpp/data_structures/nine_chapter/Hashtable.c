#include "Hashtable.h"

int primeNLT(int c, int n, char* file) {
	Bitmap B(file, n);
	while (c < n)
		if (B.test(c)) c++;
		else return c;
	return c;
}

template <typename K, typename V>
Hashtable<K, V>::Hashtable(int c) {
	M = primeNLT(c, 1048576, "primeBitmap.txt");
	N = 0;
	ht = new Entry<K, V>*[M];
	memset(ht, 0, sizeof(Entry<K, V>*) * M);
	lazyRemoval = new Bitmap(M);
}

template <typename K, typename V>
Hashtable<K, V>::~Hashtable() {
	for (int i = 0; i < M; ++i)
		if (ht[i])
			delete ht[i];
	delete ht;
	delete lazyRemoval;
	lazyRemoval = nullptr;

}

template <typename K, typename V>
V* Hashtable<K, V>::get(K k) {
	int r = probe4Hit(k);
	return ht[r] ? &(ht[r]->value) : nullptr;
}

template <typename K, typename V>
int Hashtable<K, V>::probe4Hit(const K& k) {
	int r = hashCode(k) % M;
	while ((ht[r] && (k != ht[r]->key)) || (!ht[r] && lazilyRemoved(r)))
		r = (r + 1) % M;
	return r;
}

template <typename K, typename V>
bool Hashtable<K, V>::remove(K k) {
	int r = probe4Hit(k);
	if (!ht[r])
		return false;
	delete ht[r];
	ht[r] = nullptr;
	markAsRemoved(r);
	--N;
	return true;
}

template <typename K, typename V>
bool Hashtable<K, V>::put(K k, V v) {
	if (ht[probe4Hit(k)])
		return false;
	int r = probe4Free(k);
	ht[r] = new Entry<K, V>(k, v);
	++N;
	if (N * 2 > M)
		rehash();
	return true;
}

template <typename K, typename V>
int Hashtable<K, V>::probe4Free(const K& k) {
	int r = hashCode(k) % M;
	while (ht[r])
		r = (r + 1) % M;
	return r;
}

template <typename K, typename V>
void Hashtable<K, V>::rehash() {
	int old_capacity = M;
	Entry<K, V>** old_ht = ht;
	M = primeNLT(2 * M, 1048576, "primeBitmap.txt");	//minimum prime greater c
	N = 0;
	ht = new Entry<K, V>* [M];
	memset(ht, 0, sizeof(Entry<K, V>*) * M);
	delete lazyRemoval;
	lazyRemoval = new Bitmap(M);
	for (int i = 0; i < old_capacity; ++i)
		if (old_ht[i]) {
			put(old_ht[i]->key, old_ht[i]->value);
			delete old_ht[i];
		}
	delete old_ht;
}
