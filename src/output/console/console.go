package console

import (
	"fmt"
	"../../input"
	"../../area"
)

type Console struct {

}

func New() Console {
	return Console{}
}

func ( c *Console ) DoOutput( inputer input.GetInputer ) {
	input := inputer.GetInput()
	build := input.GetBuild()
	c.DrawArea( build.GetArea() )
}

func ( c *Console ) DrawArea( a *area.Area ) {
	height, width := a.GetHeight(), a.GetWidth()
	fmt_border := "\033[34m%s\033[39m"
	for i := uint(0); i < height; i++ {
		fmt.Printf( fmt_border, "|" )
		for j := uint(0); j < width; j++ {
			cell := a.GetCell( int(j), int(i) )
			//if cell != nil {
				fmt.Printf( "%s", cell.String2Console() )
			//}// else {
			//	fmt.Printf( "+" )
			//}
			fmt.Printf( fmt_border, "|" )
		}
		fmt.Println()
	}
	//fmt.Println()
}