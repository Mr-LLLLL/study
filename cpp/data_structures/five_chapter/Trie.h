#ifndef TRIE_H
#define TRIE_H

#include <string>

using std::string;

class TrieNode {
	friend class Trie;
private:
	bool isEnd;
	TrieNode* next[26];
public:
	TrieNode() : isEnd(false) {}
	bool isWord() {
		return isEnd;
	}
	void setWord() {
		isEnd = true;
	}
};

class Trie {
private:
	TrieNode* _root;
public:
	Trie() : _root(nullptr) {}
	bool searchPrefix(string const& word) {
		if (0 == word.size())
			return true;
		if (nullptr == _root)
			return false;
		TrieNode* node = _root;
		for (char c : word) {
			int i = c - 'a';
			if (nullptr == node->next[i])
				return false;
			else
				node = node->next[i];
		}
		return true;
	}
	bool search(string const& word) {
		if (0 == word.size())
			return true;
		if (nullptr == _root)
			return false;
		TrieNode* node = _root;
		for (char c : word) {
			int i = c - 'a';
			if (nullptr == node->next[i])
				return false;
			else
				node = node->next[i];
		}
		return node->isWord();
	}
	void insert(string const& word) {
		if (0 == word.size())
			return;
		if (nullptr == _root)
			_root = new TrieNode();
		TrieNode* node = _root;
		for (char c : word) {
			int i = c - 'a';
			if (nullptr == node->next[i])
				node->next[i] = new TrieNode();
			node = node->next[i];
		}
		node->setWord();
	}
};


#endif
