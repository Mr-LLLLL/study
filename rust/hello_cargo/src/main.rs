#![allow(dead_code)]

mod back_of_house;
mod front_of_house;

use std::fmt::Display;
use std::fs::{self, File};
use std::io;
use std::io::ErrorKind;

use std::slice;
use std::thread;
use std::time::Duration;

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

fn hex_or_die_trying(maybe_string: Option<String>) -> core::result::Result<u32, String> {
    let Some(s) = maybe_string else {
        return Err(String::from("got none"));
    };

    let Some(first_byte_char) = s.chars().next() else {
        return Err(String::from("got empty string"));
    };

    let Some(digit) = first_byte_char.to_digit(16) else {
        return Err(String::from("not a hex digit"));
    };

    Ok(digit)
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

impl<T> CusOption<T> {
    pub fn unwrap(self) -> T {
        match self {
            CusOption::Some(val) => val,
            CusOption::None => panic!("called `Option::unwrap()` on a `None` value"),
        }
    }
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

struct Pair<T> {
    x: T,
    y: T,
}

impl<T> Pair<T> {
    fn new(x: T, y: T) -> Self {
        Self { x, y }
    }
}

impl<T: Display + PartialOrd> Pair<T> {
    fn cmp_display(&self) {
        if self.x >= self.y {
            println!("The largest number is x = {}", self.x);
        } else {
            println!("The largest number is y = {}", self.y);
        }
    }
}

struct Counter {
    count: u32,
}

impl Counter {
    fn new() -> Counter {
        Counter { count: 0 }
    }
}

impl Iterator for Counter {
    type Item = u32;

    fn next(&mut self) -> Option<Self::Item> {
        if self.count < 5 {
            self.count += 1;
            Some(self.count)
        } else {
            None
        }
    }
}

trait Pilot {
    fn fly(&self);
}

trait Wizard {
    fn fly(&self);
}

struct Human;

impl Pilot for Human {
    fn fly(&self) {
        println!("This is your captain speaking.");
    }
}

impl Wizard for Human {
    fn fly(&self) {
        println!("Up!");
    }
}

impl Human {
    fn fly(&self) {
        println!("*waving args furiously*");
    }
}

trait Animal {
    fn baby_name() -> String;
}

struct Dog;

impl Dog {
    pub fn baby_name() -> String {
        String::from("Spot")
    }
}

impl Animal for Dog {
    fn baby_name() -> String {
        String::from("puppy")
    }
}

use std::fmt;
trait OutlinePrint: fmt::Display {
    fn outline_print(&self) {
        let output = self.to_string();
        let len = output.len();
        println!("{}", "*".repeat(len + 4));
        println!("*{}*", " ".repeat(len + 2));
        println!("* {} *", output);
        println!("*{}*", " ".repeat(len + 2));
        println!("{}", "*".repeat(len + 4));
    }
}

struct Point_int {
    x: i32,
    y: i32,
}

impl OutlinePrint for Point_int {}

impl fmt::Display for Point_int {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}

struct Wrapper(Vec<String>);

impl fmt::Display for Wrapper {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "[{}]", self.0.join(", "))
    }
}

fn bar() -> ! {
    panic!();
}

fn returns_closure() -> Box<dyn Fn(i32) -> i32> {
    Box::new(|x| x + 1)
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

    println!("");

    let person = Human;
    person.fly();
    Human::fly(&person);
    Pilot::fly(&person);
    Wizard::fly(&person);
    <Human as Pilot>::fly(&person);

    println!("");
    println!("A baby dog is called a {}", <Dog as Animal>::baby_name());

    println!("");
    let p = Point_int { x: 1, y: 1 };
    p.outline_print();

    println!("");
    let w = Wrapper(vec![String::from("hello"), String::from("world")]);
    println!("w = {}", w);

    type Kilometers = i32;
    let x: i32 = 5;
    let y: Kilometers = 5;
    println!("x + y = {}", x + y);
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

fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

struct ImportantExcerpt<'a> {
    part: &'a str,
}

impl<'a> ImportantExcerpt<'a> {
    fn announce_and_return_part(&self, annoucement: &str) -> &str {
        println!("Attention please: {}", annoucement);
        self.part
    }
}

