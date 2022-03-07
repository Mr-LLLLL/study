#include <iostream>
#include <stack>

using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

class Solution {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        stack<TreeNode*> s;
        while (true) {
            if (!compareAlongLeftBranch(p, q, s))
                return false;
            if (s.empty())
                return true;
            q = s.top();
            s.pop();
            p = s.top();
            s.pop();
        }
    }

    bool compareAlongLeftBranch(TreeNode* p, TreeNode* q, stack<TreeNode*>& s) {
        while (p && q) {
            if (p->val != q->val)
                return false;
            s.push(p->right);
            s.push(q->right);
            p = p->left;
            q = q->left;
        }
        if (p != q)
            return false;
        return true;
    }
};

int main(int argc, char** argv)
{
    return 0;
}
