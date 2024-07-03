use bevy::prelude::*;
use std::f32::consts::PI;

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .insert_resource(AmbientLight {
            color: Color::WHITE,
            brightness: 150.,
        })
        .add_systems(Startup, (setup, bevy::window::close_on_esc))
        .run();
}

fn setup(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<StandardMaterial>>,
    mut animations: ResMut<Assets<AnimationClip>>,
) {
    commands.spawn(Camera3dBundle {
        transform: Transform::from_xyz(-2., 2.5, 5.).looking_at(Vec3::ZERO, Vec3::Y),
        ..default()
    });

    commands.spawn(PointLightBundle {
        point_light: PointLight {
            intensity: 500_000.0,
            ..default()
        },
        transform: Transform::from_xyz(0., 2.5, 0.),
        ..default()
    });

    let planet = Name::new("planet");
    let orbit_controller = Name::new("orbit_controller");
    let satellite = Name::new("satellite");

    let mut animation = AnimationClip::default();

    animation.add_curve_to_path(
        EntityPath {
            parts: vec![planet.clone()],
        },
        VariableCurve {
            keyframe_timestamps: vec![0., 1., 2., 3., 4.],
            keyframes: Keyframes::Translation(vec![
                Vec3::new(1., 0., 1.),
                Vec3::new(-1., 0., 1.),
                Vec3::new(-1., 0., -1.),
                Vec3::new(1., 0., -1.),
                Vec3::new(1., 0., 1.),
            ]),
            interpolation: Interpolation::Linear,
        },
    );

    animation.add_curve_to_path(
        EntityPath {
            parts: vec![planet.clone(), orbit_controller.clone()],
        },
        VariableCurve {
            keyframe_timestamps: vec![0., 1., 2., 3., 4.],
            keyframes: Keyframes::Rotation(vec![
                Quat::IDENTITY,
                Quat::from_axis_angle(Vec3::Y, PI / 2.),
                Quat::from_axis_angle(Vec3::Y, PI / 2. * 2.),
                Quat::from_axis_angle(Vec3::Y, PI / 2. * 3.),
                Quat::IDENTITY,
            ]),
            interpolation: Interpolation::Linear,
        },
    );

    animation.add_curve_to_path(
        EntityPath {
            parts: vec![planet.clone(), orbit_controller.clone(), satellite.clone()],
        },
        VariableCurve {
            keyframe_timestamps: vec![0., 0.5, 1., 1.5, 2., 2.5, 3., 3.5, 4.],
            keyframes: Keyframes::Scale(vec![
                Vec3::splat(0.8),
                Vec3::splat(1.2),
                Vec3::splat(0.8),
                Vec3::splat(1.2),
                Vec3::splat(0.8),
                Vec3::splat(1.2),
                Vec3::splat(0.8),
                Vec3::splat(1.2),
                Vec3::splat(0.8),
            ]),
            interpolation: Interpolation::Linear,
        },
    );

    animation.add_curve_to_path(
        EntityPath {
            parts: vec![planet.clone(), orbit_controller.clone(), satellite.clone()],
        },
        VariableCurve {
            keyframe_timestamps: vec![0., 1., 2., 3., 4.],
            keyframes: Keyframes::Rotation(vec![
                Quat::IDENTITY,
                Quat::from_axis_angle(Vec3::Y, PI / 2.),
                Quat::from_axis_angle(Vec3::Y, PI / 2. * 2.),
                Quat::from_axis_angle(Vec3::Y, PI / 2. * 3.),
                Quat::IDENTITY,
            ]),
            interpolation: Interpolation::Linear,
        },
    );

    let mut player = AnimationPlayer::default();
    player.play(animations.add(animation)).repeat();

    commands
        .spawn((
            PbrBundle {
                mesh: meshes.add(Sphere::default()),
                material: materials.add(Color::rgb(0.8, 0.7, 0.6)),
                ..default()
            },
            planet,
            player,
        ))
        .with_children(|p| {
            p.spawn((SpatialBundle::INHERITED_IDENTITY, orbit_controller))
                .with_children(|p| {
                    p.spawn((
                        PbrBundle {
                            transform: Transform::from_xyz(1.5, 0., 0.),
                            mesh: meshes.add(Cuboid::new(0.5, 0.5, 0.5)),
                            material: materials.add(Color::rgb(0.3, 0.9, 0.3)),
                            ..default()
                        },
                        satellite,
                    ));
                });
        });
}
