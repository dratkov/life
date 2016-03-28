package move

import (
	"../../strategy"
	"./random"
)

type MoveStrateg struct {
	result_strateg strategy.MoveStrateger
}

/*
type MoveStrateger interface {
	Move(c strategy.Celler)
}
*/

func (ms *MoveStrateg) Move(c strategy.Celler) {
	result_strateg := ms.GetResultStrateger()
	result_strateg.Move(c)
}

func (ms *MoveStrateg) SetResultStrateger( s strategy.MoveStrateger ) {
	ms.result_strateg = s
}

func (ms *MoveStrateg) GetResultStrateger() strategy.MoveStrateger {
	return ms.result_strateg
}

func New() MoveStrateg {
	ms := MoveStrateg{}
	result_strateg := random.New()
	ms.SetResultStrateger( &result_strateg )

	return ms
}
