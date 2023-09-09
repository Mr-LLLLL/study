/*
 * 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
 *
 * 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 例如，121 是回文，而 123 不是。
 */
pub fn is_palindrome(x: i32) -> bool {
    if x < 0 {
        false
    } else {
        let mut tmp = 0;
        let mut y = x;
        while y > 0 {
            tmp = tmp * 10 + y % 10;

            y /= 10;
        }

        tmp == x
    }
}

#[test]
fn is_palindrome_test() {
    let tests = vec![(121, true), (-121, false), (10, false)];

    for v in tests {
        assert_eq!(is_palindrome(v.0), v.1);
    }
}
