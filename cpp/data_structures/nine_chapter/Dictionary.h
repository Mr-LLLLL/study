#ifndef DICTIONARY_H
#define DICTIONARY_H
template <typename K, typename V>
class Dictionary {
public:
	virtual int size() const = 0;
	virtual bool put(K, V) = 0;
	virtual V* get(K k) = 0;
	virtual bool remove(K k) = 0;
};
#endif
