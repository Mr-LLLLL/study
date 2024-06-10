use bevy::{
    prelude::*,
    window::{close_on_esc, ReceivedCharacter},
};

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Update, (print_char_event_system, close_on_esc))
        .run();
}

fn print_char_event_system(mut char_input_events: EventReader<ReceivedCharacter>) {
    for event in char_input_events.read() {
        info!("{:?}: '{}'", event, event.char);
    }
}
