/*
 * 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
 */
// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

use std::borrow::BorrowMut;
use std::cell::RefCell;
use std::rc::Rc;
fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
    let mut ans = vec![];
    let mut stack = vec![];
    let mut node = root;

    while node.is_some() || stack.len() > 0 {
        while let Some(n) = node {
            node = n.borrow().left.clone();
            stack.push(n);
        }

        if let Some(n) = stack.pop() {
            ans.push(n.borrow().val);
            node = n.borrow().right.clone();
        }
    }

    ans
}

#[test]
fn inorder_traversal_test() {
    let tests = vec![
        (
            Some(Rc::new(RefCell::new(TreeNode {
                val: 1,
                left: None,
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 2,
                    left: Some(Rc::new(RefCell::new(TreeNode {
                        val: 3,
                        left: None,
                        right: None,
                    }))),
                    right: None,
                }))),
            }))),
            vec![1, 3, 2],
        ),
        (
            Some(Rc::new(RefCell::new(TreeNode {
                val: 1,
                left: None,
                right: None,
            }))),
            vec![1],
        ),
    ];

    for v in tests {
        assert_eq!(inorder_traversal(v.0), v.1);
    }
}
