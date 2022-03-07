#include <vector>

class Solution {
public:
	size_type removeElement(vector<int>& nums, int val) {
		size_type i = 0;
		size_type n = nums.size();
		while (i < n) {
			if (nums[i] == val) {
				nums[i] = nums[n - 1];
				--n;
			} else {
				++i;
			}
		}
	}
};
