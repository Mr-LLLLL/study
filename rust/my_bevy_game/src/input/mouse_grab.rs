use bevy::{
    prelude::*,
    window::{close_on_esc, CursorGrabMode},
};

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Update, (grab_mouse, close_on_esc))
        .run();
}

fn grab_mouse(
    mut windows: Query<&mut Window>,
    mouse: Res<ButtonInput<MouseButton>>,
    key: Res<ButtonInput<KeyCode>>,
) {
    let mut window = windows.single_mut();

    if mouse.just_pressed(MouseButton::Left) {
        window.cursor.visible = false;
        window.cursor.grab_mode = CursorGrabMode::Locked;
    }

    if key.just_pressed(KeyCode::KeyQ) {
        window.cursor.visible = true;
        window.cursor.grab_mode = CursorGrabMode::None;
    }
}
