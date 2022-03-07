#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
	void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
		size_t pos = m + n -1;
		while (n > 0)
			if (m > 0 && nums1[m - 1] > nums2[n - 1])
				nums1[pos--] = nums1[--m];
			else
				nums1[pos--] = nums2[--n];
	}
};



int main(int argc, char** argv)
{
	vector<int> nums1, nums2;
	int n;
	cout << "input first array please: " << endl;
	while (cin >> n, n)
		nums1.push_back(n);
	cout << "input second array please: " << endl;
	while (cin >> n, n)
		nums2.push_back(n);
	Solution().merge(nums1, nums1.size(), nums2, nums2.size());
	cout << nums1.size() << endl;
	n = 0;
	while (n < 10)
		cout << nums1[n++] << ' ';
	
	return 0;	
}

