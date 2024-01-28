// 类似地，`mod inaccessible` 和 `mod nested` 将找到 `nested.rs` 和
// `inaccessible.rs` 文件，并在它们放到各自的模块中。
mod inaccessible;
pub mod nested;

pub fn function() {
    println!("called `my::function()`");
}

mod cool {
    pub fn function() {
        println!("called `my::cool::function()`");
    }
}

fn private_function() {
    println!("called `my::private_function()`");
}

pub fn indirect_access() {
    print!("called `my::indirect_access()`, that\n> ");

    private_function();
}

pub fn indirect_call() {
    println!("called `my::indirect_call()`");

    self::function();
    function();

    self::cool::function();

    {
        use crate::cool::function as root_function;
        root_function();
    }
}

pub struct OpenBox<T> {
    pub contents: T,
}

#[allow(dead_code)]
pub struct CloseBox<T> {
    contents: T,
}

impl<T> CloseBox<T> {
    pub fn new(contents: T) -> CloseBox<T> {
        CloseBox { contents }
    }
}
