template <typename T>
int find(T const& e, int lo, int hi) const {
	while ((lo < hi--) && (e != _elem[hi]));
	return hi;
}
