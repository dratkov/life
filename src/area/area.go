package area

import (
	"../cell"
	"math/rand"
	"time"
	//"fmt"
)

type Area struct {
	width, height uint
	celler []cell.Celler
	live_cell_count uint
}

type Areaer interface {
    GetCell( int, int ) cell.Celler
    GetHeight() uint
    GetWidth() uint
}

func New( width, height uint ) Area {
	new_area := Area{ width: width, height: height }
	new_area.celler = make( []cell.Celler, height * width )

	return new_area
}

func ( a *Area ) AddCell( c cell.Celler ) {
	idx := c.GetY() * int(a.GetWidth()) + c.GetX()
	a.celler[ idx ] = c
	if c.IsEmptyCell(c) == false {
		a.SetLiveCellCount( a.GetLiveCellCount() + 1 )
	}
}

func ( a *Area ) GetLiveCellCount() uint {
	return a.live_cell_count
}

func ( a *Area ) SetLiveCellCount( live_cell_count uint ) {
	a.live_cell_count = live_cell_count
}

func ( a *Area ) GetWidth() uint {
	return a.width
}

func ( a *Area ) GetHeight() uint {
	return a.height
}

func ( a *Area ) GetCell( x int, y int ) cell.Celler {
	for idx := range a.celler {
		c := a.celler[ idx ]
		if c.GetX() == x && c.GetY() == y {
			return c
		}
	}

	return nil
}

func ( a *Area ) DeleteCell( to_delete_cell cell.Celler ) {
	x, y := to_delete_cell.GetX(), to_delete_cell.GetY()
	for idx := range a.celler {
		c := a.celler[ idx ]
		if c.GetX() == x && c.GetY() == y {
			new_cell := cell.New( c.GetX(), c.GetY() )
			a.celler[ idx ] = &new_cell
			a.SetLiveCellCount( a.GetLiveCellCount() - 1 )
		}
	}
}

func ( a *Area ) GetRandFreeCell() cell.Celler {
	free_cells := make( []cell.Celler, 0 )
	for i := 0; i < int(a.GetHeight()); i++ {
		for j := 0; j < int(a.GetWidth()); j++ {
			c := a.GetCell( j, i )
			if c.IsEmptyCell( c ) {
				free_cells = append( free_cells, c )
			}
		}
	}
	free_count := len( free_cells )
	if free_count > 0 {
		r := rand.New( rand.NewSource(time.Now().UnixNano()) )
		c := free_cells[ r.Intn( free_count ) ]
		return c
	}

	return nil
}

func ( a *Area ) NextTime() {
	cells := [] cell.Celler{}
	for i := uint(0); i < a.GetHeight(); i++ {
		for j := uint(0); j < a.GetWidth(); j++ {
			cell := a.GetCell( int(j), int(i) )
			if cell.GetIsDead() {
				a.DeleteCell( cell )
				continue
			}
			if cell.GetIsDoneAction() == false {
				cell.DoAction(a)
				cell.SetIsDoneAction(true)
				cells = append(cells, cell)
			}
		}
	}

	for _, c := range cells {
		c.SetIsDoneAction(false)
	}
}

