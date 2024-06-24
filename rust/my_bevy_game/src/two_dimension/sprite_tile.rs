use bevy::prelude::*;

pub fn run() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup)
        .add_systems(Update, (bevy::window::close_on_esc, animate))
        .run();
}

#[derive(Resource)]
struct AnimationState {
    min: f32,
    max: f32,
    current: f32,
    speed: f32,
}

fn setup(mut commands: Commands, aseet_server: Res<AssetServer>) {
    commands.spawn(Camera2dBundle::default());
    commands.insert_resource(AnimationState {
        min: 128.0,
        max: 512.,
        current: 128.,
        speed: 50.,
    });
    commands.spawn((
        SpriteBundle {
            texture: aseet_server.load("branding/icon.png"),
            ..default()
        },
        ImageScaleMode::Tiled {
            tile_x: true,
            tile_y: true,
            stretch_value: 0.5,
        },
    ));
}

fn animate(mut sprites: Query<&mut Sprite>, mut state: ResMut<AnimationState>, time: Res<Time>) {
    if state.current >= state.max || state.current <= state.min {
        state.speed = -state.speed;
    };

    state.current += state.speed * time.delta_seconds();
    for mut sprite in &mut sprites {
        sprite.custom_size = Some(Vec2::splat(state.current));
    }
}
