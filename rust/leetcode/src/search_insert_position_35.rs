/*
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 请必须使用时间复杂度为 O(log n) 的算法。
 */
fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
    let mut max = nums.len();
    let mut min = 0;
    while min < max {
        let mid = (max + min) >> 1;
        if target > nums[mid] {
            min = mid + 1
        } else {
            max = mid
        }
    }

    max as i32
}

#[test]
fn search_insert_test() {
    let tests = vec![
        ((vec![1, 3, 5, 6], 5), 2),
        ((vec![1, 3, 5, 6], 2), 1),
        ((vec![1, 3, 5, 6], 7), 4),
    ];

    for v in tests {
        assert_eq!(search_insert(v.0 .0, v.0 .1), v.1);
    }
}
