/*
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 */
fn climb_stairs(n: i32) -> i32 {
    let (mut g, mut f, mut n) = (0, 1, n);
    while 0 < n {
        f = g + f;
        g = f - g;
        n = n - 1;
    }

    return f;
}

#[test]
fn climb_stairs_test() {
    let tests = vec![(2, 2), (3, 3)];

    for v in tests {
        assert_eq!(climb_stairs(v.0), v.1);
    }
}
