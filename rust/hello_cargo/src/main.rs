fn main() {
    println!("Hello, world!");

    variable();

    tuple();

    array();

    let f = func(2, '2');

    println!("{f}");

    expr();
}

fn variable() {
    const X: u32 = 5;

    let y = 1;
    let mut y = y + 1;
    y = y + 1;

    let a = "hello";

    println!("{X}, {y}, {a}");
}

fn tuple() {
    let x = (500, 1.1, 2);

    let b = x.2;

    println!("{b}");
}

fn array() {
    let arr = [3; 5];

    let arr1 = arr[4];

    println!("{arr1}");
}

fn func(x: i32, label: char) -> i32 {
    println!("{x}, {label}");

    6
}

fn expr() {
    let y = {
        let x = 3;
        x + 1
    };

    println!("{y}");
}
