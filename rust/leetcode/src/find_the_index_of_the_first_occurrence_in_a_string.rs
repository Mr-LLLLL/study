use std::{mem::needs_drop, ops::Index};

/*
 * 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：haystack = "sadbutsad", needle = "sad"
 * 输出：0
 * 解释："sad" 在下标 0 和 6 处匹配。
 * 第一个匹配项的下标是 0 ，所以返回 0 。
 * 示例 2：
 *
 * 输入：haystack = "leetcode", needle = "leeto"
 * 输出：-1
 * 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
 */
fn str_str(haystack: String, needle: String) -> i32 {
    let mut i = 0;
    let mut j = 0;
    let hb = haystack.as_bytes();
    let nb = needle.as_bytes();
    while j + i < hb.len() && i < nb.len() {
        if hb[j + i] == nb[i] {
            i += 1;
        } else {
            i = 0;
            j += 1;
        }
    }

    if i == nb.len() {
        j as i32
    } else {
        -1
    }
}

#[test]
fn str_str_test() {
    let tests = vec![
        (("a".to_string(), "a".to_string()), 0),
        (("sadbutsad".to_string(), "sad".to_string()), 0),
        (("leetcode".to_string(), "leeto".to_string()), -1),
        (("aaa".to_string(), "aaaa".to_string()), -1),
        (("mississippi".to_string(), "issipi".to_string()), -1),
    ];

    for v in tests {
        assert_eq!(str_str(v.0 .0, v.0 .1), v.1);
    }
}
