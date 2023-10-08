/*
 * 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
 *
 * 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
 *
 * 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
 * 返回 k 。
 * 判题标准:
 *
 * 系统会用下面的代码来测试你的题解:
 *
 * int[] nums = [...]; // 输入数组
 * int[] expectedNums = [...]; // 长度正确的期望答案
 *
 * int k = removeDuplicates(nums); // 调用
 *
 * assert k == expectedNums.length;
 * for (int i = 0; i < k; i++) {
 *     assert nums[i] == expectedNums[i];
 * }
 * 如果所有断言都通过，那么您的题解将被 通过。
 */
fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
    let mut i = 0;
    for j in 1..nums.len() {
        if nums[j] != nums[i] {
            i = i + 1;
            nums[i] = nums[j];
        }
    }

    (i + 1) as i32
}

#[test]
fn remove_duplicates_test() {
    let tests = vec![(vec![1, 1, 2], 2), (vec![0, 0, 1, 1, 1, 2, 2, 3, 3, 4], 5)];

    for mut v in tests {
        assert_eq!(remove_duplicates(&mut v.0), v.1)
    }
}
