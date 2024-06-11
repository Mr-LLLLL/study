use bevy::{input::touch::*, prelude::*};

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Update, (touch_system, bevy::window::close_on_esc))
        .run();
}

fn touch_system(touches: Res<Touches>) {
    for touch in touches.iter_just_pressed() {
        info!(
            "just pressed touch with id: {:?}, at: {:?}",
            touch.id(),
            touch.position()
        );
    }

    for touch in touches.iter_just_released() {
        info!(
            "just released touch with id: {:?}, at: {:?}",
            touch.id(),
            touch.position()
        );
    }

    for touch in touches.iter_just_canceled() {
        info!("canceled touch with id: {:?}", touch.id());
    }

    for touch in touches.iter() {
        info!("active touch: {:?}", touch);
        info!(" just_pressed: {}", touches.just_pressed(touch.id()));
    }
}
