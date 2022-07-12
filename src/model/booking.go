package model

import (
	"github.com/google/uuid"
)

type Booking struct {
	id            string
	show          Show
	seatsBooked   []Seat
	user          string
	bookingStatus BookingStatus
}

func NewBooking(show Show, user string, seatsBooked []Seat) Booking {
	return Booking{
		id:            uuid.New().String(),
		show:          show,
		seatsBooked:   seatsBooked,
		user:          user,
		bookingStatus: Created,
	}
}

func (b Booking) IsConfirmed() bool {
	return b.bookingStatus == Confirmed
}

func (b Booking) ConfirmBooking() {
	if b.bookingStatus != Created {
		//	 return error
	}

	b.bookingStatus = Confirmed
}

func (b Booking) ExpireBooking() {
	if b.bookingStatus != Created {
		//	 return error
	}

	b.bookingStatus = Expired
}

func (b Booking) GetShow() Show {
	return b.show
}

func (b Booking) GetBoookingID() string {
	return b.id
}

func (b Booking) GetUser() string {
	return b.user
}

func (b Booking) GetSeatsBooked() []Seat {
	return b.seatsBooked
}
