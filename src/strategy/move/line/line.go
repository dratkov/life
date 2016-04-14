package line

import (
//"../../../area"
    "../../../strategy"
    "../../../history/move"
)

type Line struct {
    area_width uint
}

func (l *Line) Move( c strategy.Celler ) {
    middle_area := int( l.GetAreaWidth() / 2 )
    x := c.GetX()
    last_move := move.GetLast(c)
    var not_to_right, not_to_left, not_to_up, not_to_dawn bool
    if last_move != nil {
        if last_move.GetFromX() != last_move.GetToX() {
            if last_move.GetFromX() < last_move.GetToX() {
                not_to_left = true
            } else {
                not_to_right = true
            }
        } else {
            if last_move.GetFromY() < last_move.GetToY() {
                not_to_up = true
            } else {
                not_to_dawn = true
            }
        }
    }
    if x < middle_area {
       if not_to_right || !c.MoveRight() {
           switch {
                case !not_to_up && c.MoveUp():
                case !not_to_dawn && c.MoveDawn():
                case !not_to_left && c.MoveLeft():
           }
       }
    } else if x > middle_area {
        if not_to_left || !c.MoveLeft() {
            switch {
                case !not_to_up && c.MoveUp():
                case !not_to_dawn && c.MoveDawn():
                case !not_to_right && c.MoveRight():
            }
        }
    }
}

func (l *Line) GetAreaWidth() uint {
    return l.area_width
}

func New( area_width uint ) Line {
    return Line{ area_width: area_width }
}

