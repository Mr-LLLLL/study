use std::f32::consts::PI;
use std::time::Duration;

use bevy::{animation::RepeatAnimation, pbr::CascadeShadowConfigBuilder, prelude::*};

pub fn run() {
    App::new()
        .insert_resource(AmbientLight {
            color: Color::WHITE,
            brightness: 2000.,
        })
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup)
        .add_systems(
            Update,
            (
                setup_scene_once_loaded,
                keyboard_animation_control,
                bevy::window::close_on_esc,
            ),
        )
        .run();
}

#[derive(Resource)]
struct Animations(Vec<Handle<AnimationClip>>);

fn setup(
    mut commands: Commands,
    asset_server: Res<AssetServer>,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<StandardMaterial>>,
) {
    commands.insert_resource(Animations(vec![
        asset_server.load("models/animated/Fox.glb#Animation2"),
        asset_server.load("models/animated/Fox.glb#Animation1"),
        asset_server.load("models/animated/Fox.glb#Animation0"),
    ]));

    commands.spawn(Camera3dBundle {
        transform: Transform::from_xyz(100., 100., 150.)
            .looking_at(Vec3::new(0., 20., 0.), Vec3::Y),
        ..default()
    });

    commands.spawn(PbrBundle {
        mesh: meshes.add(Plane3d::default().mesh().size(500000., 500000.)),
        material: materials.add(Color::rgb(0.3, 0.5, 0.3)),
        ..default()
    });

    commands.spawn(DirectionalLightBundle {
        transform: Transform::from_rotation(Quat::from_euler(EulerRot::ZXY, 0., 1., -PI / 4.)),
        directional_light: DirectionalLight {
            shadows_enabled: true,
            ..default()
        },
        cascade_shadow_config: CascadeShadowConfigBuilder {
            first_cascade_far_bound: 200.,
            maximum_distance: 400.,
            ..default()
        }
        .into(),
        ..default()
    });

    commands.spawn(SceneBundle {
        scene: asset_server.load("models/animated/Fox.glb#Scene0"),
        ..default()
    });

    println!("Animation controls:");
    println!("  - spacebar: play / pause");
    println!("  - arrow up / down: speed up / slow down animation playback");
    println!("  - arrow left / right: seek backward / forward");
    println!("  - digit 1 / 3 / 5: play the animation <digit> times");
    println!("  - L: loop the animation forever");
    println!("  - return: change animation");
}

fn setup_scene_once_loaded(
    animations: Res<Animations>,
    mut players: Query<&mut AnimationPlayer, Added<AnimationPlayer>>,
) {
    for mut player in &mut players {
        player.play(animations.0[0].clone_weak()).repeat();
    }
}

fn keyboard_animation_control(
    keyboard_input: Res<ButtonInput<KeyCode>>,
    mut animation_players: Query<&mut AnimationPlayer>,
    animations: Res<Animations>,
    mut current_animation: Local<usize>,
) {
    for mut player in &mut animation_players {
        if keyboard_input.just_pressed(KeyCode::Space) {
            if player.is_paused() {
                player.resume();
            } else {
                player.pause();
            }
        }

        if keyboard_input.just_pressed(KeyCode::ArrowUp) {
            let speed = player.speed();
            player.set_speed(speed * 1.2);
        }

        if keyboard_input.just_pressed(KeyCode::ArrowDown) {
            let speed = player.speed();
            player.set_speed(speed * 0.8);
        }

        if keyboard_input.just_pressed(KeyCode::ArrowLeft) {
            let elapsed = player.seek_time();
            player.seek_to(elapsed - 0.1);
        }

        if keyboard_input.just_pressed(KeyCode::ArrowRight) {
            let elapsed = player.seek_time();
            player.seek_to(elapsed + 0.1);
        }

        if keyboard_input.just_pressed(KeyCode::Enter) {
            *current_animation = (*current_animation + 1) % animations.0.len();
            player
                .play_with_transition(
                    animations.0[*current_animation].clone_weak(),
                    Duration::from_millis(250),
                )
                .repeat();
        }

        if keyboard_input.just_pressed(KeyCode::Digit1) {
            player.set_repeat(RepeatAnimation::Count(1));
            player.replay();
        }

        if keyboard_input.just_pressed(KeyCode::Digit3) {
            player.set_repeat(RepeatAnimation::Count(3));
            player.replay();
        }

        if keyboard_input.just_pressed(KeyCode::Digit5) {
            player.set_repeat(RepeatAnimation::Count(5));
            player.replay();
        }

        if keyboard_input.just_pressed(KeyCode::KeyL) {
            player.set_repeat(RepeatAnimation::Forever);
        }
    }
}
