#include <cstdlib>
#include <iostream>

template <typename T>
void partition(T* array, int lo, int hi);

template <typename T>
void quickSort(T*  array, int lo, int hi) {
	if (hi - lo < 2)
		return;
	int mi = partition(array, lo, hi - 1);
	quickSort(array, lo, mi);
	quickSort(array, mi + 1, hi);
}

// plan A, if array all of number is equal, than O will reach O(n^2)
// but the plan wii do less exchange
template <typename T>
void partition(T* array, int lo, int hi) {
	swap(array[lo], array[lo + rand() % (hi - lo + 1)]);
	T pivot = array[lo];
	while (lo < hi) {
		while (lo < hi && pivot <= array[hi])
			--hi;
		array[lo] = array[hi];
		while (lo < hi && array[lo] <= pivot)
			++lo;
		array[hi] = array[lo];
	}
	array[lo] = pivot;
	return lo;
}

// plan B will tackle all of repeated number
// but the plan will do more exchange
template <typename T>
void partitionB(T* array, int lo, int hi) {
	swap(array[lo], array[lo + rand() % (hi - lo + 1)]);
	T pivot = array[lo];
	while (lo < hi) {
		while (lo < hi)
			if (pivot < array[hi])
				--hi;
			else {
				array[lo++] = array[hi];
				break;
			}
		while (lo < hi)
			if (array[lo] < pivot)
				++lo;
			else {
				array[hi--] = array[lo];
				break;
			}
	}
	array[lo] = pivot;
	return lo;
}

// Plan C this version is same with Plan A
template <typename T>
int partitionC(T* array, int lo, int hi) {
	swap(array[0], array[lo + rand() % (hi - lo + 1)]);
	T pivot = array[0];
	int mi = lo;
	for (int k = lo + 1; k <= hi; ++k)
		if (array[k] < pivot)
			swap(array[++mi], array[k]);
	swap(array[lo], array[mi]);
	return mi;
}

