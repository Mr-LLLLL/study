#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
	int removeDuplicates(vector<int>& nums) {
		if (!nums.size())
			return 0;
		size_t i = 0;
		for (size_t j = 0; j < nums.size(); ++j) {
			if (nums[j] != nums[i]) {
				++i;
				nums[i] = nums[j];
			}
		}
		return 1 + i;
	}
};