fn longest_with_an_announcement<'a, T>(x: &'a str, y: &'a str, ann: T) -> &'a str
where
    T: Display,
{
    println!("Announcement! {}", ann);
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn livecycle_practice() {
    let s1 = String::from("ab");
    let res;
    {
        let s2 = "xyz";

        res = longest(s1.as_str(), s2);
    }
    println!("The longest string is {res}");

    let novel = String::from("Call me Ishmael. Some years ago...");
    let first_sentence = novel.split('.').next().expect("Could not find a '.'");
    let i = ImportantExcerpt {
        part: first_sentence,
    };

    let s2 = "xyz";

    let res1 = longest_with_an_announcement(s1.as_str(), s2, "Today is someone's birthday");
    println!("The longest string is {res1}");
}

#[derive(Debug, PartialEq, Copy, Clone)]
enum ShirtColor {
    Red,
    Blue,
}

struct Inventory {
    shirts: Vec<ShirtColor>,
}

impl Inventory {
    fn giveaway(&self, user_preference: Option<ShirtColor>) -> ShirtColor {
        user_preference.unwrap_or_else(|| self.most_stocked())
    }

    fn most_stocked(&self) -> ShirtColor {
        let mut num_red = 0;
        let mut num_blue = 0;

        for color in &self.shirts {
            match color {
                ShirtColor::Red => num_red += 1,
                ShirtColor::Blue => num_blue += 1,
            }
        }
        if num_red > num_blue {
            ShirtColor::Red
        } else {
            ShirtColor::Blue
        }
    }
}

fn generate_workout(intensity: u32, random_number: u32) {
    let expensive_closure = |num: u32| -> u32 {
        println!("calculating slowly...");
        num
    };

    if intensity < 25 {
        println!("Today, do {} pushups!", expensive_closure(intensity));
        println!("Next, do {} situps!", expensive_closure(intensity));
    } else {
        if random_number == 3 {
            println!("Take a break today! Remember to stay hydrated!");
        } else {
            println!("Today, run for {} minutes!", expensive_closure(intensity));
        }
    }
}

fn closures_practice() {
    let store = Inventory {
        shirts: vec![ShirtColor::Blue, ShirtColor::Red, ShirtColor::Blue],
    };

    let user_pref1 = Some(ShirtColor::Red);
    let giveaway1 = store.giveaway(user_pref1);
    println!(
        "The user with preference {:?} gets {:?}",
        user_pref1, giveaway1
    );

    let user_pref2 = None;
    let giveaway2 = store.giveaway(user_pref2);
    println!(
        "The user with preference {:?} gets {:?}",
        user_pref2, giveaway2
    );

    let simulated_user_specified_value = 10;
    let simulated_random_number = 7;
    generate_workout(simulated_user_specified_value, simulated_random_number);

    let list = vec![1, 2, 3];
    println!("Before defining closure: {:?}", list);

    let only_borrows = || println!("From closure: {:?}", list);

    println!("Before calling closure: {:?}", list);
    only_borrows();
    println!("After calling closure: {:?}", list);

    let mut mut_list = vec![1, 2, 3, 4];
    let mut borrows_mutably = || mut_list.push(7);

    borrows_mutably();
    println!("After calling closure: {:?}", mut_list);

    let list2 = vec![1, 2, 3];
    println!("Before defining closure: {:?}", list);

    thread::spawn(move || println!("From thread: {:?}", list2))
        .join()
        .unwrap();

    let mut list = [
        Rectangle {
            width: 10,
            height: 1,
        },
        Rectangle {
            width: 3,
            height: 1,
        },
        Rectangle {
            width: 5,
            height: 3,
        },
    ];

    let mut num_sort_operations = 0;

    list.sort_by_key(|r| {
        num_sort_operations += 1;
        r.width
    });
    println!("{:#?}, {num_sort_operations}", list);
}

fn shoes_in_size(shoes: Vec<Shoe>, shoe_size: u32) -> Vec<Shoe> {
    shoes.into_iter().filter(|s| s.size == shoe_size).collect()
}

#[derive(Debug)]
struct Shoe {
    size: u32,
    style: String,
}

fn iterrator_practice() {
    let v1 = vec![1, 2, 3, 4];

    let v1_iter = v1.iter();

    for val in v1_iter {
        println!("Got: {}", val);
    }

    let mut v2_iter = v1.iter();
    println!("Got:{}", v2_iter.next().unwrap());
    println!("Got:{}", v2_iter.next().unwrap());

    let total: i32 = v2_iter.sum();
    println!("Got: {}", total);

    let v3_iter = v1.iter();
    let v2: Vec<_> = v3_iter.map(|x| x + 1).collect();
    println!("Got: {:?}", v2);

    let shoes = vec![
        Shoe {
            size: 10,
            style: String::from("sneaker"),
        },
        Shoe {
            size: 13,
            style: String::from("sandal"),
        },
    ];

    let in_my_size = shoes_in_size(shoes, 10);
    println!("{:?}", in_my_size);
}

enum List {
    Cons(i32, Box<List>),
    Nil,
}

use crate::List::{Cons, Nil};

use std::rc::{Rc, Weak};
enum List1 {
    Cons(i32, Rc<List1>),
    Nil,
}

use crate::List1::{Cons as Cons1, Nil as Nil1};

use std::cell::RefCell;
#[derive(Debug)]
enum List2 {
    Cons(Rc<RefCell<i32>>, Rc<List2>),
    Nil,
}

use crate::List2::{Cons as Cons2, Nil as Nil2};

#[derive(Debug)]
enum List3 {
    Cons(i32, RefCell<Rc<List3>>),
    Nil,
}

use crate::List3::{Cons as Cons3, Nil as Nil3};

impl List3 {
    fn tail(&self) -> Option<&RefCell<Rc<List3>>> {
        match self {
            Cons3(_, item) => Some(item),
            Nil3 => None,
        }
    }
}

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

use std::ops::Deref;
impl<T> Deref for MyBox<T> {
    type Target = T;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

struct CustomSmartPointer {
    data: String,
}

impl Drop for CustomSmartPointer {
    fn drop(&mut self) {
        println!("\nDropping CustomSmartPointer with data `{}`!", self.data);
    }
}

#[derive(Debug)]
struct Node {
    value: i32,
    children: RefCell<Vec<Rc<Node>>>,
    parent: RefCell<Weak<Node>>,
}

fn smart_pointer_practice() {
    let b = Box::new(5);
    println!("b = {b}");

    let list = Cons(1, Box::new(Cons(2, Box::new(Cons(3, Box::new(Nil))))));

    let x = 5;
    let y = MyBox::new(5);

    assert_eq!(5, x);
    assert_eq!(5, *y);

    fn hello(name: &str) {
        println!("Hello, {name}!");
    }

    let m = MyBox::new(String::from("Rust"));
    hello(&m);

    let c = CustomSmartPointer {
        data: String::from("my stuff"),
    };
    let d = CustomSmartPointer {
        data: String::from("other stuff"),
    };
    println!("CustomSmartPointer created.");
    drop(c);
    println!("CustomSmartPointer dropped before the end of main.");

    let a = Rc::new(Cons1(5, Rc::new(Cons1(10, Rc::new(Nil1)))));
    println!("count after createing a = {}", Rc::strong_count(&a));
    let b = Cons1(3, Rc::clone(&a));
    println!("count after creating b = {}", Rc::strong_count(&a));
    {
        let c = Cons1(5, Rc::clone(&a));
        println!("count after creating c = {}", Rc::strong_count(&a));
    }
    println!("count after c goes out of scope = {}", Rc::strong_count(&a));

    let value = Rc::new(RefCell::new(5));

    let a = Rc::new(Cons2(Rc::clone(&value), Rc::new(Nil2)));
    let b = Cons2(Rc::new(RefCell::new(3)), Rc::clone(&a));
    let c = Cons2(Rc::new(RefCell::new(4)), Rc::clone(&a));

    *value.borrow_mut() += 10;
    println!("a after = {:?}", a);
    println!("b after = {:?}", b);
    println!("c after = {:?}", c);

    println!("");

    let a = Rc::new(Cons3(5, RefCell::new(Rc::new(Nil3))));

    println!("a initial rc count = {}", Rc::strong_count(&a));
    println!("a next item = {:?}", a.tail());

    let b = Rc::new(Cons3(10, RefCell::new(Rc::clone(&a))));

    println!("a rc count after b creation = {}", Rc::strong_count(&a));
    println!("b initial rc count = {}", Rc::strong_count(&b));
    println!("b next item = {:?}", b.tail());

    if let Some(link) = a.tail() {
        *link.borrow_mut() = Rc::clone(&b);
    }

    println!("b rc count after changing a = {}", Rc::strong_count(&b));
    println!("a rc count after changing a = {}", Rc::strong_count(&a));

    println!("");
    let leaf = Rc::new(Node {
        value: 3,
        parent: RefCell::new(Weak::new()),
        children: RefCell::new(vec![]),
    });

    println!(
        "leaf strong = {}, weak = {}",
        Rc::strong_count(&leaf),
        Rc::weak_count(&leaf)
    );

    {
        let branch = Rc::new(Node {
            value: 5,
            parent: RefCell::new(Weak::new()),
            children: RefCell::new(vec![Rc::clone(&leaf)]),
        });

        *leaf.parent.borrow_mut() = Rc::downgrade(&branch);

        println!(
            "branch strong = {}, weak = {}",
            Rc::strong_count(&branch),
            Rc::weak_count(&branch)
        );
        println!(
            "leaf strong = {}, weak = {}",
            Rc::strong_count(&leaf),
            Rc::weak_count(&leaf)
        )
    }

    println!("leaf parent = {:?}", leaf.parent.borrow().upgrade());
    println!(
        "leaf strong = {}, weak = {}",
        Rc::strong_count(&leaf),
        Rc::weak_count(&leaf)
    );
}

use std::sync::mpsc;
use std::sync::{Arc, Mutex, MutexGuard};
fn parallel_practice() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn(move || {
        println!("hi number {:?} from the spawned thread!", v);
    });

    for i in 1..5 {
        println!("hi number {} from the main thread!", i);
        thread::sleep(Duration::from_millis(1));
    }

    handle.join().unwrap();

    let (tx, rx) = mpsc::channel();

    let tx1 = tx.clone();
    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];
        for val in vals {
            tx.send(val).unwrap();
        }
    });
    thread::spawn(move || {
        let vals = vec![
            String::from("more"),
            String::from("messages"),
            String::from("for"),
            String::from("you"),
        ];
        for val in vals {
            tx1.send(val).unwrap();
        }
    });
    for received in rx {
        println!("Got: {}", received);
    }

    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();

            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }
    println!("Result: {}", *counter.lock().unwrap());
}

