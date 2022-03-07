template <typename T>
void mergeSort(int lo, int hi) {
	if (hi - lo < 2)
		return;
	int mi = (lo + hi) >> 1;
	mergeSort(lo, mi);
	mergeSort(mi, hi);
	if (_elem[mi - 1] > _elem[mi])
	merge(lo, mi, hi);
}

template <typename T>
void merge(int lo, int mi, int hi) {
	T* A = _elem + lo;
	int lb = mi - lo;
	T* B = new T[lb];
	for (int i = 0; i < lb; B[i] = A[i++]);
	int lc = hi - mi;
	T* C = _elem + mi;
	for (int i = 0, j = 0, k = 0; j < lb) {
		if (k == lc || B[j] <= C[k])
			A[i++] = B[j++];
		else
			A[i++] = C[k++];
	}
	delete [] B;
}
