template <typename T>
ListNode<T>* find(T const& e, int n, ListNode<T>* p) const {
	while (0 < n--)
		if (e == (p = p->pred)->data)
			return p;
	return nullptr;
}

template <typename T>
ListNode<T> *find(T const& e) const {
	return find(e, _size, trailer);
}
