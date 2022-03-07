template <typename T>
ListNode<T>* insertAsFirst(T const& e) {
	++_size;
	return header->insertAsSucc(e);
}

template <typenmae T>
ListNode<T>* insertAsLast(T const& e) {
	++_size;
	return trailer->insertAsPred(e);
}

template <typename T>
ListNode<T>* insertAfter(ListNode<T>* p, T const& e) {
	++_size;
	return p->insertAsSucc(e);
}

template <typename T>
ListNode<T>* insertBefore(ListNode<T>* p, T const& e) {
	++_size;
	return p->insertAsPred(e);
}

template <typename T>
ListNode<T>* insertAsPred(T cosnt& e) {
	ListNode<T> x = new (e, pred, this);
	pred->succ = x;	
	pred = x;
	return x;
}

template <typename T>
ListNode<T>* insertAsSucc(T const& e) {
	ListNode<T> x = new (e, this, succ);
	succ = x;
	succ->pred = x;
	return x;
}
