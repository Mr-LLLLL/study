template <typename T>
template <typename VST>
void traverse(VST& vist) {
	for (int i = 0; i < _size; ++i)
		visit(_elem[i]);
}
