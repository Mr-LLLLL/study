#include <iostream>
#include <vector>

using namespace std;

template <typename T>
struct TreeNode {
	T val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(T x) : val(x), left(nullptr), right(nullptr) {}\
};

template <typename T>
using TreeNodePosi = TreeNode<T>*;

class Solution {
public:
	template <typename T>
	TreeNodePosi<T> sortedArrayToBST(vector<T> const& nums) {
		return midValueInsert(0, nums.size(), nums);
	}
	template <typename T>
	TreeNodePosi<T> midValueInsert(int lo, int hi, vector<T> const& nums) {
		if (lo >= hi)
			return nullptr;
		int mi = (lo + hi) >> 1;
		TreeNodePosi<T> p = new TreeNode<T>(nums[mi]);
		p->left = midValueInsert(lo, mi, nums);
		p->right = midValueInsert(mi + 1, hi, nums);
		return p;
	}
};

