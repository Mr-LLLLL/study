template <typename T>
void unsort (int lo, int hi) {
	T* V = _elem = lo;
	for (int i = hi - lo; i > 0; --i)
		swap (V[i - 1], V[rand() % i]);
}
