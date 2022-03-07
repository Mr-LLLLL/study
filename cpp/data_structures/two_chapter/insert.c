template <typename T>

template <typename T>
int insert(int r, T const& e) {
	expand();
	for (int i = _size; i > r; --i)
		_elem[i] = _elem[r];
	return r;
}
