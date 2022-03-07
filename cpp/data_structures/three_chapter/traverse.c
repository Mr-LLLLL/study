template <typename T> template <typename VST>
void List<T>::traverse(VST &visit) {
	for (LIstNode<T> *p = header->succ; p !=trailer; p = p->succ)
		visit(p->data);
}
