use bevy::{
    input::{
        mouse::{MouseButtonInput, MouseMotion, MouseWheel},
        touchpad::{TouchpadMagnify, TouchpadRotate},
    },
    prelude::*,
};

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(
            Update,
            (bevy::window::close_on_esc, print_mouse_events_system),
        )
        .run();
}

fn print_mouse_events_system(
    mut mouse_button_input_events: EventReader<MouseButtonInput>,
    mut mouse_motion_events: EventReader<MouseMotion>,
    mut cursor_moved_events: EventReader<CursorMoved>,
    mut mouse_wheel_events: EventReader<MouseWheel>,
    mut touchpad_magnify_events: EventReader<TouchpadMagnify>,
    mut touchpad_rorate_events: EventReader<TouchpadRotate>,
) {
    for event in mouse_button_input_events.read() {
        info!("{:?}", event);
    }

    for event in mouse_motion_events.read() {
        info!("{:?}", event);
    }

    for event in cursor_moved_events.read() {
        info!("{:?}", event);
    }

    for event in mouse_wheel_events.read() {
        info!("{:?}", event);
    }

    for event in touchpad_rorate_events.read() {
        info!("{:?}", event);
    }

    for event in touchpad_magnify_events.read() {
        info!("{:?}", event);
    }
}
