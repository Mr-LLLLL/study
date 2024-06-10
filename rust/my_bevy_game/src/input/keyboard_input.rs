use bevy::prelude::*;

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Update, (keyboard_input_system, bevy::window::close_on_esc))
        .run();
}

fn keyboard_input_system(keyboard_input: Res<ButtonInput<KeyCode>>) {
    if keyboard_input.pressed(KeyCode::KeyA) {
        info!("'A' currently pressd");
    }

    if keyboard_input.just_pressed(KeyCode::KeyA) {
        info!("'A' just pressed");
    }

    if keyboard_input.just_released(KeyCode::KeyA) {
        info!("'A' just released");
    }
}
