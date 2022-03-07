// version 1
template <typename T>
void List<T>::insertionsort(ListNode<T> *p, int n) {
	for (int r = 0; r < n; ++r) {
		insertafter(search(p->data, r, p), p->data);
		p = p->succ;
		remove(p->pred);
	}
}

// version 2

template<typename T>
ListNode<T>* insetA(ListNode<T>* p, ListNode<T>* q) {
	if (q == p)
		return q;
	q->Pred->Succ = q->Succ;
	q->Succ->Pred = q->Pred;
	q->Pred = p;
	q->Succ = p->Succ;
	p->succ->Pred = q;
	p->Succ = q;
	return q;
}



template <typename T>
void List<T>::insertionsort(ListNode<T> *p, int n) {
	for (int r = 0; r < n; ++r) {
		ListNode<T> *temp = p;
		insertA(search(p->data, r, p), p);
		temp = temp->succ;
	}
}
