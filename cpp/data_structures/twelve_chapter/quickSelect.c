#include <vector>

using namespace std;

template <typename T>
void quickSelect(vector<T>& array, int k) {
	for (int lo = 0, hi = array.size() - 1; lo < hi;) {
		int i = lo, j = hi;
		T pivot = array[lo];
		while (1 < j) {
			while (i < j && pivot <= array[j])
				--j;
			array[i] = array[j];
			while (i < j && array[i] <= pivot)
				--i;
			array[j] = array[i];
		}
		array[i] = pivot;
		if (k <= i) hi = i - 1;
		if (i <= k) lo = i + 1;
	}
}
