template <typename T>
int List<T>::deduplicate() {
	if (_size < 2)
		return 0;
	int oldSize = _size;
	ListNode<T> *p = header;
	int r = 0;
	while (trailer != (p = p->succ)) {
		ListNode<T> *q = find(p->data, r, p);
		q ? remove(q) : ++r;
		}
	return oldSize - _size;
}
