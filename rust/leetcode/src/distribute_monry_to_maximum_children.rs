/*
 * 给你一个整数 money ，表示你总共有的钱数（单位为美元）和另一个整数 children ，表示你要将钱分配给多少个儿童。
 *
 * 你需要按照如下规则分配：
 *
 * 所有的钱都必须被分配。
 * 每个儿童至少获得 1 美元。
 * 没有人获得 4 美元。
 *
 * 请你按照上述规则分配金钱，并返回 最多 有多少个儿童获得 恰好 8 美元。如果没有任何分配方案，返回 -1 。
 *
 * 1 <= money <= 200
 * 2 <= children <= 30
 */
pub fn dist_money(money: i32, children: i32) -> i32 {
    if money < children {
        return -1;
    } else if money > 8 * children {
        return children - 1;
    } else if money == children * 8 - 4 {
        return children - 2;
    } else {
        return (money - children) / 7;
    }
}

#[test]
fn dist_money_test() {
    let tests = vec![((20, 3), 1), ((16, 2), 2)];

    for v in tests {
        assert_eq!(dist_money(v.0 .0, v.0 .1), v.1)
    }
}
