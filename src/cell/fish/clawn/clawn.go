package clawn

import (
	"../.."
	"fmt"
	//"reflect"
)

const (
	StringConsoleFormatCell string = "\033[31m%s\033[39m"
	StringConsoleCell string = "."
	MaxLiveCycle int = 40
	MaxLiveCycleDeltaPercent int = 20
)

type Clawn struct {
	*cell.Cell
}

func NewClawn( x int, y int ) *Clawn {
	c := &Clawn{ cell.NewCell( x, y ) }

	return c
}

func ( c *Clawn ) New( x int, y int ) cell.Celler {
    c.SetX( x )
    c.SetY( y )

    return c
}

func ( c *Clawn ) Init() {
    c.SetMaxLiveCycle( MaxLiveCycle, MaxLiveCycleDeltaPercent )
    c.SetConsoleCellSign( StringConsoleCell )
}

func (c *Clawn) DoAction( celler cell.Celler ) {
	c.Cell.DoAction(c)
	c.Move()
	/*x := c.GetX()
	if x > 0 {
    	c.SetX( x - 0 )
	}
	*/
}

func ( c *Clawn ) String2Console() string {
	return fmt.Sprintf( StringConsoleFormatCell, c.GetConsoleCellSign() )
}

