package build

import (
    "../area"
    "../cell"
    "../cell/fish/clawn"
    "../cell/fish/predator/shark"
    //"fmt"
    "reflect"
    "../strategy"
    "../strategy/move/random"
    "../strategy/move/line"
)

type Build struct {
    area area.Area
    //cell []cell.Cell
    //clawn []clawn.Clawn

}

type Builder interface {
    BuildArea( uint, uint )
    BuildCell( reflect.Type, int )
}

func New() Build {
    return Build{}
}

func ( b *Build ) BuildArea( width, height uint ) {
    b.area = area.New( width, height )
    b.FillArea()
}

func ( b *Build ) GetArea() *area.Area {
    return &b.area
    a := b.area
    return &a
}

func ( b *Build ) BuildCell( c reflect.Type, count int ) {
    area := b.GetArea()
    for i := 0; i < count; i++ {
        var free_cell = area.GetRandFreeCell()
        if free_cell != nil {
            cellPtr := reflect.New( c )
            var cell_general cell.Celler
            var move_strategy strategy.MoveStrateger
            switch cellPtr.Elem().Interface().(type) {
                case clawn.Clawn:
                    r := cellPtr.Elem().Interface().(clawn.Clawn)
                    cell_general = &r
                    ms := random.New()
                    move_strategy = &ms
                case shark.Shark:
                    r := cellPtr.Elem().Interface().(shark.Shark)
                    cell_general = &r
                    ms := line.New( area.GetWidth() )
                    move_strategy = &ms
            }
            cell_general.SetX( free_cell.GetX() )
            cell_general.SetY( free_cell.GetY() )
            cell_general.Init()
            cell_general.SetStrategy( strategy.New( move_strategy ) )
            cell_general.SetArea( area )
            area.AddCell( cell_general )
        }
    }
}

func ( b Build ) FillArea() {
    area := b.GetArea()
    height, width := area.GetHeight(), area.GetWidth()
    for i := 0; i < int(height); i++ {
        for j := 0; j < int(width); j++ {
            cell_el := cell.New( j, i )
            area.AddCell( &cell_el )
        }
    }
}