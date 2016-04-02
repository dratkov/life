package cell

import (
    "../history/move"
)

type MoveDirectFunc func( Areaer ) bool

func ( c *Cell ) Move() {
    strategy := c.GetStrategy()
    strategy.Move(c)
}


func ( c *Cell ) DoMove( target_x, target_y int ) {
    var x, y = c.GetX(), c.GetY()
    //area := c.GetArea()
    target_cell := area.GetCell( target_x, target_y )
    c.SetX( target_x )
    c.SetY( target_y )
    target_cell.SetX( x )
    target_cell.SetY( y )
    move.New(c, x, y, target_x, target_y)
}

func ( c *Cell ) MoveLeft() bool {
    target_x := c.GetX() - 1
    //area := c.GetArea()
    if target_x == -1 {
        return false
    }
    target_y := c.GetY()
    target_cell := area.GetCell( target_x, target_y )
    if target_cell.IsEmptyCell( target_cell ) {
        c.DoMove( target_x, target_y )
        return true
    }

    return false
}

func ( c *Cell ) MoveUp() bool {
    target_y := c.GetY() - 1
    //area := c.GetArea()
    if target_y == -1 {
        return false
    }
    target_x := c.GetX()
    target_cell := area.GetCell( target_x, target_y )
    if target_cell.IsEmptyCell( target_cell ) {
        c.DoMove( target_x, target_y )
        return true
    }

    return false
}

func ( c *Cell ) MoveRight() bool {
    target_x := c.GetX() + 1
    //area := c.GetArea()
    if target_x == int(area.GetWidth()) {
        return false
    }
    target_y := c.GetY()
    target_cell := area.GetCell( target_x, target_y )
    if target_cell.IsEmptyCell( target_cell ) {
        c.DoMove( target_x, target_y )
        return true
    }

    return false
}

func ( c *Cell ) MoveDawn() bool {
    target_y := c.GetY() + 1
    //area := c.GetArea()
    if target_y == int(area.GetHeight()) {
        return false
    }
    target_x := c.GetX()
    target_cell := area.GetCell( target_x, target_y )
    if target_cell.IsEmptyCell( target_cell ) {
        c.DoMove( target_x, target_y )
        return true
    }

    return false
}
