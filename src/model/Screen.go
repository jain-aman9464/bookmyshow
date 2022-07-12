package model

import (
	"github.com/google/uuid"
)

type Screen struct {
	name, id string
	theatre  Theatre
	seats    []Seat
}

func NewScreen(name string, theatre Theatre) Screen {
	return Screen{
		name:    name,
		id:      uuid.New().String(),
		theatre: theatre,
		seats:   make([]Seat, 0),
	}
}

func (s Screen) GetScreenID() string {
	return s.id
}

func (s *Screen) Addseat(seat Seat) {
	s.seats = append(s.seats, seat)
}

func (s Screen) GetSeats() []Seat {
	return s.seats
}
