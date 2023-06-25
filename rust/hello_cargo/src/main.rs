#[derive(Debug)]
struct User {
    active: bool,
    username: String,
    email: String,
    sign_in_count: u64,
}

fn build_user(email: String, username: String) -> User {
    User {
        active: true,
        username,
        email,
        sign_in_count: 1,
    }
}

struct Color(i32, i32, i32);

#[derive(Debug)]
struct AlwaysEqual;

fn main() {
    println!("Hello, world!");

    variable();

    tuple();

    array();

    let f = func(2, '2');

    println!("{f}");

    expr();

    circle();

    owner();

    slice();

    let user1 = User {
        active: true,
        username: String::from("name"),
        email: String::from("@"),
        sign_in_count: 1,
    };

    let user2 = build_user(String::from("email"), String::from("username"));

    let user3 = User {
        email: String::from("eamil"),
        ..user1
    };

    println!("{:?}", user3);

    println!(
        "{}, {}, {}, {}",
        user2.username, user2.email, user2.sign_in_count, user2.active
    );

    let black = Color(0, 0, 0);
    println!("{}", black.0);

    let subject = AlwaysEqual;
    println!("{:?}", subject);

    dbg();

    println!("{}", area((30, 50)));
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

fn circle() {
    let mut s = 'tag: loop {
        loop {
            println!("loop in");
            break 'tag 2;
        }
    };

    println!("{s}");

    while s < 3 {
        s = s + 1;
        println!("{s}");
    }

    for n in (1..4).rev() {
        println!("{n}");
    }
}

fn owner() {
    let s = String::from("hello");
    let s1 = s;

    println!("{}", s1);

    let s2 = s1.clone();
    println!("{}, {}", s2, s1);

    let s3 = String::from("hello");
    takes_ownership(s3, String::from("wolrd"));

    makes_copy(5);

    let s4 = gives_ownership();
    let s5 = String::from("hello");
    let s6 = takes_and_gives_back(s5);

    println!("{}, {}", s4, s6);

    println!("{}", calculate_length(&s6));

    let mut s7 = String::from("hello");
    change(&mut s7);

    println!("{}", s7);

    let mut s8 = String::from("hello");
    let s9 = &s8;
    let s10 = &s8;
    println!("{}, {}", s9, s10);

    let s11 = &mut s8;
    println!("{}", s11);
}

fn takes_ownership(s: String, s1: String) {
    println!("{}, {}", s, s1);
}

fn makes_copy(i: i32) {
    println!("{}", i);
}

fn gives_ownership() -> String {
    let s = String::from("yours");

    s
}

fn takes_and_gives_back(a_string: String) -> String {
    a_string
}

fn calculate_length(s: &String) -> usize {
    s.len()
}

fn change(s: &mut String) {
    s.push_str(",wolrd");
}

fn slice() {
    let s = String::from("Hello world");

    let hello = &s[..5];
    let world = &s[6..];

    println!("{}, {}", hello, world);
}

fn dbg() {
    let user = User {
        username: dbg!(String::from("name")),
        active: true,
        email: String::from("email"),
        sign_in_count: 1,
    };

    dbg!(&user);
}

fn area(dimensions: (u32, u32)) -> u32 {
    dimensions.0 * dimensions.1
}
