use bevy::{
    gltf::GltfPlugin,
    prelude::*,
    reflect::TypePath,
    render::{
        mesh::{MeshVertexAttribute, MeshVertexBufferLayout},
        render_resource::*,
    },
    sprite::{Material2d, Material2dKey, Material2dPlugin, MaterialMesh2dBundle, Mesh2dHandle},
};

const ATTRIBUTE_BARYCENTRIC: MeshVertexAttribute =
    MeshVertexAttribute::new("Barycentric", 2139464976, VertexFormat::Float32x3);

pub fn run() {
    App::new()
        .insert_resource(AmbientLight {
            color: Color::WHITE,
            brightness: 1.0 / 5.0f32,
        })
        .add_plugins((
            DefaultPlugins.set(
                GltfPlugin::default()
                    .add_custom_vertex_attribute("_BARYCENTRIC", ATTRIBUTE_BARYCENTRIC),
            ),
            Material2dPlugin::<CustomMaterial>::default(),
        ))
        .add_systems(Startup, setup)
        .run();
}

fn setup(
    mut commands: Commands,
    asset_server: Res<AssetServer>,
    mut materials: ResMut<Assets<CustomMaterial>>,
) {
    let mesh = asset_server.load("models/barycentric/barycentric.gltf#Mesh0/Primitive0");
    commands.spawn(MaterialMesh2dBundle {
        mesh: Mesh2dHandle(mesh),
        material: materials.add(CustomMaterial {}),
        transform: Transform::from_scale(150.0 * Vec3::ONE),
        ..default()
    });

    commands.spawn(Camera2dBundle { ..default() });
}

#[derive(Asset, TypePath, AsBindGroup, Debug, Clone)]
struct CustomMaterial {}

impl Material2d for CustomMaterial {
    fn vertex_shader() -> ShaderRef {
        "shaders/custom_gltf_2d.wgsl".into()
    }
    fn fragment_shader() -> ShaderRef {
        "shaders/custom_gltf_2d.wgsl".into()
    }

    fn specialize(
        descriptor: &mut RenderPipelineDescriptor,
        layout: &MeshVertexBufferLayout,
        _key: Material2dKey<Self>,
    ) -> Result<(), SpecializedMeshPipelineError> {
        let vertex_layout = layout.get_layout(&[
            Mesh::ATTRIBUTE_POSITION.at_shader_location(0),
            Mesh::ATTRIBUTE_COLOR.at_shader_location(1),
            ATTRIBUTE_BARYCENTRIC.at_shader_location(2),
        ])?;
        descriptor.vertex.buffers = vec![vertex_layout];
        Ok(())
    }
}
