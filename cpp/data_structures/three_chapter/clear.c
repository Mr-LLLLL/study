template <typename T>
List<T>::~List() {
	clear();
	delete header;
	delete trailer;
}

template <typename T>
int List<T>::clear() {
	int oldsize = _size;
	while (0 < _size)
		remove (header->succ);
	return oldsize;
}
