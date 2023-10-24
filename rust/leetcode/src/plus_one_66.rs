/*
* 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

* 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

* 你可以假设除了整数 0 之外，这个整数不会以零开头。
*
*/
fn plus_one(digists: Vec<i32>) -> Vec<i32> {
    let mut digists = digists;
    for i in (0..digists.len()).rev() {
        if digists[i] == 9 {
            digists[i] = 0;
        } else {
            digists[i] += 1;
            return digists;
        }
    }
    digists.insert(0, 1);

    digists
}

#[test]
fn plus_one_test() {
    let tests = vec![
        (vec![1, 2, 3], vec![1, 2, 4]),
        (vec![4, 3, 2, 1], vec![4, 3, 2, 2]),
        (vec![0], vec![1]),
    ];

    for v in tests {
        assert_eq!(plus_one(v.0), v.1);
    }
}
