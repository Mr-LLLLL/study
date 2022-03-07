#include <iostream>
#include <vector>
#include <queue>

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
	vector<vector<T>> levelOrderBottom(TreeNodePosi<T> root) {
		vector<vector<T>> res;
		if (!root)
			return res;
		queue<TreeNodePosi<T>> q;
		q.push(root);
		while (!q.empty()) {
			vector<int> temp;
			int len = q.size();
			for (int i = 0; i < len; ++i) {
				TreeNodePosi<T> now = q.front();
				q.pop();
				temp.push_back(now->val);
				if (now->left)
					q.push(now->left);
				if (now->rigth)
					q,push(now->right);
			}
			res.insert(res.begin(), temp);
		}
		return res;
	}
};

