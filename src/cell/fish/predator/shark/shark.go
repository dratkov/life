package shark

import (
	"../../../"
	"fmt"
)

const (
	StringConsoleFormatCell string = "\033[34m%s\033[39m"
	StringConsoleCell string = "-"
	MaxLiveCicle int = 100
	MaxLiveCicleDeltaPercent int = 17
)

type Shark struct {
	cell.Cell
}

func New( x int, y int ) Shark {
	s := Shark{ cell.New( x, y ) }

	return s
}

func ( s *Shark ) New( x int, y int ) cell.Celler {
    s.SetX( x )
    s.SetY( y )
 
    return s
}

func ( s *Shark ) Init() {
    s.SetMaxLiveCicle( MaxLiveCicle, MaxLiveCicleDeltaPercent )
    s.SetConsoleCellSign(StringConsoleCell)
}

func ( s *Shark ) DoAction( a cell.Areaer ) {    
	s.IncrementLiveCicle()
    s.Move()
    //s.Move()
}

func ( s *Shark ) String2Console() string {
	return fmt.Sprintf( StringConsoleFormatCell, s.GetConsoleCellSign() )
}

