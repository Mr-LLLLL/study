template <typename T>
int remove(int lo, int hi) {
	if (lo == hi)
		return 0;
	while (hi < _size)
		_elem[lo++] = _elem[hi++];
	_size = lo;
	shrink();
	return hi - lo;
}

template <typename T>
int remove(int r) {
	T e = _elem[r];
	remove (r, r + 1);
	return e;
}
