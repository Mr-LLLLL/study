template <typename T>
int deduplicate() {
	int oldsize = _size;
	int i = 1;
	while (i < _size)
		find(_elem[i], 0, i) < 0 ? ++i : remove(i);
	return oldsize - _size;
}

