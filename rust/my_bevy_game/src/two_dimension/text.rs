use bevy::{
    prelude::*,
    sprite::Anchor,
    text::{BreakLineOn, Text2dBounds},
};

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup)
        .add_systems(
            Update,
            (
                animate_translation,
                animate_rotation,
                animate_scale,
                bevy::window::close_on_esc,
            ),
        )
        .run();
}

#[derive(Component)]
struct AnimateTranslation;

#[derive(Component)]
struct AnimateRotation;

#[derive(Component)]
struct AnimateScale;

fn setup(mut commands: Commands, asset_server: Res<AssetServer>) {
    let font = asset_server.load("fonts/FiraSans-Bold.ttf");
    let text_style = TextStyle {
        font: font.clone(),
        font_size: 60.,
        color: Color::WHITE,
    };
    let text_justification = JustifyText::Center;
    commands.spawn(Camera2dBundle::default());
    commands.spawn((
        Text2dBundle {
            text: Text::from_section("translation", text_style.clone())
                .with_justify(text_justification),
            ..default()
        },
        AnimateTranslation,
    ));
    commands.spawn((
        Text2dBundle {
            text: Text::from_section("rotation", text_style.clone())
                .with_justify(text_justification),
            ..default()
        },
        AnimateRotation,
    ));
    commands.spawn((
        Text2dBundle {
            text: Text::from_section("scale", text_style).with_justify(text_justification),
            ..default()
        },
        AnimateScale,
    ));
    let slightly_smaller_text_style = TextStyle {
        font,
        font_size: 42.,
        color: Color::WHITE,
    };
    let box_size = Vec2::new(300., 200.);
    let box_position = Vec2::new(0., -250.);
    commands
        .spawn(SpriteBundle {
            sprite: Sprite {
                color: Color::rgb(0.25, 0.25, 0.75),
                custom_size: Some(Vec2::new(box_size.x, box_size.y)),
                ..default()
            },
            transform: Transform::from_translation(box_position.extend(0.)),
            ..default()
        })
        .with_children(|builder| {
            builder.spawn(Text2dBundle {
                text: Text {
                    sections: vec![TextSection::new(
                        "this text wraps in the box\n(Unicode linebreaks)",
                        slightly_smaller_text_style.clone(),
                    )],
                    justify: JustifyText::Left,
                    linebreak_behavior: BreakLineOn::WordBoundary,
                },
                text_2d_bounds: Text2dBounds { size: box_size },
                transform: Transform::from_translation(Vec3::Z),
                ..default()
            });
        });

    let other_box_size = Vec2::new(300., 200.);
    let other_box_position = Vec2::new(320., -250.);
    commands
        .spawn(SpriteBundle {
            sprite: Sprite {
                color: Color::rgb(0.20, 0.3, 0.7),
                custom_size: Some(Vec2::new(other_box_size.x, other_box_size.y)),
                ..default()
            },
            transform: Transform::from_translation(other_box_position.extend(0.)),
            ..default()
        })
        .with_children(|builder| {
            builder.spawn(Text2dBundle {
                text: Text {
                    sections: vec![TextSection::new(
                        "this text wraps in the box\n(AnyCharater linebreaks)",
                        slightly_smaller_text_style.clone(),
                    )],
                    justify: JustifyText::Left,
                    linebreak_behavior: BreakLineOn::AnyCharacter,
                },
                text_2d_bounds: Text2dBounds {
                    size: other_box_size,
                },
                transform: Transform::from_translation(Vec3::Z),
                ..default()
            });
        });

    for (text_anchor, color) in [
        (Anchor::TopLeft, Color::RED),
        (Anchor::TopRight, Color::GREEN),
        (Anchor::BottomRight, Color::BLUE),
        (Anchor::BottomLeft, Color::YELLOW),
    ] {
        commands.spawn(Text2dBundle {
            text: Text {
                sections: vec![TextSection::new(
                    format!(" Anchor::{text_anchor:?} "),
                    TextStyle {
                        color,
                        ..slightly_smaller_text_style.clone()
                    },
                )],
                ..default()
            },
            transform: Transform::from_translation(250. * Vec3::Y),
            text_anchor,
            ..default()
        });
    }
}

fn animate_translation(
    time: Res<Time>,
    mut query: Query<&mut Transform, (With<Text>, With<AnimateTranslation>)>,
) {
    for mut transform in &mut query {
        transform.translation.x = 100. * time.elapsed_seconds().sin() - 400.;
        transform.translation.y = 100. * time.elapsed_seconds().cos();
    }
}

fn animate_rotation(
    time: Res<Time>,
    mut query: Query<&mut Transform, (With<Text>, With<AnimateRotation>)>,
) {
    for mut transform in &mut query {
        transform.rotation = Quat::from_rotation_z(time.elapsed_seconds().cos());
    }
}

fn animate_scale(
    time: Res<Time>,
    mut query: Query<&mut Transform, (With<Text>, With<AnimateScale>)>,
) {
    for mut transform in &mut query {
        transform.translation = Vec3::new(400., 0., 0.);

        let scale = (time.elapsed_seconds().sin() + 1.1) * 2.0;
        transform.scale.x = scale;
        transform.scale.y = scale;
    }
}
