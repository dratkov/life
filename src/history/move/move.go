package move

import (
    "../../history"
)

type Move struct {
    history.History
    from_x, from_y, to_x, to_y int
}

type Mover interface {
    GetFromX() int
    GetFromY() int
    GetToX() int
    GetToY() int
}

func GetLast(c history.Celler) Mover {
    history := history.GetCellActions(c)
    if history == nil {
        return nil
    }
    idx := len(history) - 1
    for idx >= 0 {
        h := history[idx]
        switch h.(type) {
            case Mover:
                return h.(Mover)
        }
        idx--
    }

    return nil
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

func New(c history.Celler, from_x, from_y, to_x, to_y int) *Move {
    h := history.New(c)
    m := new(Move)
    m.History = *h
    m.SetFromX(from_x)
    m.SetFromY(from_y)
    m.SetToX(to_x)
    m.SetToY(to_y)

    history.AppendAction(m, c)

    return m
}
