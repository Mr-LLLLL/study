template <typename T>
ListNode<T>* List<T>::search(T const &e, int n, ListNode<T> *p) const {
	while (0 <= n--)
		if (((p = p->pred)->data) <= e)
			break;
	return p;
}
