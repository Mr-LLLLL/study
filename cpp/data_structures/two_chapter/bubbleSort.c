template <typename T>
void bubblesort(int lo, int hi) {
	while ((lo = bubbleMin(lo, hi)) < (hi = bubble(lo, hi)));
}

template <typename T>
int bubbleMax(int lo, int hi) {
	int last = lo;
	while (++lo < hi)
		if (_elem[lo - 1] > _elem[lo]) {
			last = lo;
			swap (_elem[lo - 1], _elem[lo]);
		}
	return last;
}
template <typename T>
int bubbleMin(int lo, int hi) {
	int firts = hi;
	while (lo < --hi)
		if (_elem[hi] < _elem[hi - 1]) {
			first = hi;
			swap (_elem[hi], _elem[hi - 1]);
		}
	return first;
}
