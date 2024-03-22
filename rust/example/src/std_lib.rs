#![allow(dead_code)]

use std::{
    collections::{HashMap, HashSet},
    mem,
};

use std::rc::Rc;

use self::checked::sqrt;

#[derive(Debug, Clone, Copy)]
struct Point {
    x: f64,
    y: f64,
}

struct Rectangle {
    p1: Point,
    p2: Point,
}

fn origin() -> Point {
    Point { x: 0.0, y: 0.0 }
}

fn boxed_origin() -> Box<Point> {
    Box::new(Point { x: 0.0, y: 0.0 })
}

fn practise_box() {
    let point: Point = origin();
    let rectangle: Rectangle = Rectangle {
        p1: origin(),
        p2: Point { x: 3.0, y: 4.0 },
    };

    let boxed_rectangle: Box<Rectangle> = Box::new(Rectangle {
        p1: origin(),
        p2: origin(),
    });

    let boxed_point: Box<Point> = Box::new(origin());

    let box_in_a_box: Box<Box<Point>> = Box::new(boxed_origin());

    println!(
        "Point occupies {} bytes on the stack",
        mem::size_of_val(&point)
    );
    println!(
        "Rectangle occupies {} bytes on the stack",
        mem::size_of_val(&rectangle)
    );

    println!(
        "Boxed point occupies {} bytes in the stack",
        mem::size_of_val(&boxed_point)
    );
    println!(
        "Boxed rectangle occupies {} bytes in the stack",
        mem::size_of_val(&boxed_rectangle)
    );
    println!(
        "Boxed box occupies {} bytes in the stack",
        mem::size_of_val(&box_in_a_box)
    );

    let unboxed_point: Point = *boxed_point;
    println!(
        "Unboxed point occupies {} bytes in the stack",
        mem::size_of_val(&unboxed_point)
    );
}

fn practise_vector() {
    let collected_iterator: Vec<i32> = (0..10).collect();
    println!("Collectd (0..10): {:?}", collected_iterator);

    let mut xs = vec![1i32, 2, 3];
    println!("Initial vector: {:?}", xs);

    println!("Push 4 into the vector");
    xs.push(4);
    println!("Vector after push: {:?}", xs);

    // collected_iterator.push(0);
    println!("Vector size: {}", xs.len());

    println!("Second element: {}", xs[1]);

    println!("Pop last element: {:?}", xs.pop());

    // println!("Fourth element: {}", xs[3]);

    println!("Contents of xs:");
    for x in xs.iter() {
        println!("> {}", x);
    }

    for (i, x) in xs.iter().enumerate() {
        println!("In Position {} we have value {}", i, x);
    }

    for x in xs.iter_mut() {
        *x *= 3;
    }
    println!("Updated vector: {:?}", xs);
}

fn checked_division(dividend: i32, divisor: i32) -> Option<i32> {
    if divisor == 0 {
        None
    } else {
        Some(dividend / divisor)
    }
}

fn try_division(divident: i32, divisor: i32) {
    match checked_division(divident, divisor) {
        None => println!("{} / {} failed!", divident, divisor),
        Some(quotient) => println!("{} / {} = {}", divident, divisor, quotient),
    }
}

fn practise_option() {
    try_division(4, 2);
    try_division(1, 0);

    let none: Option<i32> = None;
    let _equivalent_none = None::<i32>;

    let optional_float = Some(0f32);

    println!(
        "{:?} unwrap to {:?}",
        optional_float,
        optional_float.unwrap()
    );

    println!("{:?} unwrap to {:?}", none, none.unwrap());
}

