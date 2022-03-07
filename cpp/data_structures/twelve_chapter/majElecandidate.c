#include <vector>

using namespace std;

template <typename T>
T majEleCandidate(vector<T> v) {
	T maj;
	for (int c = 0, i = 0; i < v.size(); ++i)
		if (0 == c){
			maj = v[i];
			c = 1;
		} else
			maj == v[i] ? ++c : --c;
	return maj;
}
