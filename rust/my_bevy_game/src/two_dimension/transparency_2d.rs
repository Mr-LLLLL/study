use bevy::prelude::*;

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup)
        .run();
}

fn setup(mut commands: Commands, asset_server: Res<AssetServer>) {
    commands.spawn(Camera2dBundle::default());

    let sprite_handle = asset_server.load("branding/icon.png");

    commands.spawn(SpriteBundle {
        texture: sprite_handle.clone(),
        ..default()
    });

    commands.spawn(SpriteBundle {
        sprite: Sprite {
            color: Color::rgba(0., 0., 1., 0.7),
            ..default()
        },
        texture: sprite_handle.clone(),
        transform: Transform::from_xyz(100., 0., 0.),
        ..default()
    });
    commands.spawn(SpriteBundle {
        sprite: Sprite {
            color: Color::rgba(0., 1., 0., 0.3),
            ..default()
        },
        texture: sprite_handle,
        transform: Transform::from_xyz(200., 0., 0.),
        ..default()
    });
}
