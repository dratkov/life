package build

import (
    "../area"
    "../cell"
    "../cell/fish/clawn"
    "../cell/fish/predator/shark"
    //"fmt"
    "../strategy"
    "../strategy/move/random"
    "../strategy/move/line"
)

type Build struct {
    area *area.Area
}

type Builder interface {
    BuildArea( uint, uint )
    BuildCell( interface{}, int )
}

func New() Build {
    return Build{}
}

func ( b *Build ) BuildArea( width, height uint ) {
    b.area = area.New( width, height )
    cell.SetArea(b.area)
    b.FillArea()
}

func ( b *Build ) GetArea() *area.Area {
    return b.area
    a := b.area
    return a
}

func ( b *Build ) BuildCell( c interface{}, count int ) {
    area := b.GetArea()
    for i := 0; i < count; i++ {
        var free_cell = area.GetRandFreeCell()
        if free_cell != nil {
            var cell_general cell.Celler
            var move_strategy strategy.MoveStrateger
            x, y := free_cell.GetX(), free_cell.GetY()
            switch c.(type) {
                case clawn.Clawn:
                    cell_general = clawn.New(x, y)
                    ms := random.New()
                    move_strategy = &ms
                case shark.Shark:
                    cell_general = shark.New(x, y)
                    ms := line.New( area.GetWidth() )
                    move_strategy = &ms
            }
            cell_general.Init()
            cell_general.SetStrategy( strategy.New( move_strategy ) )
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