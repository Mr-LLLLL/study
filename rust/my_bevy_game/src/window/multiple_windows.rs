use bevy::{prelude::*, render::camera::RenderTarget, window::WindowRef};

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup_scene)
        .add_systems(Update, bevy::window::close_on_esc)
        .run();
}

fn setup_scene(mut commands: Commands, asset_server: Res<AssetServer>) {
    commands.spawn(SceneBundle {
        scene: asset_server.load("models/torus/torus.gltf#Scene0"),
        ..default()
    });

    commands.spawn(DirectionalLightBundle {
        transform: Transform::from_xyz(4., 5., 4.).looking_at(Vec3::ZERO, Vec3::Y),
        ..default()
    });

    let first_window_camera = commands
        .spawn(Camera3dBundle {
            transform: Transform::from_xyz(0., 0., 6.).looking_at(Vec3::ZERO, Vec3::Y),
            ..default()
        })
        .id();

    let second_window = commands
        .spawn(Window {
            title: "Second window".to_owned(),
            ..default()
        })
        .id();

    let second_window_camera = commands
        .spawn(Camera3dBundle {
            transform: Transform::from_xyz(6., 0., 0.).looking_at(Vec3::ZERO, Vec3::Y),
            camera: Camera {
                target: RenderTarget::Window(WindowRef::Entity(second_window)),
                ..default()
            },
            ..default()
        })
        .id();

    commands
        .spawn((NodeBundle::default(), TargetCamera(first_window_camera)))
        .with_children(|parent| {
            parent.spawn(TextBundle::from_section(
                "First window",
                TextStyle::default(),
            ));
        });

    commands
        .spawn((NodeBundle::default(), TargetCamera(second_window_camera)))
        .with_children(|parent| {
            parent.spawn(TextBundle::from_section(
                "Second window",
                TextStyle::default(),
            ));
        });
}
