//! This example demonstrates Bevy's immediate mode drawing API intended for visual debugging.

use std::f32::consts::{PI, TAU};

use bevy::prelude::*;

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .init_gizmo_group::<MyRoundGizmos>()
        .add_systems(Startup, setup)
        .add_systems(
            Update,
            (
                draw_example_collection,
                update_config,
                bevy::window::close_on_esc,
            ),
        )
        .run();
}

#[derive(Default, Reflect, GizmoConfigGroup)]
struct MyRoundGizmos;

fn setup(mut commands: Commands, asset_server: Res<AssetServer>) {
    commands.spawn(Camera2dBundle::default());

    commands.spawn(TextBundle::from_section(
        "Hold 'Left' or 'Right' to change the line width of straight gizmos\n\
        Hold 'Up' or 'Down' to change the line width of round gizmos\n\
        Press '1' or '2' to toggle the visibility of straight gizmos or round gizmos",
        TextStyle {
            font: asset_server.load("fonts/FiraMono-Medium.ttf"),
            font_size: 24.,
            color: Color::WHITE,
        },
    ));
}

fn draw_example_collection(
    mut gizmos: Gizmos,
    mut my_gizmos: Gizmos<MyRoundGizmos>,
    time: Res<Time>,
) {
    let sin = time.elapsed_seconds().sin() * 50.;
    gizmos.line_2d(Vec2::Y * -sin, Vec2::splat(-80.), Color::RED);
    gizmos.ray_2d(Vec2::Y * sin, Vec2::splat(80.), Color::GREEN);

    gizmos.linestrip_gradient_2d([
        (Vec2::Y * 300., Color::BLUE),
        (Vec2::new(-255., -255.), Color::RED),
        (Vec2::new(255., -155.), Color::GREEN),
        (Vec2::Y * 300., Color::BLUE),
    ]);

    gizmos.rect_2d(
        Vec2::ZERO,
        time.elapsed_seconds() / 3.,
        Vec2::splat(300.),
        Color::BLACK,
    );

    my_gizmos.circle_2d(Vec2::ZERO, 120., Color::BLACK);

    my_gizmos.ellipse_2d(
        Vec2::ZERO,
        time.elapsed_seconds() % TAU,
        Vec2::new(100., 200.),
        Color::YELLOW_GREEN,
    );
    my_gizmos
        .circle_2d(Vec2::ZERO, 300., Color::NAVY)
        .segments(64);

    my_gizmos.arc_2d(Vec2::ZERO, sin / 10., PI / 2., 350., Color::ORANGE_RED);

    gizmos.arrow_2d(
        Vec2::ZERO,
        Vec2::from_angle(sin / -10. + PI / 2.) * 50.,
        Color::YELLOW,
    );
}

fn update_config(
    mut config_store: ResMut<GizmoConfigStore>,
    keyboard: Res<ButtonInput<KeyCode>>,
    time: Res<Time>,
) {
    let (config, _) = config_store.config_mut::<DefaultGizmoConfigGroup>();
    if keyboard.pressed(KeyCode::ArrowRight) {
        config.line_width += 5. * time.delta_seconds();
        config.line_width = config.line_width.clamp(0., 50.);
    }
    if keyboard.pressed(KeyCode::ArrowLeft) {
        config.line_width -= 5. * time.delta_seconds();
        config.line_width = config.line_width.clamp(0., 50.);
    }
    if keyboard.just_pressed(KeyCode::Digit1) {
        config.enabled ^= true;
    }

    let (my_config, _) = config_store.config_mut::<MyRoundGizmos>();
    if keyboard.pressed(KeyCode::ArrowUp) {
        my_config.line_width += 5. * time.delta_seconds();
        my_config.line_width = my_config.line_width.clamp(0., 50.);
    }
    if keyboard.pressed(KeyCode::ArrowDown) {
        my_config.line_width -= 5. * time.delta_seconds();
        my_config.line_width = my_config.line_width.clamp(0., 50.);
    }
    if keyboard.just_pressed(KeyCode::Digit2) {
        my_config.enabled ^= true;
    }
}
