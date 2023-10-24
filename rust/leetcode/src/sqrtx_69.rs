/*
 * 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
 *
 * 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
 *
 * 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
 */
fn my_sqrt(x: i32) -> i32 {
    let x_half = 0.5f64 * x as f64;
    let mut i = (x as f64).to_bits();
    // magic number
    i = 0x5FE6EC85E7DE30DA_u64 - (i >> 1);
    let mut f = f64::from_bits(i);
    f = f * (1.5 - x_half * f * f);
    f = f * (1.5 - x_half * f * f);
    f = f * (1.5 - x_half * f * f);
    (1.0f64 / f) as i32
}

#[test]
fn my_sqrt_test() {
    let tests = vec![(4, 2), (8, 2)];

    for v in tests {
        assert_eq!(my_sqrt(v.0), v.1);
    }
}
