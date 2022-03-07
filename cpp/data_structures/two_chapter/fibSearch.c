#include "fibonacci\Fib.h"

template <typename T>
static int fibSreach(T* A, T const& e, int lo, int hi) {
	Fib fib (hi - lo);
	while (lo < hi) {
		while (hi - lo < fib.get())
			fib.prev();
		int mi = lo + fib.get() - 1;
		if (e < A[mi])
			hi = mi;
		else if (A[mi] < e)
			lo = mi + 1;
		else
			return mi;
	}
	return -1;
}
	
