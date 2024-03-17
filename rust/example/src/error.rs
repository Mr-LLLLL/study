use std::num::ParseIntError;
use std::{error, fmt};

fn give_princess(gift: &str) {
    if gift == "snake" {
        panic!("AAAaaaaa!!!");
    }

    println!("I love {}s!!!!", gift);
}

fn give_commoner(gift: Option<&str>) {
    match gift {
        Some("snake") => println!("Yuck! I'm throwing that snake in a fire."),
        Some(inner) => println!("{}? How nice.", inner),
        None => println!("No gift? Oh well."),
    }
}

fn give_princess1(gift: Option<&str>) {
    let inside = gift.unwrap();
    if inside == "snake" {
        panic!("AAAaaa!!!");
    }

    println!("I love {}s!!!", inside);
}

fn next_birthday(current_age: Option<u8>) -> Option<String> {
    let next_age: u8 = current_age?;
    Some(format!("Next year I will be {}", next_age))
}

struct Person {
    job: Option<Job>,
}

#[derive(Clone, Copy)]
struct Job {
    phone_number: Option<PhoneNumber>,
}

#[derive(Clone, Copy)]
struct PhoneNumber {
    area_code: Option<u8>,
    number: u32,
}

impl Person {
    fn work_phone_area_code(&self) -> Option<u8> {
        self.job?.phone_number?.area_code
    }
}

#[derive(Debug)]
enum Food {
    Apple,
    Carrot,
    Potato,

    CordonBlue,
    Steak,
    Sushi,
}

#[derive(Debug)]
enum Day {
    Monday,
    Tuesday,
    Wednesday,
}

#[derive(Debug)]
struct Peeled(Food);

#[derive(Debug)]
struct Chopped(Food);

#[derive(Debug)]
struct Cooked(Food);

fn peel(food: Option<Food>) -> Option<Peeled> {
    match food {
        Some(food) => Some(Peeled(food)),
        None => None,
    }
}

fn chop(peeled: Option<Peeled>) -> Option<Chopped> {
    match peeled {
        Some(Peeled(food)) => Some(Chopped(food)),
        None => None,
    }
}

fn cook(chopped: Option<Chopped>) -> Option<Cooked> {
    chopped.map(|Chopped(food)| Cooked(food))
}

fn process(food: Option<Food>) -> Option<Cooked> {
    food.map(|f| Peeled(f))
        .map(|Peeled(f)| Chopped(f))
        .map(|Chopped(f)| Cooked(f))
}

fn eat(food: Option<Cooked>) {
    match food {
        Some(food) => println!("Mnn. I love {:?}", food),
        None => println!("Oh no! It wasn't edible."),
    }
}

fn have_ingredients(food: Food) -> Option<Food> {
    match food {
        Food::CordonBlue => None,
        _ => Some(food),
    }
}

fn have_recipe(food: Food) -> Option<Food> {
    match food {
        Food::CordonBlue => None,
        _ => Some(food),
    }
}

fn cookable_v1(food: Food) -> Option<Food> {
    match have_ingredients(food) {
        None => None,
        Some(food) => match have_recipe(food) {
            None => None,
            Some(food) => Some(food),
        },
    }
}

fn cookable_v2(food: Food) -> Option<Food> {
    have_ingredients(food).and_then(have_recipe)
}

fn eat_v2(food: Food, day: Day) {
    match cookable_v2(food) {
        Some(food) => println!("Yay! On {:?} we get to eat {:?}", day, food),
        None => println!("Oh no. We don't get to eat on {:?}", day),
    }
}

type AliasedResult<T> = Result<T, ParseIntError>;

fn multiply2(first_number_str: &str, second_number_str: &str) -> AliasedResult<i32> {
    let first_number = r#try!(first_number_str.parse::<i32>());
    let second_number = r#try!(second_number_str.parse::<i32>());

    Ok(first_number * second_number)
}

fn multiply1(first_number_str: &str, second_number_str: &str) -> AliasedResult<i32> {
    let first_number = first_number_str.parse::<i32>()?;
    let second_number = second_number_str.parse::<i32>()?;

    Ok(first_number * second_number)
}

fn multiply(first_number_str: &str, second_number_str: &str) -> AliasedResult<i32> {
    first_number_str.parse::<i32>().and_then(|first_number| {
        second_number_str
            .parse::<i32>()
            .map(|second_number| first_number * second_number)
    })
}

fn print(result: AliasedResult<i32>) {
    match result {
        Ok(n) => println!("n is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

#[derive(Debug, Clone)]
struct DoubleError;

type Result1<T> = std::result::Result<T, DoubleError>;

impl fmt::Display for DoubleError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "invalid first item to double")
    }
}

impl error::Error for DoubleError {
    fn source(&self) -> Option<&(dyn error::Error + 'static)> {
        None
    }
}

fn double_first1(vec: &Vec<&str>) -> Result1<i32> {
    vec.first()
        .ok_or(DoubleError)
        .and_then(|s| s.parse::<i32>().map_err(|_| DoubleError).map(|i| 2 * i))
}

