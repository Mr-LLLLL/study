use hello_cargo;

mod common;

#[test]
fn it_adds_two() {
    common::setup();
    assert_eq!(4, hello_cargo::add_two(2))
}
