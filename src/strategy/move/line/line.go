package line

import (
//"../../../area"
    "../../../strategy"
)

type Line struct {
    area_width uint
}

func (l *Line) Move( c strategy.Celler ) {
    middle_area := int( l.GetAreaWidth() / 2 )
    x := c.GetX()
    if x < middle_area {
       if c.MoveRight() == false {
           switch {
                case c.MoveUp() == true:
                case c.MoveDawn() == true:
           }
       }
    } else if x > middle_area {
        if c.MoveLeft() == false {
            switch {
                case c.MoveUp() == true:
                case c.MoveDawn() == true:
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

