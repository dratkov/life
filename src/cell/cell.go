package cell

import (
    "fmt"
    "math/rand"
    "time"
    "../strategy"
)

const (
    StringConsoleCell string = " "
    MaxLiveCycle int = -1
    MaxLiveCycleDeltaPercent int = 0
)

type Celler interface {
    DoAction( c Celler )
    GetX() int
    GetY() int
    SetX( int )
    SetY( int )
    SetID(uint)
    GetID() uint
    String2Console() string
    IsEmptyCell( Celler ) bool
    IncrementLiveCycle()
    GetIsDead() bool
    GetMaxLiveCycle() int
    SetMaxLiveCycle(int, int)
    GetLiveCycle() int
    Init()
    SetStrategy( *strategy.Strategy )
    GetIsDoneAction() bool
    SetIsDoneAction(bool)
}

type Areaer interface {
    GetCell( int, int ) Celler
    GetHeight() uint
    GetWidth() uint
    GetMaxCellID() uint
    SetMaxCellID(uint)
}

type Cell struct {
    x, y int
    live_cycle, max_live_cycle int
    is_dead bool
    console_cell_sign string
    strategy *strategy.Strategy
    is_done_action bool
    id uint
}

func NewCell( x int, y int ) *Cell {
    c := &Cell{ x: x, y: y }
    c.SetIsDead( false )
    c.SetConsoleCellSign( StringConsoleCell )
    max_cell_id := area.GetMaxCellID()
    c.SetID(max_cell_id)
    //area.SetMaxCellID(max_cell_id + 1)
 
    return c
}

func ( c *Cell ) Init() {
    c.SetMaxLiveCycle( MaxLiveCycle, MaxLiveCycleDeltaPercent )
    c.SetConsoleCellSign( StringConsoleCell )
    c.SetIsDoneAction(false)
}

var area Areaer = nil

func ( c *Cell ) GetX() int {
    return c.x
}

func ( c *Cell ) GetY() int {
    return c.y
}

func ( c *Cell ) SetX( x int ) {
    c.x = x
}

func ( c *Cell ) GetID() uint {
    return c.id
}

func ( c *Cell ) SetID( id uint ) {
    c.id = id
}

func SetArea( a Areaer ) {
    area = a
}

func (c *Cell) SetIsDoneAction( is_done bool )  {
    c.is_done_action = is_done
}

func (c *Cell) GetIsDoneAction() bool {
    return c.is_done_action
}

func GetArea() Areaer {
    return area
}

func ( c *Cell ) GetStrategy() *strategy.Strategy {
    return c.strategy
}

func ( c *Cell ) SetStrategy( strategy *strategy.Strategy ) {
    c.strategy = strategy
}

func ( c *Cell ) SetY( y int ) {
    c.y = y
}

func ( c *Cell ) IncrementLiveCycle() {
    if c.GetMaxLiveCycle() < 0 {
        return
    }
    if c.GetLiveCycle() == c.GetMaxLiveCycle() {
        c.SetIsDead( true )
        return
    }
    c.SetLiveCycle( c.GetLiveCycle() + 1 )
}

func ( c *Cell ) SetMaxLiveCycle( max_live_cycle, max_live_cycle_delta_percent int ) {
    if max_live_cycle < 0 {
        c.max_live_cycle = max_live_cycle
        return
    }
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    delta := float32(max_live_cycle_delta_percent) / 100 * float32(max_live_cycle)
    rand_delta := r.Intn( int( delta ) )
    if r.Intn( 2 ) == 0 {
        c.max_live_cycle = max_live_cycle + rand_delta
    } else {
        c.max_live_cycle = max_live_cycle - rand_delta
    }
}

func ( c *Cell ) SetIsDead( is_dead bool ) {
    c.is_dead = is_dead
    c.SetConsoleCellSign( "x" )
}

func ( c *Cell ) SetLiveCycle( next_live_cycle int ) {
    c.live_cycle = next_live_cycle
}

func ( c *Cell ) SetConsoleCellSign( console_cell_sign string ) {
    c.console_cell_sign = console_cell_sign
}

func ( c *Cell ) GetIsDead() bool {
    return c.is_dead
}

func ( c *Cell ) GetConsoleCellSign() string {
    return c.console_cell_sign
}

func ( c *Cell ) GetMaxLiveCycle() int {
    return c.max_live_cycle
}

func ( c *Cell ) GetLiveCycle() int {
    return c.live_cycle
}

func (c *Cell) DoAction( celler Celler ) {
    switch celler.(type) {
        case *Cell:
            return
    }

    celler.IncrementLiveCycle()
}

func ( c *Cell ) IsEmptyCell( celler Celler ) bool {
    switch celler.(type) {
        case *Cell:
            return true
    }
    return false
}

/*
func ( c *Cell ) Move( a Areaer ) {
    
}
*/

func ( c *Cell ) String2Console() string {
    return fmt.Sprintf( c.GetConsoleCellSign() )
}