fn practise_string() {
    let pangram: &'static str = "the quick brown fox jumps over the lazy dog";
    println!("Pangram: {}", pangram);

    println!("Words in revers");
    for word in pangram.split_whitespace().rev() {
        println!("> {}", word);
    }

    let mut chars: Vec<char> = pangram.chars().collect();
    chars.sort();
    chars.dedup();

    let mut string = String::new();
    for c in chars {
        string.push(c);
        string.push_str(", ");
    }

    let chars_to_trim: &[char] = &[' ', ','];
    let trimmed_str: &str = string.trim_matches(chars_to_trim);
    println!("Used string: {}", trimmed_str);

    let alice = String::from("I like dogs");
    let bob: String = alice.replace("dog", "cat");

    println!("Alice says: {}", alice);
    println!("Bob says: {}", bob);

    println!("");

    let byte_escape = "I'm writing \x52\x75\x73\x74";
    println!("What are you doing\x3F (\\x3F means ?) {}", byte_escape);

    let unicode_codepoint = "\u{211D}";
    let character_name = "\"Double-Struck Capital R\"";
    println!(
        "Unicode character {} (U+211D) is called {}",
        unicode_codepoint, character_name
    );

    let long_string = "String literals
                            can span multiple lines,
                            The linebreak and indentation here->\
                            <- can be escaped too!";
    println!("{}", long_string);

    let raw_str = r"Escapes don't work here: \x3F \u{211D}";
    println!("{}", raw_str);

    let quotes = r#"And then I said: "There is no escape!""#;
    println!("{}", quotes);

    let longer_delimiter = r###"A string with "# in it. And ever "##!"###;
    println!("{}", longer_delimiter);

    println!("");

    let bytestring: &[u8; 20] = b"this is a bytestring";
    println!("A bytestring: {:?}", bytestring);

    let escaped = b"\x52\x75\x73\x74 as bytes";
    println!("Some escapted bytes: {:?}", escaped);

    let raw_bytestring = br"\u{211D} is not escaped here";
    println!("{:?}", raw_bytestring);

    if let Ok(my_str) = std::str::from_utf8(bytestring) {
        println!("And the same as text: '{}'", my_str);
    }

    // let quotes = br#"You can also use "fancier" formatting, \
    //                 like with normal raw strings"#;

    // 字节串可以不使用 UTF-8 编码
    let shift_jis = b"\x82\xe6\x82\xa8\x82\xb1\x82"; // SHIFT-JIS 编码的 "ようこそ"

    // 但这样的话它们就无法转换成 &str 了
    match std::str::from_utf8(shift_jis) {
        Ok(my_str) => println!("Conversion successful: '{}'", my_str),
        Err(e) => println!("Conversion failed: {:?}", e),
    };
}

mod checked {
    #[derive(Debug)]
    pub enum MathError {
        DivideByZero,
        NegativeLogarithm,
        NegativeSquareRoot,
    }

    pub type MathResult = Result<f64, MathError>;

    pub fn div(x: f64, y: f64) -> MathResult {
        if y == 0.0 {
            Err(MathError::DivideByZero)
        } else {
            Ok(x / y)
        }
    }

    pub fn sqrt(x: f64) -> MathResult {
        if x < 0.0 {
            Err(MathError::NegativeSquareRoot)
        } else {
            Ok(x.sqrt())
        }
    }

    pub fn ln(x: f64) -> MathResult {
        if x < 0.0 {
            Err(MathError::NegativeLogarithm)
        } else {
            Ok(x.ln())
        }
    }
}

fn op(x: f64, y: f64) -> f64 {
    match checked::div(x, y) {
        Err(why) => panic!("{:?}", why),
        Ok(ratio) => match checked::ln(ratio) {
            Err(why) => panic!("{:?}", why),
            Ok(ln) => match checked::sqrt(ln) {
                Err(why) => panic!("{:?}", why),
                Ok(sqrt) => sqrt,
            },
        },
    }
}

fn op_with_question_expresssion(x: f64, y: f64) -> checked::MathResult {
    let ratio = checked::div(x, y)?;

    let ln = checked::ln(ratio)?;

    checked::sqrt(ln)
}

fn division(divident: i32, divisor: i32) -> i32 {
    if divisor == 0 {
        panic!("Cannot divide by zero");
    } else {
        divident / divisor
    }
}

fn practise_result() {
    // println!("{}", op(1.0, 10.0));

    let _x = Box::new(0i32);
    division(3, 0);

    println!("This point won't be reached!");
}

