#include <iostream>
#include <stack>
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
	bool hasPathSum(TreeNodePosi<T> root, T sum) {
		if (!root)
			return false;
		stack<TreeNodePosi<T>> node_stack;
		stack<T> t_stack;
		node_stack.push(root);
		t_stack.push(sum - root->val);
		while (!node_stack.empty()) {
			TreeNodePosi<T> node = node_stack.top();
			int curr_sum = t_stack.top();
			node_stack.pop();
			t_stack.pop();
			if (!node->right && !node->left && 0 == curr_sum)
				return true;
			if (node->right) {
				node_stack.push(node->right);
				t_stack.push(curr_sum - node->right->val);
			}
			if (node->left) {
				node_stack.push(node->left);
				t_stack.push(curr_sum - node->left->val);
			}
		}
		return false;
	}
};
