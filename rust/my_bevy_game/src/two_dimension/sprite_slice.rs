use bevy::prelude::*;

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins.set(WindowPlugin {
            primary_window: Some(Window {
                resolution: (1350., 700.).into(),
                ..default()
            }),
            ..default()
        }))
        .add_systems(Startup, setup)
        .add_systems(Update, bevy::window::close_on_esc)
        .run();
}

fn spawn_sprites(
    commands: &mut Commands,
    texture_handle: Handle<Image>,
    mut position: Vec3,
    slice_border: f32,
    style: TextStyle,
    gap: f32,
) {
    let cases = [
        ("Original texture", style.clone(), Vec2::splat(100.), None),
        (
            "Streched texture",
            style.clone(),
            Vec2::new(100., 200.),
            None,
        ),
        (
            "Streched and sliced",
            style.clone(),
            Vec2::new(100., 200.),
            Some(ImageScaleMode::Sliced(TextureSlicer {
                border: BorderRect::square(slice_border),
                center_scale_mode: SliceScaleMode::Stretch,
                ..default()
            })),
        ),
        (
            "Scaled and Tiled",
            style.clone(),
            Vec2::new(100., 200.),
            Some(ImageScaleMode::Sliced(TextureSlicer {
                border: BorderRect::square(slice_border),
                center_scale_mode: SliceScaleMode::Tile { stretch_value: 0.5 },
                sides_scale_mode: SliceScaleMode::Tile { stretch_value: 0.2 },
                ..default()
            })),
        ),
        (
            "Sliced and Tiled",
            style.clone(),
            Vec2::new(300., 200.),
            Some(ImageScaleMode::Sliced(TextureSlicer {
                border: BorderRect::square(slice_border),
                center_scale_mode: SliceScaleMode::Tile { stretch_value: 0.2 },
                sides_scale_mode: SliceScaleMode::Tile { stretch_value: 0.3 },
                ..default()
            })),
        ),
        (
            "Sliced and Tiled with corner constraint",
            style,
            Vec2::new(300., 200.),
            Some(ImageScaleMode::Sliced(TextureSlicer {
                border: BorderRect::square(slice_border),
                center_scale_mode: SliceScaleMode::Tile { stretch_value: 0.1 },
                sides_scale_mode: SliceScaleMode::Tile { stretch_value: 0.2 },
                max_corner_scale: 0.2,
            })),
        ),
    ];

    for (label, text_style, size, scale_mode) in cases {
        position.x += 0.5 * size.x;
        let mut cmd = commands.spawn(SpriteBundle {
            transform: Transform::from_translation(position),
            texture: texture_handle.clone(),
            sprite: Sprite {
                custom_size: Some(size),
                ..default()
            },
            ..default()
        });
        if let Some(scale_mode) = scale_mode {
            cmd.insert(scale_mode);
        }
        cmd.with_children(|builder| {
            builder.spawn(Text2dBundle {
                text: Text::from_section(label, text_style).with_justify(JustifyText::Center),
                transform: Transform::from_xyz(0., -0.5 * size.y - 10., 0.),
                text_anchor: bevy::sprite::Anchor::TopCenter,
                ..default()
            });
        });
        position.x += 0.5 * size.x + gap;
    }
}

fn setup(mut commands: Commands, asset_server: Res<AssetServer>) {
    commands.spawn(Camera2dBundle::default());
    let font = asset_server.load("fonts/FiraSans-Bold.ttf");
    let style = TextStyle {
        font: font.clone(),
        font_size: 16.,
        color: Color::WHITE,
    };

    let handle_1 = asset_server.load("textures/slice_square.png");
    let handle_2 = asset_server.load("textures/slice_square_2.png");

    spawn_sprites(
        &mut commands,
        handle_1,
        Vec3::new(-600., 200., 0.),
        200.,
        style.clone(),
        50.,
    );
    spawn_sprites(
        &mut commands,
        handle_2,
        Vec3::new(-600., -200., 0.),
        80.,
        style,
        50.,
    );
}
