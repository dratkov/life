package history

import (
    "./id"
    "time"
)

type History struct {
    id *id.ID
    t time.Time
}

var actions map[uint][]Historyer
var max_cell_actions int = 10

type Historyer interface {
    GetID() *id.ID
    SetID(*id.ID)
    GetTime() time.Time
    SetTime(time.Time)
}

type Celler interface {
    GetID() uint
}

func NewID(c Celler) *id.ID {
    last_history := GetLast(c)
    if last_history == nil {
        return id.New(1)
    }
    return last_history.GetID().Increment()
}

func GetLast(c Celler) Historyer {
    cell_history := actions[c.GetID()]
    if cell_history == nil || len(cell_history) == 0 {
        return nil
    }

    return cell_history[len(cell_history)-1]
}

func (h *History) SetTime(t time.Time) {
    h.t = t
}

func (h *History) GetTime() time.Time {
    return h.t
}

func (h *History) SetID(id *id.ID) {
    h.id = id
}

func (h *History) GetID() *id.ID {
    return h.id
}

func New(c Celler) *History {
    if actions == nil {
        actions = make(map[uint][]Historyer)
    }
    h := &History{}
    h.SetID(NewID(c))
    h.SetTime(time.Now())

    cell_actions := actions[c.GetID()]
    if cell_actions != nil && len(cell_actions) == max_cell_actions {
        cell_actions = cell_actions[1:]
    }
    cell_actions = append(cell_actions, h)
    actions[c.GetID()] = cell_actions

    return h
}
