/* 给定一个  无重复元素 的 有序 整数数组 nums 。
 *
 *  返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
 *
 *  列表中的每个区间范围 [a,b] 应该按如下格式输出：
 *
 *  "a->b" ，如果 a != b
 *  "a" ，如果 a == b
 *
 * NOTE:
 * 0 <= nums.length <= 20
 * -231 <= nums[i] <= 231 - 1
 * nums 中的所有值都 互不相同
 * nums 按升序排列
*/
pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
    let mut res = vec![];
    let mut i = 0;

    while i < nums.len() {
        let mut j = i + 1;
        while j < nums.len() && nums[j - 1] + 1 == nums[j] {
            j += 1;
        }

        if j - 1 == i {
            res.push(nums[i].to_string());
        } else {
            res.push(format!("{}->{}", nums[i], nums[j - 1]));
        }
        i = j;
    }

    res
}

#[test]
fn summary_ranges_test() {
    let tests = vec![
        (vec![0, 1, 2, 4, 5, 7], vec!["0->2", "4->5", "7"]),
        (vec![0, 2, 3, 4, 6, 8, 9], vec!["0", "2->4", "6", "8->9"]),
    ];

    for v in tests {
        assert_eq!(summary_ranges(v.0), v.1);
    }
}
