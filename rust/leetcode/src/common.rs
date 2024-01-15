use std::cell::RefCell;
use std::collections::LinkedList;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    val: i32,
    left: Option<Rc<RefCell<TreeNode>>>,
    right: Option<Rc<RefCell<TreeNode>>>,
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

    #[inline]
    pub fn set_left(&mut self, node: Option<Rc<RefCell<TreeNode>>>) -> &Self {
        self.left = node;

        self
    }

    #[inline]
    pub fn set_right(&mut self, node: Option<Rc<RefCell<TreeNode>>>) -> &Self {
        self.left = node;

        self
    }
}

pub fn vec_2_tree(vec: Vec<Option<i32>>) -> Rc<RefCell<TreeNode>> {
    let mut queue: LinkedList<Rc<RefCell<TreeNode>>> = LinkedList::new();
    let root = Rc::new(RefCell::new(TreeNode::new(vec[0].unwrap())));
    queue.push_back(root.clone());

    let mut i = 0;
    loop {
        let old_node = queue.pop_front().unwrap();

        i = i + 1;
        if i == vec.len() {
            break;
        }

        if let Some(v) = vec[i] {
            let new_node = Rc::new(RefCell::new(TreeNode::new(v)));
            queue.push_back(new_node.clone());
            old_node.borrow_mut().set_left(Some(new_node.clone()));
        }

        i = i + 1;
        if i == vec.len() {
            break;
        }

        if let Some(v) = vec[i] {
            let new_node = Rc::new(RefCell::new(TreeNode::new(v)));
            queue.push_back(new_node.clone());
            old_node.borrow_mut().set_right(Some(new_node.clone()));
        }
    }

    return root;
}

#[derive(Debug)]
struct Queue<T> {
    qdata: Vec<T>,
}

impl<T> Queue<T> {
    fn new() -> Self {
        Queue { qdata: Vec::new() }
    }

    fn push(&mut self, item: T) {
        self.qdata.push(item);
    }

    fn pop(&mut self) -> T {
        self.qdata.remove(0)
    }
}
