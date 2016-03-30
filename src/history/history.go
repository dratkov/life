package history

type History struct {
    cells map[uint][]Historyer
}

type Historyer interface {
    GetFromX() int
    GetFromY() int
    GetToX() int
    GetToY() int
}

type Celler interface {
    GetID() uint
}

func GetLast(c Celler) Historyer {
    hs := New()

    cell_history := hs.cells[c.GetID()]
    if cell_history == nil {
        return nil
    }
    for i, v := range cell_history {
        if v == nil && i > 0 {
            return cell_history[i-1]
        }
    }

    return cell_history[len(cell_history)-1]
}

func Store(c Celler, h Historyer) {
    hs := New()
    cell_history := hs.cells[c.GetID()]
    if cell_history == nil {
        cell_history = make([]Historyer, 10)
    }
    cell_history = append_history(cell_history, h)
    hs.cells[c.GetID()] = cell_history
}

var history_singletone *History = nil

func New() *History {
    if history_singletone != nil {
        return history_singletone
    }
    history_singletone = &History{}
    history_singletone.cells = make(map[uint][]Historyer)

    return history_singletone
}

func append_history(cell_history []Historyer, history Historyer) []Historyer {
    for i, v := range cell_history {
        if v == nil {
            cell_history[i] = history
            return cell_history
        }
    }
    cell_history = cell_history[1:]
    return append(cell_history, history)
}