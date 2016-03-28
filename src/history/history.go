package history

type History struct {

}

var history_singletone *History = nil

func New() *History {
    if history_singletone != nil {
        return history_singletone
    }
    history_singletone = &History{}

    return history_singletone
}