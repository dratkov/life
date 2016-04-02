package id

type ID struct {
    id uint
}

func (i *ID) SetID(id uint) {
    i.id = id
}

func (i *ID) GetID() uint {
    return i.id
}

func (i *ID) Increment() *ID {
    return New(i.GetID()+1)
}

func New(id uint) *ID {
    i := &ID{}
    i.SetID(id)

    return i
}
