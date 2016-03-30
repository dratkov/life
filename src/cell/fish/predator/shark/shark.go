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

	last_move := move.GetLast(s)
    s.Move()
	last_move_2 := move.GetLast(s)
	if last_move != nil && last_move_2 != nil &&
			last_move.GetFromX() == last_move_2.GetToX() &&
			last_move.GetFromY() == last_move_2.GetToY() {
		if last_move_2.GetToY() == last_move_2.GetFromY() {
			if last_move_2.GetToX() < last_move_2.GetFromX() {
				s.MoveRight()
			} else {
				s.MoveLeft()
			}
		} else {
			if last_move_2.GetToY() < last_move_2.GetFromY() {
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

