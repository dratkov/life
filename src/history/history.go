package history

type History struct {
}

var actions map[uint][]Historyer
var max_cell_actions int = 50

type Historyer interface {
}

type Celler interface {
    GetID() uint
}

func GetCellActions(c Celler) []Historyer {
    return actions[c.GetID()]
}

func DeleteActions(c Celler) {
    delete(actions, c.GetID())
}

func AppendAction(h Historyer, c Celler) {
    cell_actions := actions[c.GetID()]
    if cell_actions != nil && len(cell_actions) == max_cell_actions {
        cell_actions = cell_actions[1:]
    }
    cell_actions = append(cell_actions, h)
    actions[c.GetID()] = cell_actions
}

func New(c Celler) *History {
    if actions == nil {
        actions = make(map[uint][]Historyer)
    }
    h := &History{}
/*
    cell_actions := actions[c.GetID()]
    if cell_actions != nil && len(cell_actions) == max_cell_actions {
        cell_actions = cell_actions[1:]
    }
    cell_actions = append(cell_actions, h)
    actions[c.GetID()] = cell_actions
*/

    return h
}
