package main

import (
	"flag"
	"fmt"

	"github.com/tilleyd/scratch/env"
	"github.com/tilleyd/scratch/stages"
	"github.com/tilleyd/scratch/stages/cube"
)

func main() {
	flag.Parse()
	stageName := flag.Arg(0)
	var stage stages.Stage

	switch stageName {
	case "cube":
		stage = cube.NewCubeStage()
	default:
		panic(fmt.Errorf("unknown stage '%s'", stageName))
	}

	env.Run(stage)
}
