#include <iostream>
#include <queue>
#include <utility>

using namespace std;

template <typename T>
struct TreeNode {
	T val;
	TreeNode *left;
	TreeNode* right;
	TreeNode(T x) : val(x), left(nullptr), right(nullptr) {}\
};

template <typename T>
using TreeNodePosi = TreeNode<T>*;

class Solution {
public:
	template <typename T>
	int maxDepth(TreeNodePosi<T> root) {
		queue< pair<TreeNodePosi<T>, int> > q;
		int maxDepth = 0;
		if (!root)
			return maxDepth;
		q.push(pair<TreeNodePosi<T>, int>(root, 1));
		while (!q.empty()) {
			TreeNodePosi<T> p = q.front().first;
			int currentDepth = q.front().second;
			q.pop();
			if (p->left)
				q.push(pair<TreeNodePosi<T>, int>(p->left, currentDepth + 1));
			if (p->right)
				q.push(pair<TreeNodePosi<T>, int>(p->right, currentDepth + 1));
			maxDepth = maxDepth < currentDepth ? currentDepth : maxDepth;
		}
		return maxDepth;
	}
};
