mod back_of_house;
mod front_of_house;

use std::fmt::{Display, Result};
use std::fs::{self, File};
use std::io::{self, Read, Write};
use std::io::{ErrorKind, Result as IoResult};
use std::{collections::*, panic, result, vec};

fn deliver_order() {}

pub use crate::front_of_house::hosting;

pub fn eat_at_restaurant() {
    hosting::add_to_awatlist();
}

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

#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.width
    }
}

impl Rectangle {
    fn square(size: u32) -> Self {
        Self {
            width: size,
            height: size,
        }
    }
}

#[derive(Debug)]
enum IpAddrKind {
    V4,
    V6,
}

#[derive(Debug)]
enum IpAddr {
    V4(u8, u8, u8, u8),
    V6(String),
}

enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

impl Message {
    fn call(&self) {
        println!("hello")
    }
}

// enum Option<T> {
//    ONone,
//     Some(T),
// }

#[derive(Debug)]
enum UsState {
    Alabama,
    Alaska,
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quater(UsState),
}

fn base_practice() {
    deliver_order();

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
    let rect1 = Rectangle {
        width: 40,
        height: 50,
    };

    let rect2 = Rectangle::square(30);

    println!("{}", rect1.area());
    println!("rect1 include rect2 {}", rect1.can_hold(&rect2));
    dbg!(&rect1);
    println!("rect1 is {:?}", rect1);

    let four = IpAddrKind::V4;
    let six = IpAddrKind::V6;

    let home = IpAddr::V4(1, 2, 3, 4);

    let loopback = IpAddr::V6(String::from("::1"));

    println!("{:?}, {:?}, {:?}, {:?}", four, six, home, loopback);

    let m = Message::Write(String::from("hello"));
    m.call();

    let x: i8 = 5;
    let y: Option<i8> = Some(5);

    println!("{:?},{:?}", x, y);

    value_in_cents(Coin::Quater(UsState::Alaska));

    let five = Some(5);
    let six = plus_one(five);
    let none = plus_one(None);

    println!("{:?}, {:?}", six, none);

    let dice_roll = 9;
    match dice_roll {
        3 => add_fancy_hat(),
        7 => remove_fancy_hat(),
        other => move_player(other),
    }

    let config_max = Some(3u8);
    match config_max {
        Some(max) => println!("The maximum is configured to be {}", max),
        _ => (),
    }

    if let Some(max) = config_max {
        println!("The maximum is configured to be {}", max);
    }

    string_practice();
    hashmap_practice();
}

fn hashmap_practice() {
    use std::collections::HashMap;

    let mut scores = HashMap::new();

    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);

    let team_name = String::from("Blue");
    let score = scores.get(&team_name).copied().unwrap_or(0);
    println!("{score}");

    for (key, value) in &scores {
        println!("{key}, {value}");
    }

    scores.insert(String::from("Blue"), 25);
    println!("{:?}", scores);

    scores.entry(String::from("Yellow")).or_insert(50);
    scores.entry(String::from("Blue")).or_insert(50);
    println!("{:?}", scores);

    let text = "hello world wonderful world";
    let mut map = HashMap::new();
    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;
    }
    println!("{:?}", map)
}
fn vector_practice() {
    let mut v: Vec<i32> = Vec::new();
    v.push(5);

    let first: &i32 = &v[0];

    println!("{first}");

    let v1 = vec![1, 2, 3];
    let third: Option<&i32> = v1.get(2);
    match third {
        Some(third) => println!("{third}"),
        None => println!("no"),
    }

    for i in &mut v {
        *i += 50;
        println!("{i}");
    }

    enum SpreadsheetCell {
        Int(i32),
        Float(f64),
        Text(String),
    }

    let row = vec![
        SpreadsheetCell::Int(3),
        SpreadsheetCell::Text(String::from("blue")),
        SpreadsheetCell::Float(10.12),
    ];
}

fn string_practice() {
    let data = "initial contents";

    let mut s = "initial contents".to_string();

    s.push_str("bar");

    s = s + data;

    println!("{data}, {s}");

    let s = format!("{s}-{data}");
    println!("{s}");

    let hello = String::from("Здравствуйте");
    println!("{hello}");
    println!("{}", &hello[0..4]);

    for c in "Зд".chars() {
        println!("{c}");
    }
}

