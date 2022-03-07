#include <vector>

using namespace std;

// brude force
template <typename T>
T trivialMedian(vector<T>& s1, int lo1, int n1, vector<T>& s2, int lo2, int n2) {
	int hi1 = lo1 + n1;
	int hi2 = lo2 + n2;
	vector<T> s;
	while ((lo1 < hi1) && (lo2 < hi2)) {
		while (lo1 < hi1 && s1[lo1] <= s2[lo2])
			s.push_back(s1[lo1++]);
		while (lo2 < hi2 && s2[lo2] <= s1[lo1])
			s.push_back(s2[lo2++]);
	}
	while (lo1 < hi1)
		s.push_back(s1[lo1++]);
	while (lo2 < hi2)
		s.push_back(s2[lo2++]);
	return s[(n1 + n2) >> 1];
}

template <typename T>
T median(vector<T>& s1, int lo1, int n1, vector<T>& s2, int lo2, int n2) {
	// make n1 <= s2
	if (n1 > n2)
		return median(s2, lo2, n2, s1, lo1, n1);
	// if n2 < 6, may n1 < 3; so only brude force will do it
	if (n2 < 6)
		return trivialMedian(s1, lo1, n1, s2, lo2, n2);
	// sheer branch
	if (2 * n1 < n2)
		return median(s1, lo1, n1, s2, lo2 + (n2 - n1 - 1) >> 1, n1 + 2 - (n2 - n1) % 2);
	int mi1 = lo1 + n1 >> 1;
	int mi2a = lo2 + (n1 - 1) >> 1;
	int mi2b = lo2 + n2 - 1 - n1 >> 1;
	// retain left half of s1 and right half of s2;
	if (s1[mi1] > s2[mi2b])
		return median(s1, lo1, n1 >> 1 + 1, s2, mi2a, n2 - (n1 - 1) >> 1);
	// retain right half of s1 and left half of s2;
	else if (s1[mi1] < s2[mi2a])
		return median(s1, mi1, (n1 + 1) >> 1, s2, lo2, n2 - n1 >> 1);
	// retain s1 wholy and median partion of s2;
	else
		return median(s1, lo1, n1, s2, mi2a, n2 - (n1 - 1));
}

