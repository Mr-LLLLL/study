use bevy::{input::keyboard::KeyboardInput, prelude::*};

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(
            Update,
            (print_keyboard_event_system, bevy::window::close_on_esc),
        )
        .run();
}

fn print_keyboard_event_system(mut keybaord_input_events: EventReader<KeyboardInput>) {
    for event in keybaord_input_events.read() {
        info!("{:?}", event);
    }
}
