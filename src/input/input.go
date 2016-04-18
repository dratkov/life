package input

import (
	"flag"
	"./console"
	"../build"
	//"../cell"
	//"fmt"
	"./file/json"
	"./lua"
	//"../cell"
)

type Input struct {
	build *build.Build
	console console.Console
	json json.JSON
	lua lua.LUA
}

type GetInputer interface {
	GetInput() Input
}

type Initer interface {
	BuildAreaAndCells(build.Builder)
}

func ( input *Input ) GetBuild() *build.Build {
	return input.build
}

func NewInput() Input {
	return Input{ build: build.NewBuilder() }
}

/*
func ( input *Input ) GetData() {
	switch {
		case input.console.IsSourceInit():
			input.console.GetDataAndBuildAreaAndCell(input)
		case input.json.IsSourceInit():
			input.json.BuildAreaAndCells(input)
		case input.lua.IsSourceInit():
			input.lua.BuildAreaAndCells(input)
	}
}
*/

func (input *Input) InitFromSource()  {
	var flagJSONFile, flagLUAScript string
	flag.StringVar(&flagJSONFile, "init-json-file", "", "JSON file name for init")
	flag.StringVar(&flagLUAScript, "init-lua-script", "", "LUA script name for init")
	flag.Parse()
	var initer Initer
	switch {
	case flagLUAScript != "":
		initer = lua.NewLua(&flagLUAScript)
	case flagJSONFile != "":
		initer = json.NewJSON(&flagJSONFile)
	default:
		initer = console.New()
	}

	initer.BuildAreaAndCells(input)
}

func (in *Input) BuildArea( width, height uint ) {
	in.build.BuildArea( width, height )
}

func (in *Input) BuildCell( cell interface{}, count int ) {
	in.build.BuildCell( cell, count )

	//return c
}
