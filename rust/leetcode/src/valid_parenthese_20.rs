/*
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 每个右括号都有一个对应的相同类型的左括号。
 */
struct Solution {}

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack = vec![' '];
        for c in s.chars() {
            match c {
                '(' | '[' | '{' => stack.push(c),
                ')' => {
                    if stack.pop().unwrap() != '(' {
                        return false;
                    }
                }
                ']' => {
                    if stack.pop().unwrap() != '[' {
                        return false;
                    }
                }
                '}' => {
                    if stack.pop().unwrap() != '{' {
                        return false;
                    }
                }
                _ => (),
            }
        }

        stack.len() == 1
    }
}

#[test]
fn is_valid_test() {
    let tests = vec![
        ("()".to_string(), true),
        ("()[]{}".to_string(), true),
        ("(]".to_string(), false),
    ];

    for v in tests {
        assert_eq!(Solution::is_valid(v.0), v.1);
    }
}
