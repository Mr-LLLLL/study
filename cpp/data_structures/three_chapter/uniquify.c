template <typename T>
int List<T>::uniquify() {
	if (_size < 2)
		return 0;
	int oldSize = _size;
	ListNode<T> *p = first();
	ListNode<T> *q;
	while (trailer != (q = p->scuu))
		if (p->data != q->data)
			p = q;
		else 
			remove(q);
	return oldSize - _size;
}
