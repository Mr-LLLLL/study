use std::ops::Add;
use std::ops::Mul;
use std::ops::Sub;

macro_rules! say_hello {
    () => {
        println!("Hello");
    };
}

macro_rules! create_function {
    // 此宏接受一个 `ident` 指示符表示的参数，并创建一个名为 `$func_name` 的函数。
    // `ident` 指示符用于变量名或函数名
    ($func_name:ident) => {
        fn $func_name() {
            // `stringify!` 宏把 `ident` 转换成字符串。
            println!("You called {:?}()", stringify!($func_name))
        }
    };
}

create_function!(foo1);
create_function!(bar);

// block
// expr 用于表达式
// ident 用于变量名或函数名
// item
// literal 用于字面常量
// pat (模式 pattern)
// path
// stmt (语句 statement)
// tt (标记树 token tree)指示符表示运算符和标记。
// ty (类型 type)
// vis (可见性描述符)
macro_rules! print_result {
    // 此宏接受一个 `expr` 类型的表达式，并将它作为字符串，连同其结果一起
    // 打印出来。
    // `expr` 指示符表示表达式。
    ($expression:expr) => {
        // `stringify!` 把表达式*原样*转换成一个字符串。
        println!("{:?} = {:?}", stringify!($expression), $expression)
    };
}

macro_rules! test {
    // 参数不需要使用逗号隔开。
    // 参数可以任意组合！
    ($left:expr; and $right:expr) => {
        println!(
            "{:?} and {:?} is {:?}",
            stringify!($left),
            stringify!($right),
            $left && $right
        )
    };
    // ^ 每个分支都必须以分号结束。
    ($left:expr; or $right:expr) => {
        println!(
            "{:?} or {:?} is {:?}",
            stringify!($left),
            stringify!($right),
            $left || $right
        )
    };
}

// 宏在参数列表中可以使用 + 来表示一个参数可能出现一次或多次，使用 * 来表示该参数可能出现零次或多次。
macro_rules! find_min {
    ($x:expr) => ($x);
    ($x:expr, $($y:expr),+) => {
        std::cmp::min($x, find_min!($($y),+))
    }
}

macro_rules! assert_equal_len {
    // `tt`（token tree，标记树）指示符表示运算符和标记。
    ($a:ident, $b:ident, $func:ident, $op:tt) => {
        assert!(
            $a.len() == $b.len(),
            "{:?}: dimension mismatch: {:?} {:?} {:?}",
            stringify!($func),
            ($a.len(),),
            stringify!($op),
            ($b.len(),)
        )
    };
}

macro_rules! op {
    ($func:ident, $bound:ident, $op:tt, $method:ident) => {
        fn $func<T: $bound<T, Output = T> + Copy>(xs: &mut Vec<T>, ys: &Vec<T>) {
            assert_equal_len!(xs, ys, $func, $op);

            for (x, y) in xs.iter_mut().zip(ys.iter()) {
                *x = $bound::$method(*x, *y);
            }
        }
    };
}

op!(add_assign, Add, +=, add);
op!(mul_assign, Mul, *=, mul);
op!(sub_assign, Sub, -=, sub);

#[cfg(test)]
mod test {
    use std::iter;

    macro_rules! test {
        ($func:ident, $x:expr, $y:expr, $z:expr) => {
            #[test]
            fn $func() {
                for size in 0usize..10 {
                    let mut x: Vec<_> = iter::repeat($x).take(size).collect();
                    let y: Vec<_> = iter::repeat($y).take(size).collect();
                    let z: Vec<_> = iter::repeat($z).take(size).collect();

                    super::$func(&mut x, &y);

                    assert_eq!(x, z);
                }
            }
        };
    }

    test!(add_assign, 1u32, 2u32, 3u32);
    test!(mul_assign, 2u32, 3u32, 6u32);
    test!(sub_assign, 3u32, 2u32, 1u32);
}

macro_rules! calculate {
    (eval $e:expr) => {
        let val: usize = $e;
        println!("{} = {}", stringify!{$e}, val);
    };

    (eval $e:expr, $(eval $es:expr),+) => {
        calculate!{ eval $e }
        calculate!{ $(eval $es),+ }
    }
}

pub fn practise_macro_rules() {
    say_hello!();
    foo1();
    bar();

    print_result!(1u32 + 1);

    print_result!({
        let x = 1u32;

        x * x + 2 * x - 1
    });

    test!(1i32+1==2i32; and 2i32*2==4i32);
    test!(true;or false);

    println!("{}", find_min!(1u32));
    println!("{}", find_min!(1u32 + 2, 2u32));
    println!("{}", find_min!(5u32, 2u32 * 3, 4u32));

    calculate! {
        eval 1+2
    }

    calculate! {
        eval (1+2)*(3/4)
    }

    calculate! {
        eval 1+2,
        eval 3+4,
        eval (2*3)+1
    }
}
