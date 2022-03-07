template <typename T>
T& operator[] (int r) const {
	ListNode<T>* p = first();
	while (0 < r--)
		p = p->succ;;
	return p->data;
}