fn read_username_from_file() -> io::Result<String> {
    fs::read_to_string("hello.txt")
}

fn err_result() {
    let greeting_file_result = File::open("hello.text");
    let greeting_file = match greeting_file_result {
        Ok(file) => file,
        Err(error) => panic!("Problem Opening the file: {:?}", error),
    };

    let greeting_file_result1 = File::open("hello.txt").unwrap_or_else(|error| {
        if error.kind() == ErrorKind::NotFound {
            File::create("hello.txt").unwrap_or_else(|error| {
                panic!("Problem creating the file: {:?}", error);
            })
        } else {
            panic!("Problem opening the file: {:?}", error);
        }
    });

    let greeting_file_result2 =
        File::open("hello.txt").expect("hello.text should be included in this project");
}

fn largest<T: std::cmp::PartialOrd>(list: &[T]) -> &T {
    let mut largest = &list[0];

    for item in list {
        if item > largest {
            largest = item;
        }
    }

    largest
}

fn template() {
    let mut number_list = vec![1, 2, 34, 5];
    number_list.push(40);
    let result = largest(&number_list);
    println!("{result}");

    let char_list = vec!['y', 'm', 'z'];
    let result = largest(&char_list);
    println!("{result}");

    let interger = Point { x: 5, y: 10 };
    let float = Point { x: 1.0, y: 4.0 };
}

struct Point<T> {
    x: T,
    y: T,
}

impl<T> Point<T> {
    fn x(&self) -> &T {
        &self.x
    }
}

impl Point<f32> {
    fn distance_from_origin(&self) -> f32 {
        (self.x.powi(2) + self.y.powi(2)).sqrt()
    }
}

enum CusOption<T> {
    Some(T),
    None,
}

pub trait Summary {
    fn summarize(&self) -> String {
        format!("(Read more from {}...)", self.summarize_author())
    }

    fn summarize_author(&self) -> String {
        format!("@unknow")
    }
}

pub struct NewsArticle {
    headline: String,
    localtion: String,
    author: String,
    content: String,
}

impl Summary for NewsArticle {
    fn summarize(&self) -> String {
        format!("(Read more from {}...)", self.summarize_author())
    }
}

pub struct Tweet {
    username: String,
    content: String,
    reply: bool,
    retweet: bool,
}

impl Summary for Tweet {
    fn summarize_author(&self) -> String {
        format!("@{}", self.username)
    }
}

pub fn notify(item: &impl Summary) {
    println!("Breaking news! {}", item.summarize());
}

pub fn notify_t<T: Summary>(item1: &T, item2: &T) {
    println!("Breaking new! {}", item1.summarize());
    println!("Breaking new author {}", item2.summarize_author());
}

pub fn notify_plus(item: &(impl Summary + Display)) {
    println!("{item}");
}

pub fn notify_t_plus<T: Summary + Display>(item: &T) {
    println!("{item}");
}

fn trait_practice() {
    let tweet = Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    };
    notify(&tweet);

    let article = NewsArticle {
        localtion: String::from("location"),
        author: String::from("article_author"),
        content: String::from("nothing"),
        headline: String::from("News"),
    };
    notify(&article);

    notify_t(&tweet, &tweet);
}

fn some_function<T: Display + Clone, U: Clone + Summary>(t: &T, u: &U) {
    println!("{t} {}", u.summarize());
}

fn some_function_where<T, U>(t: &T, u: &U)
where
    T: Display + Clone,
    U: Clone + Summary,
{
    println!("{t} {}", u.summarize());
}

fn main() {
    trait_practice();

    let i = 32;

    if i == 32 {
        println!("1")
    }

    println!("2")
}

fn add_fancy_hat() {}

fn remove_fancy_hat() {}

fn move_player(num_spaces: u8) {}

fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        None => None,
        Some(i) => Some(i + 1),
    }
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quater(state) => {
            println!("State quater from {:?}!", state);
            26
        }
    }
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
