template <typename T>
int disorder() const {		//check Vector whether sorted
	int n = 0;
	for (int i = 1; i < _size; ++i)
		if (_elem[i - 1] > _elem[i])
			++n;
	return n;
}
