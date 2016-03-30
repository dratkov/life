package move

import (
    "time"
    "../../history"
)

type Move struct {
    from_x, from_y, to_x, to_y int
    t time.Time
}

func (m *Move) SetFromX(from_x int) {
    m.from_x = from_x
}

func (m *Move) GetFromX() int {
    return m.from_x
}

func (m *Move) SetFromY(from_y int) {
    m.from_y = from_y
}

func (m *Move) GetFromY() int {
    return m.from_y
}

func (m *Move) SetToX(to_x int) {
    m.to_x = to_x
}

func (m *Move) GetToX() int {
    return m.to_x
}

func (m *Move) SetToY(to_y int) {
    m.to_y = to_y
}

func (m *Move) GetToY() int {
    return m.to_y
}

func (m *Move) SetTime(t time.Time) {
    m.t = t
}

func (m *Move) GetTime() time.Time {
    return m.t
}

func (m *Move) Store(c history.Celler) {
    history.Store(c, m)
}

func GetLast(c history.Celler) history.Historyer {
    return history.GetLast(c)
}

func New(from_x, from_y, to_x, to_y int) *Move {
    m := &Move{}
    m.SetFromX(from_x)
    m.SetFromY(from_y)
    m.SetToX(to_x)
    m.SetToY(to_y)
    m.SetTime(time.Now())

    return m
}
