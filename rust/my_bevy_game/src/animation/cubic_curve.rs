use bevy::{
    math::{cubic_splines::CubicCurve, vec3},
    prelude::*,
};

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup)
        .add_systems(Update, (animate_cube, bevy::window::close_on_esc))
        .run();
}

#[derive(Component)]
struct Curve(CubicCurve<Vec3>);

fn setup(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<StandardMaterial>>,
) {
    let points = [[
        vec3(-6., 2., 0.),
        vec3(12., 8., 0.),
        vec3(-12., 8., 0.),
        vec3(6., 2., 0.),
    ]];

    let bezier = CubicBezier::new(points).to_curve();

    commands.spawn((
        PbrBundle {
            mesh: meshes.add(Cuboid::default()),
            material: materials.add(Color::ORANGE),
            transform: Transform::from_translation(points[0][0]),
            ..default()
        },
        Curve(bezier),
    ));

    commands.spawn(PointLightBundle {
        point_light: PointLight {
            shadows_enabled: true,
            intensity: 10_000_000.,
            range: 100.,
            ..default()
        },
        transform: Transform::from_xyz(8., 16., 8.),
        ..default()
    });

    commands.spawn(PbrBundle {
        mesh: meshes.add(Plane3d::default().mesh().size(50., 50.)),
        material: materials.add(Color::SILVER),
        ..default()
    });

    commands.spawn(Camera3dBundle {
        transform: Transform::from_xyz(0., 6., 12.).looking_at(Vec3::new(0., 3., 0.), Vec3::Y),
        ..default()
    });
}

fn animate_cube(time: Res<Time>, mut query: Query<(&mut Transform, &Curve)>, mut gizmos: Gizmos) {
    let t = (time.elapsed_seconds().sin() + 1.) / 2.;

    for (mut transform, cubic_curve) in &mut query {
        gizmos.linestrip(cubic_curve.0.iter_positions(50), Color::WHITE);
        transform.translation = cubic_curve.0.position(t);
    }
}
