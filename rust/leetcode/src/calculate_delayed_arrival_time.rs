pub fn find_delayed_arrival_time(arrival_time: i32, delayed_time: i32) -> i32 {
    (arrival_time + delayed_time) % 24
}

#[test]
fn find_delayed_arrival_time_test() {
    let tests = vec![((15, 5), 20), ((13, 11), 0)];

    for v in tests {
        assert_eq!(find_delayed_arrival_time(v.0 .0, v.0 .1), v.1);
    }
}
