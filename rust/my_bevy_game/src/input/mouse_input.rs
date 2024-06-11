use bevy::prelude::*;

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Update, (mouse_click_system, bevy::window::close_on_esc))
        .run();
}

fn mouse_click_system(mouse_button_input: Res<ButtonInput<MouseButton>>) {
    if mouse_button_input.pressed(MouseButton::Left) {
        info!("left mouse currently pressed");
    }

    if mouse_button_input.just_pressed(MouseButton::Left) {
        info!("left mouse just pressed");
    }

    if mouse_button_input.just_released(MouseButton::Left) {
        info!("left mouse just released");
    }
}
