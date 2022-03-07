template <typename T>
struct List {
	T data;
	ListNode<T> *pred;
	ListNode<T> *succ;
	ListNode() {}
	ListNode(T e, ListNode<T> p = null, ListNode<T> s = null) : data(e), pred(p), succ(s){}
};

template <typename T> class List {
private:
	int _size;
	ListNode<T> *header;
	LIstNode<T> *trailer;
protected:
	void init();
}:

template <typename T>
void List<T>::init() {
	header = new ListNode<T>;
	trailer = new ListNode<T>;
	header->succ = trailer;
	header->pred = nullptr;
	trailer->pred = header;
	trailer->succ = nullptr;
	_size = 0;
}
