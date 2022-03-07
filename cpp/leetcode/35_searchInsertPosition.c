class Solution {
public:
	int searchInsert(vector<int>& nums, int target) {
		size_type max = nums.size() - 1;
		if (target > nums[max])
			return max + 1;
		size_type min = 0;
		size_type mi = (max + min) >> 1;
		while (max - min) {
			if (target <= nums[mi])
				max = mi;
			else
				min = mi + 1;
			mi = (max + min) >> 1;
		}

		return max;
	}
};
