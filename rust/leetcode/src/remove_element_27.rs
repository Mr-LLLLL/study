use crate::palindrome_number_9::is_palindrome;

/*
 * 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
 *
 * 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
 *
 * 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
 *
 *
 *
 * 说明:
 *
 * 为什么返回数值是整数，但输出的答案是数组呢?
 *
 * 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
 *
 * 你可以想象内部操作如下:
 *
 * // nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
 * int len = removeElement(nums, val);
 *
 * // 在函数里修改输入数组对于调用者是可见的。
 * // 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
 * for (int i = 0; i < len; i++) {
 *     print(nums[i]);
 * }
 */
fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
    let mut index = 0;

    for i in 0..nums.len() {
        if nums[i] != val {
            nums[index] = nums[i];
            index += 1;
        }
    }

    index as i32
}

#[test]
fn remove_element_test() {
    let tests = vec![
        ((vec![3, 2, 2, 3], 3), 2),
        ((vec![0, 1, 2, 2, 3, 0, 4, 2], 2), 5),
    ];

    for mut v in tests {
        assert_eq!(remove_element(&mut v.0 .0, v.0 .1), v.1)
    }
}
