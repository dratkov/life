package cell

import (
    "fmt"
    "math/rand"
    "time"
    "../strategy"
)

const (
    StringConsoleCell string = " "
    MaxLiveCicle int = -1
    MaxLiveCicleDeltaPercent int = 0
)

type Celler interface {
    DoAction( a Areaer )
    GetX() int
    GetY() int
    SetX( int )
    SetY( int )
    SetID(uint)
    GetID() uint
    String2Console() string
    New( int, int ) Celler
    IsEmptyCell( Celler ) bool
    IncrementLiveCicle()
    GetIsDead() bool
    GetMaxLiveCicle() int
    GetLiveCicle() int
    Init()
    SetStrategy( strategy.Strategy )
    SetArea(Areaer)
    GetArea() Areaer
    GetIsDoneAction() bool
    SetIsDoneAction(bool)
}

type Areaer interface {
    GetCell( int, int ) Celler
    GetHeight() uint
    GetWidth() uint
}

type Cell struct {
    x, y int
    live_cicle, max_live_cicle int
    is_dead bool
    console_cell_sign string
    strategy strategy.Strategy
    area Areaer
    is_done_action bool
    id uint
}

func New( x int, y int ) Cell {
    c := Cell{ x: x, y: y }
    c.SetIsDead( false )
    c.SetConsoleCellSign( StringConsoleCell )
 
    return c
}

func ( c *Cell ) New( x int, y int ) Celler {
    c.SetX( x )
    c.SetY( y )

    return c
}

func ( c *Cell ) Init() {
    c.SetMaxLiveCicle( MaxLiveCicle, MaxLiveCicleDeltaPercent )
    c.SetConsoleCellSign( StringConsoleCell )
    c.SetIsDoneAction(false)
}

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

func ( c *Cell ) SetArea( a Areaer ) {
    c.area = a
}

func (c *Cell) SetIsDoneAction( is_done bool )  {
    c.is_done_action = is_done
}

func (c *Cell) GetIsDoneAction() bool {
    return c.is_done_action
}

func (c *Cell) GetArea() Areaer {
    return c.area
}

func ( c *Cell ) GetStrategy() strategy.Strategy {
    return c.strategy
}

func ( c *Cell ) SetStrategy( strategy strategy.Strategy ) {
    c.strategy = strategy
}

func ( c *Cell ) SetY( y int ) {
    c.y = y
}

func ( c *Cell ) IncrementLiveCicle() {
    if c.GetLiveCicle() == c.GetMaxLiveCicle() {
        c.SetIsDead( true )
        return
    }
    c.SetLiveCicle( c.GetLiveCicle() + 1 )
}

func ( c *Cell ) SetMaxLiveCicle( max_live_cicle, max_live_cicle_delta_percent int ) {
    if max_live_cicle < 0 {
        c.max_live_cicle = max_live_cicle
        return
    }
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    delta := float32(max_live_cicle_delta_percent) / 100 * float32(max_live_cicle)
    rand_delta := r.Intn( int( delta ) )
    if r.Intn( 2 ) == 0 {
        c.max_live_cicle = max_live_cicle + rand_delta
    } else {
        c.max_live_cicle = max_live_cicle - rand_delta
    }
}

func ( c *Cell ) SetIsDead( is_dead bool ) {
    c.is_dead = is_dead
    c.SetConsoleCellSign( "x" )
}

func ( c *Cell ) SetLiveCicle( next_live_cicle int ) {
    c.live_cicle = next_live_cicle
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

func ( c *Cell ) GetMaxLiveCicle() int {
    return c.max_live_cicle
}

func ( c *Cell ) GetLiveCicle() int {
    return c.live_cicle
}

func ( c *Cell ) DoAction( a Areaer ) {
    return
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