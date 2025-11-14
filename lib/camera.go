package lib

import rl "github.com/gen2brain/raylib-go/raylib"

// Default 3D camera with sensible defaults for orbiting around the center.
func NewCamera3DDefault(distance float32) rl.Camera3D {
	return rl.NewCamera3D(
		rl.NewVector3(distance, distance, distance),
		rl.Vector3Zero(),
		rl.NewVector3(0, 1, 0),
		90,
		rl.CameraPerspective,
	)
}

// Rotates and zooms the camera around its target when mouse buttons are held.
func CameraOrbit(camera *rl.Camera) {
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		rl.CameraYaw(camera, rl.GetMouseDelta().X*-0.01, 1)
		rl.CameraPitch(camera, rl.GetMouseDelta().Y*-0.01, 1, 1, 0)
	}
	if rl.IsMouseButtonDown(rl.MouseButtonRight) {
		rl.CameraMoveToTarget(camera, rl.GetMouseDelta().Y*0.1)
	}
}
