#include <iostream>
#include <vector>
#include <stack>

using namespace std;

class Solution {
public:
	bool stackPermutation(vector<int> const& v, vector<int> const& after) {
		stack<int> s;
		for (size_t i = 0, j = 0; i < v.size(); ++i) {
			while (s.empty() || after[i] != s.top()) {
				if (j >= v.size())
					return false;
				s.push(v[j++]);
			}
			s.pop();
		}
		return true;
	}
};

int main(int argc, char** argv)
{
	vector<int> original, after;
	int n;
	cout << "please input original vector:";
	while (n != 0) {
		cin >> n;
		original.push_back(n);
	}
	n = 1;
	cout << "please input agter vector:";
	while (n != 0) {
		cin >> n;
		after.push_back(n);
	}
	cout << "the after vector is whether the original's stackpermutation: " << Solution().stackPermutation(original, after) << endl;



	return 0;
}
