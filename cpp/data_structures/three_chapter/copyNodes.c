template <typename T>
void copyNodes(ListNode<T>* p, int n) {
	init();
	while(n--) {
		insertLast(p->data);
		p = p->succ;
	}
}

template <typename T>
List<T>::List(ListNode<T>* p, int n) {
	copyNodes(p, n);
}

template <typename T>
List<T>::List(List<T> const& L) {
	copyNodes(L.first(), L._size);
}

template <typename T>
List<T>::List(List<T> const& L, int r, int n) {
	copyNodes(L[r], n);
}
