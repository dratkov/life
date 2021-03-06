package shark

import (
	"../../../"
	"fmt"
)

const (
	StringConsoleFormatCell string = "\033[34m%s\033[39m"
	StringConsoleCell string = "-"
	MaxLiveCycle int = 100
	MaxLiveCycleDeltaPercent int = 17
)

type Shark struct {
	*cell.Cell
}

func NewShark( x int, y int ) *Shark {
	s := &Shark{ cell.NewCell( x, y ) }

	return s
}

func ( s *Shark ) New( x int, y int ) cell.Celler {
    s.SetX( x )
    s.SetY( y )
 
    return s
}

func ( s *Shark ) Init() {
    s.SetMaxLiveCycle( MaxLiveCycle, MaxLiveCycleDeltaPercent )
    s.SetConsoleCellSign(StringConsoleCell)
}

func (s *Shark ) DoAction( c cell.Celler ) {
	s.Cell.DoAction(s)
	s.Move(c)
/*
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
	*/
}

func (s *Shark ) Move( c cell.Celler ) {
	s.Cell.Move()
	s.Cell.Move()
}


func ( s *Shark ) String2Console() string {
	return fmt.Sprintf( StringConsoleFormatCell, s.GetConsoleCellSign() )
}

