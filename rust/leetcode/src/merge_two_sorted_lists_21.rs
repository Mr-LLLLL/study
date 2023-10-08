/*
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 */
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

fn merge_two(list1: Option<Box<ListNode>>, list2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    match (list1, list2) {
        (None, None) => None,
        (None, r) => r,
        (l, None) => l,
        (Some(mut l), Some(mut r)) => {
            if l.val <= r.val {
                l.next = merge_two(l.next, Some(r));
                Some(l)
            } else {
                r.next = merge_two(Some(l), r.next);
                Some(r)
            }
        }
    }
}

#[test]
fn merge_two_test() {
    let tests = vec![
        ((vec![1, 2, 4], vec![1, 3, 4]), vec![1, 1, 2, 3, 4, 4]),
        ((vec![], vec![]), vec![]),
        ((vec![], vec![0]), vec![0]),
        (
            (vec![1, 2, 3, 4, 5], vec![2, 3, 4, 9]),
            vec![1, 2, 2, 3, 3, 4, 4, 5, 9],
        ),
    ];

    'tag: for v in tests {
        let mut l1: Option<Box<ListNode>> = None;
        let mut l2: Option<Box<ListNode>> = None;
        let mut l3: Option<Box<ListNode>> = None;
        for vv in v.0 .0.into_iter().rev() {
            l1 = Some(Box::new(ListNode { val: vv, next: l1 }))
        }
        for vv in v.0 .1.into_iter().rev() {
            l2 = Some(Box::new(ListNode { val: vv, next: l2 }))
        }
        for vv in v.1.into_iter().rev() {
            l3 = Some(Box::new(ListNode { val: vv, next: l3 }))
        }

        let mut l4 = merge_two(l1, l2);

        while !l3.is_none() && !l4.is_none() {
            if l3.as_ref().unwrap().val != l4.as_ref().unwrap().val {
                assert!(false);
                continue 'tag;
            } else {
                l3 = l3.unwrap().next;
                l4 = l4.unwrap().next;
            }
        }

        if l3.is_none() && l4.is_none() {
            assert!(true);
        } else {
            assert!(false);
        }
    }
}
