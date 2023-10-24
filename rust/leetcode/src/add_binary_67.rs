/*
 * 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
 */
fn add_binary(a: String, b: String) -> String {
    let (mut al, mut bl) = (a.len() as i32 - 1, b.len() as i32 - 1);
    let mut carry = 0;
    let mut res = String::from("");

    let a = a.as_bytes();
    let b = b.as_bytes();

    while al > -1 || bl > -1 {
        let num1 = a.get(al as usize).unwrap_or(&b'0') - b'0';
        let num2 = b.get(bl as usize).unwrap_or(&b'0') - b'0';

        let mut val = num1 + num2 + carry;
        carry = (val / 2) as u8;
        val = val % 2;
        res.insert(0, char::from(b'0' + val));

        al -= 1;
        bl -= 1;
    }

    if carry > 0 {
        res.insert(0, '1');
    }

    res
}

#[test]
fn add_binary_test() {
    let tests = vec![
        (("11".to_string(), "1".to_string()), "100".to_string()),
        (
            ("1010".to_string(), "1011".to_string()),
            "10101".to_string(),
        ),
    ];

    for v in tests {
        assert_eq!(add_binary(v.0 .0, v.0 .1), v.1);
    }
}
