package shark

import (
	"../../../"
	"fmt"
	"../../../../history/move"
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

func New( x int, y int ) *Shark {
	s := &Shark{ cell.New( x, y ) }

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

	history_first_step := move.GetLast(s)
    s.Move()
	history_second_step := move.GetLast(s)
	if history_first_step != nil && history_second_step != nil &&
	history_first_step.GetFromX() == history_second_step.GetToX() &&
	history_first_step.GetFromY() == history_second_step.GetToY() {
		if history_second_step.GetFromY() == history_second_step.GetToY() {
			if history_second_step.GetToX() < history_second_step.GetFromX() {
				s.MoveRight()
			} else {
				s.MoveLeft()
			}
		} else {
			if history_second_step.GetToY() < history_second_step.GetFromY() {
				s.MoveDawn()
			} else {
				s.MoveUp()
			}
		}
	}
}

func ( s *Shark ) String2Console() string {
	return fmt.Sprintf( StringConsoleFormatCell, s.GetConsoleCellSign() )
}

