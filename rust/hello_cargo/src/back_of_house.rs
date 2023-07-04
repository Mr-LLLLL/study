fn fix_incorrect_order() {
    cook_order();
    super::deliver_order();
}

fn cook_order() {}

pub enum Appetizer {
    Soup,
    Salad,
}
