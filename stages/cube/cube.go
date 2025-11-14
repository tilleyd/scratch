package cube

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tilleyd/scratch/lib"
)

const vertexShader string = `
#version 330

in vec3 v_Position;
in vec2 v_TexCoord;
in vec3 v_Normal;

uniform mat4 mvp;
uniform mat4 matModel;
uniform mat4 matNormal;

out vec3 f_Position;
out vec2 f_TexCoord;
out vec3 f_Normal;

void main() {
	f_Position = vec3(matModel*vec4(v_Position, 1));
	f_TexCoord = v_TexCoord;
	f_Normal = normalize(vec3(matNormal*vec4(v_Normal, 1)));
	gl_Position = mvp*vec4(v_Position, 1);
}
`

const fragmentShader string = `
#version 330

in vec3 f_Position;
in vec2 f_TexCoord;
in vec3 f_Normal;

uniform sampler2D texture0;

out vec4 o_Color;

void main() {
	o_Color = texture(texture0, f_TexCoord);
}
`

type CubeStage struct {
	mesh     rl.Mesh
	texture  rl.Texture2D
	material rl.Material
	camera   rl.Camera3D
}

func NewCubeStage() *CubeStage {
	return &CubeStage{}
}

func (s *CubeStage) Setup() {
	s.mesh = rl.GenMeshCube(1, 1, 1)

	s.material = rl.LoadMaterialDefault()
	image := rl.GenImageChecked(2, 2, 1, 1, rl.White, rl.Red)
	defer rl.UnloadImage(image)

	s.texture = rl.LoadTextureFromImage(image)

	rl.SetMaterialTexture(&s.material, rl.MapDiffuse, s.texture)

	s.camera = lib.NewCamera3DDefault(5)

	shader := rl.LoadShaderFromMemory(vertexShader, fragmentShader)
	s.material.Shader = shader
}

func (s *CubeStage) End() {
	rl.UnloadMaterial(s.material)
	rl.UnloadTexture(s.texture)
	rl.UnloadMesh(&s.mesh)
}

func (s *CubeStage) Draw(delta float32) {
	lib.CameraOrbit(&s.camera)

	rl.BeginMode3D(s.camera)
	transform := rl.MatrixIdentity()
	rl.DrawMesh(s.mesh, s.material, transform)
	rl.DrawGrid(8, 1)
	rl.EndMode3D()
}