fn print1(result: Result1<i32>) {
    match result {
        Ok(n) => println!("The first doubled is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

type Result2<T> = std::result::Result<T, Box<dyn error::Error>>;

#[derive(Debug, Clone)]
struct EmptyVec;

impl fmt::Display for EmptyVec {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "invalid first item to double")
    }
}

impl error::Error for EmptyVec {
    fn description(&self) -> &str {
        "inalid first item to double"
    }

    fn cause(&self) -> Option<&dyn error::Error> {
        None
    }
}

fn double_first2(vec: &Vec<&str>) -> Result2<i32> {
    vec.first()
        .ok_or_else(|| EmptyVec.into())
        .and_then(|s| s.parse::<i32>().map_err(|e| e.into()).map(|i| 2 * i))
}

fn double_first3(vec: &Vec<&str>) -> Result2<i32> {
    let first = vec.first().ok_or(EmptyVec)?;
    let parsed = first.parse::<i32>()?;
    Ok(2 * parsed)
}

fn print2(result: Result2<i32>) {
    match result {
        Ok(n) => println!("The first doubled is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

#[derive(Debug)]
enum DoubleErrorEnum {
    EmptyVec,
    Parse(ParseIntError),
}

type Result3<T> = std::result::Result<T, DoubleErrorEnum>;

impl fmt::Display for DoubleErrorEnum {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match *self {
            DoubleErrorEnum::EmptyVec => {
                write!(f, "please use a vector with at least one element")
            }
            DoubleErrorEnum::Parse(ref e) => e.fmt(f),
        }
    }
}

impl error::Error for DoubleErrorEnum {
    fn source(&self) -> Option<&(dyn error::Error + 'static)> {
        match *self {
            DoubleErrorEnum::EmptyVec => None,
            DoubleErrorEnum::Parse(ref e) => Some(e),
        }
    }
}

impl From<ParseIntError> for DoubleErrorEnum {
    fn from(value: ParseIntError) -> Self {
        DoubleErrorEnum::Parse(value)
    }
}

fn double_first4(vec: &Vec<&str>) -> Result3<i32> {
    let first = vec.first().ok_or(DoubleErrorEnum::EmptyVec)?;
    let parsed = first.parse::<i32>()?;

    Ok(2 * parsed)
}

fn print4(result: Result3<i32>) {
    match result {
        Ok(n) => println!("The first doubled is {}", n),
        Err(e) => println!("Error: {}", e),
    }
}

pub fn practise_error() {
    let food = Some("chicken");
    let snake = Some("snake");
    // let void = None;

    give_commoner(food);
    give_commoner(snake);
    // give_commoner(void);

    let bird = Some("robin");
    // let nothing = None;

    give_princess1(bird);
    // give_princess1(nothing);

    let p = Person {
        job: Some(Job {
            phone_number: Some(PhoneNumber {
                area_code: Some(61),
                number: 439222222,
            }),
        }),
    };

    assert_eq!(p.work_phone_area_code(), Some(61));

    let apple = Some(Food::Apple);
    let carrot = Some(Food::Carrot);
    let potato = Some(Food::Potato);

    let cooked_apple = cook(chop(peel(apple)));
    let cooked_carrot = cook(chop(peel(carrot)));

    let cooked_potato = process(potato);

    eat(cooked_apple);
    eat(cooked_carrot);
    eat(cooked_potato);

    let (cordon_blue, steak, sushi) = (Food::CordonBlue, Food::Steak, Food::Sushi);

    eat_v2(cordon_blue, Day::Monday);
    eat_v2(steak, Day::Tuesday);
    eat_v2(sushi, Day::Wednesday);

    let twenty = multiply("10", "2");
    print(twenty);

    let tt = multiply("t", "2");
    let tt1 = multiply1("t", "2");
    let tt2 = multiply1("t", "2");
    print(tt);
    print(tt1);
    print(tt2);

    let numbers = vec!["42", "93", "18"];
    let empty = vec![];
    let strings = vec!["tofu", "93", "18"];

    print1(double_first1(&numbers));
    print1(double_first1(&empty));
    print1(double_first1(&strings));

    println!("");

    print2(double_first2(&numbers));
    print2(double_first2(&empty));
    print2(double_first2(&strings));

    println!("");

    print2(double_first3(&numbers));
    print2(double_first3(&empty));
    print2(double_first3(&strings));

    println!("");

    print4(double_first4(&numbers));
    print4(double_first4(&empty));
    print4(double_first4(&strings));

    let numbers: Vec<i32> = strings
        .iter()
        .filter_map(|s| s.parse::<i32>().ok())
        .collect();
    println!("Results: {:?}", numbers);

    let numbers: Result<Vec<_>, _> = strings.iter().map(|s| s.parse::<i32>()).collect();
    println!("Result: {:?}", numbers);

    let (numbers, errors): (Vec<_>, Vec<_>) = strings
        .iter()
        .map(|s| s.parse::<i32>())
        .partition(Result::is_ok);

    let numbers: Vec<_> = numbers.iter().map(|e| e.clone().unwrap()).collect();
    let errors: Vec<_> = errors.into_iter().map(Result::unwrap_err).collect();

    println!("Numbers: {:?}", numbers);
    println!("Errors: {:?}", errors);
}
