template <typename T>
void List<T>::selectionSort(ListNode<T> *p, int n) {
	ListNode<T> *head = p->pred;
	ListNode<T> *tail = p;
	for (int i = 0; i < n; ++i)
		tail = tail->succ;
	while (1 < n) {
		ListNode<T> *max = selectMax(head->succ, n);
		swap(tail->pred->data, max->data);
		tail = tail-pred;
		--n;
	}
}

template <typename T>
ListNode<T>* List<T>::selectMax(ListNode<T> *p, int n) {
	ListNode<T> *max = p;
	for (ListNode<T> *cur = p; 1 < n; --n)
		if (!lt((cur = cur->succ)->data, max->data))
			max = cur;
	return max;
}
