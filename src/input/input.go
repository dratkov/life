package input

import (
	//"flag"
	"./console"
	"../build"
	//"../cell"
	//"fmt"
	"./file/json"
	"reflect"
)

type Input struct {
	build build.Build
	console console.Console
	json json.JSON
}

type GetInputer interface {
	GetInput() Input
}

func ( input *Input ) GetBuild() build.Build {
	return input.build
}

func New() Input {
	return Input{ build: build.New() }
}

func ( input *Input ) GetData() {
	switch {
		case input.console.CheckInput():
			input.console.GetDataAndBuildAreaAndCell( input )
		case input.json.CheckFile():
			input.json.BuildAreaAndCells( input )
	}
}

func ( in *Input ) BuildArea( width, height uint ) {
	in.build.BuildArea( width, height )
}

func ( in *Input ) BuildCell( cell reflect.Type, count int ) {
	in.build.BuildCell( cell, count )
}
