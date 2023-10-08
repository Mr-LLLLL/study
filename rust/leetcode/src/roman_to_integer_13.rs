/*
 * 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
 *
 * 字符          数值
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 * 例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
 *
 * 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
 *
 * I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
 * X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
 * C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
 * 给定一个罗马数字，将其转换成整数。
 */
struct Solution {}

impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        s.chars()
            .fold((0, ' '), |res, ch| match (res.1, ch) {
                ('I', 'V') => (res.0 + 3, 'V'),
                ('I', 'X') => (res.0 + 8, 'X'),
                ('X', 'L') => (res.0 + 30, 'L'),
                ('X', 'C') => (res.0 + 80, 'C'),
                ('C', 'D') => (res.0 + 300, 'D'),
                ('C', 'M') => (res.0 + 800, 'M'),
                (_, 'I') => (res.0 + 1, 'I'),
                (_, 'V') => (res.0 + 5, 'V'),
                (_, 'X') => (res.0 + 10, 'X'),
                (_, 'L') => (res.0 + 50, 'L'),
                (_, 'C') => (res.0 + 100, 'C'),
                (_, 'D') => (res.0 + 500, 'D'),
                (_, 'M') => (res.0 + 1000, 'M'),
                (_, _) => unreachable!(),
            })
            .0
    }
}

#[test]
fn roman_to_int_test() {
    let tests = vec![
        (String::from("III"), 3),
        (String::from("IV"), 4),
        (String::from("IX"), 9),
        (String::from("LVIII"), 58),
        (String::from("MCMXCIV"), 1994),
    ];

    for v in tests {
        assert_eq!(Solution::roman_to_int(v.0), v.1);
    }
}
