use bevy::prelude::*;

const BOUNDS: Vec2 = Vec2::new(1200., 640.);

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .insert_resource(Time::<Fixed>::from_hz(60.))
        .add_systems(Startup, setup)
        .add_systems(
            FixedUpdate,
            (
                player_movement_system,
                snap_to_player_system,
                rotate_to_player_system,
            ),
        )
        .add_systems(Update, bevy::window::close_on_esc)
        .run();
}

#[derive(Component)]
struct Player {
    movement_speed: f32,
    rotation_speed: f32,
}

#[derive(Component)]
struct SnapToPlayer;

#[derive(Component)]
struct RotateToPlayer {
    rotation_speed: f32,
}

fn setup(mut commands: Commands, asset_server: Res<AssetServer>) {
    let ship_handle = asset_server.load("textures/simplespace/ship_C.png");
    let enemy_a_handle = asset_server.load("textures/simplespace/enemy_A.png");
    let enemy_b_handle = asset_server.load("textures/simplespace/enemy_B.png");

    commands.spawn(Camera2dBundle::default());

    let horizontal_margin = BOUNDS.x / 4.;
    let vertical_margin = BOUNDS.y / 4.;

    commands.spawn((
        SpriteBundle {
            texture: ship_handle,
            ..default()
        },
        Player {
            movement_speed: 500.,
            rotation_speed: f32::to_radians(360.),
        },
    ));

    commands.spawn((
        SpriteBundle {
            texture: enemy_a_handle.clone(),
            transform: Transform::from_xyz(0. - horizontal_margin, 0., 0.),
            ..default()
        },
        SnapToPlayer,
    ));

    commands.spawn((
        SpriteBundle {
            texture: enemy_a_handle.clone(),
            transform: Transform::from_xyz(0., 0. - vertical_margin, 0.),
            ..default()
        },
        SnapToPlayer,
    ));

    commands.spawn((
        SpriteBundle {
            texture: enemy_b_handle.clone(),
            transform: Transform::from_xyz(0. + horizontal_margin, 0., 0.),
            ..default()
        },
        RotateToPlayer {
            rotation_speed: f32::to_radians(45.),
        },
    ));
    commands.spawn((
        SpriteBundle {
            texture: enemy_b_handle,
            transform: Transform::from_xyz(0., 0. + vertical_margin, 0.),
            ..default()
        },
        RotateToPlayer {
            rotation_speed: f32::to_radians(90.),
        },
    ));
}

fn player_movement_system(
    time: Res<Time>,
    keyboard_input: Res<ButtonInput<KeyCode>>,
    mut query: Query<(&Player, &mut Transform)>,
) {
    let (ship, mut transform) = query.single_mut();

    let mut rotation_factor = 0.;
    let mut movement_factor = 0.;

    if keyboard_input.pressed(KeyCode::ArrowLeft) {
        rotation_factor += 1.0;
    }

    if keyboard_input.pressed(KeyCode::ArrowRight) {
        rotation_factor -= 1.0;
    }

    if keyboard_input.pressed(KeyCode::ArrowUp) {
        movement_factor += 1.0;
    }

    transform.rotate_z(rotation_factor * ship.rotation_speed * time.delta_seconds());

    let movement_direction = transform.rotation * Vec3::Y;

    let movement_distance = movement_factor * ship.movement_speed * time.delta_seconds();

    let translation_delta = movement_direction * movement_distance;
    transform.translation += translation_delta;

    let extents = Vec3::from((BOUNDS / 2., 0.));
    transform.translation = transform.translation.min(extents).max(-extents);
}

fn snap_to_player_system(
    mut query: Query<&mut Transform, (With<SnapToPlayer>, Without<Player>)>,
    player_query: Query<&Transform, With<Player>>,
) {
    let player_transform = player_query.single();
    let player_translation = player_transform.translation.xy();

    for mut enemy_transform in &mut query {
        let to_player = (player_translation - enemy_transform.translation.xy()).normalize();

        let rotate_to_player = Quat::from_rotation_arc(Vec3::Y, to_player.extend(0.));

        enemy_transform.rotation = rotate_to_player;
    }
}

fn rotate_to_player_system(
    time: Res<Time>,
    mut query: Query<(&RotateToPlayer, &mut Transform), Without<Player>>,
    player_query: Query<&Transform, With<Player>>,
) {
    let player_transform = player_query.single();
    let player_translation = player_transform.translation.xy();

    for (config, mut enemy_transform) in &mut query {
        let enemy_forward = (enemy_transform.rotation * Vec3::Y).xy();

        let to_player = (player_translation - enemy_transform.translation.xy()).normalize();

        let forward_dot_player = enemy_forward.dot(to_player);

        if (forward_dot_player - 1.).abs() < f32::EPSILON {
            continue;
        }

        let enemy_right = (enemy_transform.rotation * Vec3::X).xy();

        let right_dot_player = enemy_right.dot(to_player);

        let rotation_sign = -f32::copysign(1., right_dot_player);

        let max_angle = forward_dot_player.clamp(-1., 1.).acos();

        let rotation_angle =
            rotation_sign * (config.rotation_speed * time.delta_seconds()).min(max_angle);

        enemy_transform.rotate_z(rotation_angle);
    }
}
