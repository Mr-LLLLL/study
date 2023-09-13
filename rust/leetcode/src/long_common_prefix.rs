/*
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 *
 *
 * 示例 1：
 *
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 * 示例 2：
 *
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 */
pub fn longest_common_prefix(strs: Vec<String>) -> String {
    let mut ans = 0;
    for i in 0..strs[0].len() {
        for s in strs.iter() {
            let s = s.as_bytes();

            if s.len() == i || strs[0].as_bytes()[i] != s[i] {
                return strs[0][0..ans].to_string();
            }
        }

        ans += 1;
    }

    strs[0][0..ans].to_string()
}

#[test]
fn longest_common_prefix_test() {
    let tests = vec![
        (
            vec![
                "flower".to_string(),
                "flow".to_string(),
                "flight".to_string(),
            ],
            "fl".to_string(),
        ),
        (
            vec!["dog".to_string(), "racecar".to_string(), "car".to_string()],
            "".to_string(),
        ),
    ];

    for v in tests {
        assert_eq!(longest_common_prefix(v.0), v.1);
    }
}