#[derive(PartialEq, Eq, Hash)]
struct Account<'a> {
    username: &'a str,
    password: &'a str,
}

struct AccountInfo<'a> {
    name: &'a str,
    email: &'a str,
}

type Accounts<'a> = HashMap<Account<'a>, AccountInfo<'a>>;

fn try_logon<'a>(accounts: &Accounts<'a>, username: &'a str, password: &'a str) {
    println!("Username: {}", username);
    println!("Password: {}", password);
    println!("Attempting login...");

    let logon = Account { username, password };

    match accounts.get(&logon) {
        Some(account_info) => {
            println!("Login successful!");
            println!("Name: {}", account_info.name);
            println!("Email: {}", account_info.email);
        }
        _ => println!("Login failed!"),
    }
}

fn practise_hashmap() {
    let mut contacts = HashMap::new();

    contacts.insert("Daniel", "798-1364");
    contacts.insert("Ashley", "645-8689");
    contacts.insert("Katie", "435-8291");
    contacts.insert("Robert", "956-1745");

    match contacts.get(&"Daniel") {
        Some(&number) => println!("Calling Daniel: {}", call(number)),
        _ => println!("Do' know Daniel's number."),
    }

    contacts.insert("Daniel", "164-6743");

    match contacts.get(&"Ashley") {
        Some(&number) => println!("Calling Ashley: {}", call(number)),
        _ => println!("Do' know Ashley's number."),
    }

    contacts.remove(&"Ashley");

    for (contact, &number) in contacts.iter() {
        println!("Calling {}: {}", contact, call(number));
    }

    let mut accounts: Accounts = HashMap::new();

    let account = Account {
        username: "j.everyman",
        password: "password123",
    };

    let account_info = AccountInfo {
        name: "John Everyman",
        email: "j.everyman@email.com",
    };

    accounts.insert(account, account_info);
    try_logon(&accounts, "j.everyman", "password123");
    try_logon(&accounts, "j.everyman", "password");

    let mut a: HashSet<i32> = vec![1, 2, 3].into_iter().collect();
    let mut b: HashSet<i32> = vec![2, 3, 4].into_iter().collect();

    assert!(a.insert(4));
    assert!(a.contains(&4));

    assert!(!a.insert(4));

    b.insert(5);

    println!("A: {:?}", a);
    println!("B: {:?}", b);

    println!("Union: {:?}", a.union(&b).collect::<Vec<&i32>>());
    println!(
        "Intersaction: {:?}",
        a.intersection(&b).collect::<Vec<&i32>>()
    );
    println!(
        "Symmetric Difference: {:?}",
        a.symmetric_difference(&b).collect::<Vec<&i32>>()
    );
}

fn call(number: &str) -> &str {
    match number {
        "798-1364" => {
            "We're sorry, the call cannot be completed as dialed. 
            Please hang up and try again."
        }
        "645-7689" => {
            "Hello, this is Mr. Awesome's Pizza. My name is Fred.
            What can I get for you today?"
        }
        _ => "Hi! Who is this again?",
    }
}

fn practise_rc() {
    let rc_examples = "Rc examples".to_string();
    {
        println!("--- rc_a is created ---");

        let rc_a: Rc<String> = Rc::new(rc_examples);
        println!("Reference Count of rc_a: {}", Rc::strong_count(&rc_a));

        {
            println!("--- rc_a is cloned to rc_b ---");

            let rc_b: Rc<String> = Rc::clone(&rc_a);
            println!("Reference Count of rc_a: {}", Rc::strong_count(&rc_a));
            println!("Reference Count of rc_b: {}", Rc::strong_count(&rc_b));

            println!("rc_a and rc_b are equal: {}", rc_a.eq(&rc_b));

            println!("Length of the value inside rc_a: {}", rc_a.len());
            println!("Value of rc_b: {}", rc_b);

            println!("--- rc_b is dropped out of scope ---");
        }

        println!("Reference Count of rc_a: {}", Rc::strong_count(&rc_a));

        println!("--- rc_a is dropped out of scope ---");
    }
}

pub fn practise_std() {
    practise_rc();
}
