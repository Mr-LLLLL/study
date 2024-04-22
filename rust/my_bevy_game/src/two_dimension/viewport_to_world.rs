use bevy::prelude::*;

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup)
        .add_systems(Update, draw_cursor)
        .run();
}

fn draw_cursor(
    camera_query: Query<(&Camera, &GlobalTransform)>,
    windows: Query<&Window>,
    mut gizmos: Gizmos,
) {
    let (camera, camera_transform) = camera_query.single();

    let Some(cursor_position) = windows.single().cursor_position() else {
        return;
    };

    let Some(point) = camera.viewport_to_world_2d(camera_transform, cursor_position) else {
        return;
    };

    gizmos.circle_2d(point, 10., Color::WHITE);
}

fn setup(mut commands: Commands) {
    commands.spawn(Camera2dBundle::default());
}
