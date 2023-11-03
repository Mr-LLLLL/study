/*
 * 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
 */
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32, next: Option<Box<ListNode>>) -> Self {
        ListNode { next, val }
    }
}

fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut head = head;
    let mut cur = head.as_mut();
    while cur.is_some() && cur.as_ref().unwrap().next.is_some() {
        if cur.as_ref().unwrap().val == cur.as_ref().unwrap().next.as_ref().unwrap().val {
            let next = cur.as_mut().unwrap().next.as_mut().unwrap().next.take();
            cur.as_mut().unwrap().next = next;
        } else {
            cur = cur.unwrap().next.as_mut();
        }
    }

    head
}

#[test]
fn delete_duplicates_test() {
    let tests = vec![
        (vec![1, 1, 2], vec![1, 2]),
        (vec![1, 2, 2, 3, 3], vec![1, 2, 3]),
    ];

    'tag: for v in tests {
        let mut input: Option<Box<ListNode>> = None;
        let mut want: Option<Box<ListNode>> = None;
        for vv in v.0.into_iter().rev() {
            input = Some(Box::new(ListNode::new(vv, input)))
        }
        for vv in v.1.into_iter().rev() {
            want = Some(Box::new(ListNode::new(vv, want)))
        }
        let mut res = delete_duplicates(input);

        while !want.is_none() && !res.is_none() {
            if want.as_ref().unwrap().val != res.as_ref().unwrap().val {
                assert!(false);
                continue 'tag;
            } else {
                want = want.unwrap().next;
                res = res.unwrap().next;
            }
        }

        if res.is_none() && want.is_none() {
            assert!(true)
        } else {
            assert!(false)
        }
    }
}
