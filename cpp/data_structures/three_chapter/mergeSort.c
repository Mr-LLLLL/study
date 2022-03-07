template <typename T>
void List<T>::merge(ListNode<T> *&p, int n, List<T>&L, ListNode<T> *q, int m) {
	ListNode<T> *pp = p->pred;
	while (0 < m)
		if ((0 < n) && (p->data <= q->data)) {
			if (q == (p = p->succ)) 
				break;
				--n;
		}
		else {
			insertBefore(p, L.remove((q = (q = q->succ)->pred));
			--m;
		}
	p = pp->succ;
}


template <typename T>
void List<T>::mergeSort(ListNode<T> *&p, int n) {
	if (n < 2)
	return;
	int m = n >> 1;
	ListNode<T> *q = p;
	for (int i = 0; i < m; ++i)
		q = q->succ;
	mergeSort(p, m);
	mergeSort(q, n - m);
	merge(p, m, *this, q, n - m);
}

