package model

import (
	"github.com/google/uuid"
)

type Seat struct {
	id      string
	rowNum  int
	seatNum int
}

func NewSeat(rowNum, seatNum int) Seat {
	return Seat{
		rowNum:  rowNum,
		id:      uuid.New().String(),
		seatNum: seatNum,
	}
}

func (s Seat) GetSeatID() string {
	return s.id
}
