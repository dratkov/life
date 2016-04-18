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
    BuildCell( interface{} ) cell.Celler
    BuildCells( interface{}, int )
}

func NewBuilder() *Build {
    return &Build{}
}

func ( b *Build ) BuildArea( width, height uint ) {
    b.area = area.NewArea( width, height )
    cell.SetArea(b.area)
    b.FillArea()
}

func ( b *Build ) GetArea() *area.Area {
    return b.area
    a := b.area
    return a
}

func ( b *Build ) BuildCell( c interface{} ) cell.Celler {
    area := b.GetArea()
    var free_cell = area.GetRandFreeCell()
    if free_cell != nil {
        var cell_general cell.Celler
        var move_strategy strategy.MoveStrateger
        x, y := free_cell.GetX(), free_cell.GetY()
        switch c.(type) {
        case clawn.Clawn:
            cell_general = clawn.NewClawn(x, y)
            ms := random.New()
            move_strategy = &ms
        case shark.Shark:
            cell_general = shark.NewShark(x, y)
            ms := line.New( area.GetWidth() )
            move_strategy = &ms
        }
        cell_general.Init()
        cell_general.SetStrategy( strategy.NewStrategy( move_strategy ) )
        area.AddCell( cell_general )

        return cell_general
    }

    return nil
}

func ( b *Build ) BuildCells( c interface{}, count int ) {
    for i := 0; i < count; i++ {
       b.BuildCell(c)
    }
}

func ( b Build ) FillArea() {
    area := b.GetArea()
    height, width := area.GetHeight(), area.GetWidth()
    for i := 0; i < int(height); i++ {
        for j := 0; j < int(width); j++ {
            cell_el := cell.NewCell( j, i )
            area.AddCell( cell_el )
        }
    }
}