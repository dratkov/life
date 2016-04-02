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

type History struct {
    cells map[uint][]Mover
}

var actions map[uint][]Mover
var max_cell_actions int = 10

func GetLast(c history.Celler) Mover {
    cell_history := actions[c.GetID()]
    if cell_history == nil || len(cell_history) == 0 {
        return nil
    }

    return cell_history[len(cell_history)-1]
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

func Undo(c history.Celler) {

}

func New(c history.Celler, from_x, from_y, to_x, to_y int) *Move {
    if actions == nil {
        actions = make(map[uint][]Mover)
    }
    h := history.New(c)
    m := new(Move)
    m.History = *h
    m.SetFromX(from_x)
    m.SetFromY(from_y)
    m.SetToX(to_x)
    m.SetToY(to_y)

    cell_actions := actions[c.GetID()]
    if cell_actions != nil && len(cell_actions) == max_cell_actions {
        cell_actions = cell_actions[1:]
    }
    cell_actions = append(cell_actions, m)
    actions[c.GetID()] = cell_actions

    return m
}
