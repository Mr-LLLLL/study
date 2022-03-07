template <typename T>
Vector<T>& operator= (Vector<T> const& v) {
	if (_elem)
		delete [] _elem;
	copyFrom (v._elem, 0, v.size());
	return *this;
}

template <typename T>
T& operator[] (int r) const {
	return _elem[r];
}
