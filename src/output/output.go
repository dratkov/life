package output

import (
	"../input"
	"./console"
	"os"
    "os/exec"
)

type Output struct {
	input input.Input
	console console.Console
}

func NewOutput( input input.Input ) *Output {
	return &Output{ input: input }
}

func ( o *Output ) GetInput() input.Input {
	return o.input
}

func ( o *Output ) DoOutput() {
	clearTerminal()
	o.console.DoOutput( o )
}

func ( o *Output ) NextTime() {
	
}

func clearTerminal() {
	cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}