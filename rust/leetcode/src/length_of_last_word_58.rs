/*
* 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

* 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
*/
fn length_of_last_word(s: String) -> i32 {
    s.into_bytes()
        .into_iter()
        .rev()
        .skip_while(|&c| c == b' ')
        .take_while(|&c| c != b' ')
        .count() as i32
}

#[test]
fn length_of_last_word_test() {
    let tests = vec![
        ("Hello World".to_string(), 5),
        ("   fly me   to   the moon  ".to_string(), 4),
        ("luffy is still joyboy".to_string(), 6),
    ];

    for v in tests {
        assert_eq!(length_of_last_word(v.0), v.1);
    }
}
