template <typename T>
struct Increase {
	virtual void operator() (T& e) {
		++e;
	}
};

template <typename T>
void increase(Vector<T>& V) {
	V.traverse(Increase<T>());
}
