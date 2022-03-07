#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
	int maxSubArray(vector<int> const& nums) {
		int ret = nums[0];
		int sum = 0;
		for (int num : nums) {
			if (sum > 0)
				sum += num;
			else
				sum = num;
			ret = ret > sum ? ret : sum;
		}
		return ret;
	}
};

int main(int argc, char** argv)
{
	int n;
	vector<int> vec;
	cout << "please input vector<int>:";
	while (cin >> n)
		vec.push_back(n);
	cout << Solution().maxSubArray(vec) << endl;

	return 0;
}
