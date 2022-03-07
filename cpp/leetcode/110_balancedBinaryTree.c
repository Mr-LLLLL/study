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
	bool isBalanced(TreeNodePosi<T> root) {
		int height;
		return isBalanced(root, height);
	}
	template <typename T>
	bool isBalanced(TreeNodePosi<T> p, int& height) {
		if (!p) {
			height = 0;
			return true;
		}
		int h1, h2;
		bool flag = isBalanced(p->left, h1) && isBalanced(p->right, h2);
		height = max(h1, h2) + 1;
		return flag && abs(h1 - h2) <= 1;
	}
};
