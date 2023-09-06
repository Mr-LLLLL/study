use std::{collections::HashMap, vec};

fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut map = HashMap::with_capacity(nums.len());

    for i in 0..nums.len() {
        if let Some(k) = map.get(&(target - nums[i])) {
            if *k != i {
                return vec![*k as i32, i as i32];
            }
        }
        map.insert(nums[i], i);
    }
    panic!("not found")
}

#[test]
fn two_sum_test() {
    let tests = vec![
        ((vec![2, 7, 11, 15], 9), vec![0, 1]),
        ((vec![3, 3], 6), vec![0, 1]),
        ((vec![3, 2, 4], 6), vec![1, 2]),
    ];

    for v in tests {
        assert_eq!(two_sum(v.0 .0, v.0 .1), v.1);
    }
}
