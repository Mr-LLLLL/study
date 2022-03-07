#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
	int singleNumber(vector<int>& nums) {
		int res = 0;
		for (int i : nums)
			res ^= i;
		return res;
	}
};

int main(int argc, char** argv)
{
	int n;
	while (cout << "input vector please:") {
		vector<int> v;
		while (cin >> n)
			v.push_back(n);
		cin.clear();
		cin.ignore();
		cout << "the result is: " << Solution().singleNumber(v) << endl;
	}

	return 0;
}
