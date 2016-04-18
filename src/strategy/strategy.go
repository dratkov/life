package strategy

import (
	//"./move"
)

type Celler interface {
	GetX() int
	GetY() int
	SetX(int)
	SetY(int)
	MoveLeft() bool
	MoveUp() bool
	MoveRight() bool
	MoveDawn() bool
	GetID() uint
}

type Strategy struct {
	move_strateger Strategyer
}

type MoveStrateger interface {
	Move(c Celler)
}

type Strategyer interface {
	Move( Celler )
}

func (s *Strategy) GetMoveStrateger() Strategyer {
	return s.move_strateger
}

func (s *Strategy) SetMoveStrateger( m Strategyer ) {
	s.move_strateger = m
}

func (s *Strategy) Move( c Celler ) {
	move_strateger := s.GetMoveStrateger()
	move_strateger.Move(c)
}

func NewStrategy( move_strategy MoveStrateger ) *Strategy {
	strategy := &Strategy{}
	strategy.SetMoveStrateger( move_strategy )

	return strategy
}