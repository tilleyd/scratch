package env

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tilleyd/scratch/stages"
)

type runner struct {
	stage      stages.Stage
	shouldExit bool
}

var globalRunner runner

func init() {
	globalRunner = runner{
		stage:      nil,
		shouldExit: false,
	}

}

func Run(stage stages.Stage) {
	globalRunner.stage = stage

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint)
	rl.InitWindow(1280, 720, "Scratch")

	for !(globalRunner.shouldExit || rl.WindowShouldClose()) {
		rl.BeginDrawing()
		rl.ClearScreenBuffers()
		delta := rl.GetFrameTime()
		globalRunner.stage.Draw(delta)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func Exit() {
	globalRunner.shouldExit = true
}
