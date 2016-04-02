package clawn

import (
	"../.."
	"fmt"
	//"reflect"
)

const (
	StringConsoleFormatCell string = "\033[31m%s\033[39m"
	StringConsoleCell string = "."
	MaxLiveCicle int = 40
	MaxLiveCicleDeltaPercent int = 20
)

type Clawn struct {
	cell.Cell
}

func New( x int, y int ) *Clawn {
	c := &Clawn{ cell.New( x, y ) }

	return c
}

func ( c *Clawn ) New( x int, y int ) cell.Celler {
    c.SetX( x )
    c.SetY( y )

    return c
}

func ( c *Clawn ) Init() {
    c.SetMaxLiveCicle( MaxLiveCicle, MaxLiveCicleDeltaPercent )
    c.SetConsoleCellSign( StringConsoleCell )
}

func ( c *Clawn ) DoAction( a cell.Areaer ) {
	c.IncrementLiveCicle()
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

