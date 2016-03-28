package random

import (
//"../../../area"
	"../../../strategy"
	"math/rand"
	"time"
)

type Random struct {

}

type MoveDirectFunc func() bool

func (r *Random) Move(c strategy.Celler) {
	directs := []MoveDirectFunc{c.MoveLeft, c.MoveUp, c.MoveRight, c.MoveDawn }
	shuffle(directs)
	switch {
		case directs[ 0 ]() == true:

		case directs[ 1 ]() == true:

		case directs[ 2 ]() == true:

		case directs[ 3 ]() == true:

	}
}

func shuffle(d []MoveDirectFunc) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		j := r.Intn(len(d))
		d[i], d[j] = d[j], d[i]
	}
}

func New() Random {
	return Random{}
}