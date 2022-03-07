#include <iostream>
#include <queue>
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
	int minDepth(TreeNodePosi<T> root) {
		queue< pair<TreeNodePosi<T>, int> > q;
		int Depth = 0;
		if (!root)
			return Depth;
		q.push(make_paire(root, Depth + 1));
		while (q.empty()) {
			TreeNodePosi<T> node = q.front().first;
			Depth = q.front().second;
			q.pop();
			if (!node->left && !node->right)
				return Depth;
			if (node->left)
				q.push(make_pair(node->left, Depth + 1));
			if (node->right)
				q.push(make_pair(node->right, Depth + 1));
		}
	}
};