pub struct AveragedCollection {
    list: Vec<i32>,
    average: f64,
}

impl AveragedCollection {
    pub fn add(&mut self, value: i32) {
        self.list.push(value);
        self.update_average();
    }

    pub fn remove(&mut self) -> Option<i32> {
        let result = self.list.pop();
        match result {
            Some(value) => {
                self.update_average();
                Some(value)
            }
            None => None,
        }
    }

    pub fn average(&self) -> f64 {
        self.average
    }

    fn update_average(&mut self) {
        let total: i32 = self.list.iter().sum();
        self.average = total as f64 / self.list.len() as f64;
    }
}

fn split_at_mut(values: &mut [i32], mid: usize) -> (&mut [i32], &mut [i32]) {
    let len = values.len();
    let ptr = values.as_mut_ptr();

    assert!(mid <= len);

    unsafe {
        (
            slice::from_raw_parts_mut(ptr, mid),
            slice::from_raw_parts_mut(ptr.add(mid), len - mid),
        )
    }
}

extern "C" {
    fn abs(input: i32) -> i32;
}

static HELLO_WORLD: &str = "Hello, world!";

static mut COUTNER: u32 = 0;

fn add_to_count(inc: u32) {
    unsafe {
        COUTNER += inc;
    }
}

fn parallel_unsafe() {
    let mut vector = vec![1, 2, 3, 4, 5, 6];
    let (left, right) = split_at_mut(&mut vector, 3);

    unsafe {
        println!("Absolute value of -3 according to C: {}", abs(-3));
    }

    println!("name is {}", HELLO_WORLD);

    add_to_count(3);
    unsafe {
        println!("COUNTER: {}", COUTNER);
    }
}

#[macro_export]
macro_rules! vec_cus {
    ($ ($x:expr), *) => {
        {
        let mut temp_vec = Vec::new();
        $(
            temp_vec.push($x);
        )*
        temp_vec
        }
    };
}

pub trait HelloMacro {
    fn hello_macro();
}
use hello_macro_derive::HelloMacro;

#[derive(HelloMacro)]
struct Pancakes;

fn practice_macro() {
    let v = vec_cus![1, 2, 3];
    println!("{:?}", v);

    Pancakes::hello_macro();
}

fn main() {
    practice_macro()
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
