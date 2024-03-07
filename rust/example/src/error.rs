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
}
