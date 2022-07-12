package repo

import (
	"github.com/tokopedia/test/bookmyshow/src/model"
)

type BookingRepo struct {
	showBookings     map[string]model.Booking
	seatLockProvider SeatLockProvider
}

func NewBookingRepo(lock SeatLockProvider) BookingRepo {
	return BookingRepo{
		showBookings:     make(map[string]model.Booking, 0),
		seatLockProvider: lock,
	}
}

func (b BookingRepo) GetBooking(bookingID string) model.Booking {
	if _, ok := b.showBookings[bookingID]; !ok {
		//	return err
	}

	return b.showBookings[bookingID]
}

func (b BookingRepo) GetAllBookings(show model.Show) []model.Booking {
	bookings := make([]model.Booking, 0)
	for _, booking := range b.showBookings {
		if booking.GetShow().GetShowID() == show.GetShowID() {
			bookings = append(bookings, booking)
		}
	}

	return bookings
}

func (b BookingRepo) isAnySeatAlreadyBooked(show model.Show, seats []model.Seat) bool {
	bookedSeats := b.getBookedSeats(show)
	bookedSeatMap := make(map[string]bool, 0)

	for _, bookedSeat := range bookedSeats {
		bookedSeatMap[bookedSeat.GetSeatID()] = true
	}

	for _, seat := range seats {
		if isBooked, ok := bookedSeatMap[seat.GetSeatID()]; ok && isBooked {
			return true
		}
	}

	return false
}

func (b BookingRepo) getBookedSeats(show model.Show) []model.Seat {
	seats := make([]model.Seat, 0)
	bookings := b.GetAllBookings(show)

	for _, booking := range bookings {
		if booking.IsConfirmed() {
			seats = append(seats, booking.GetShow().GetScreen().GetSeats()...)
		}
	}

	return seats
}

func (b BookingRepo) CreateBooking(userID string, show model.Show, seats []model.Seat) model.Booking {
	if b.isAnySeatAlreadyBooked(show, seats) {
		//	return err
	}

	b.seatLockProvider.LockSeats(show, seats, userID)
	newBooking := model.NewBooking(show, userID, seats)
	b.showBookings[newBooking.GetBoookingID()] = newBooking
	// TODO: Create timer for booking expiry
	return newBooking
}

func (b BookingRepo) ConfirmBooking(booking model.Booking, user string) {
	if booking.GetUser() != user {
		//	return err
	}

	for _, seat := range booking.GetSeatsBooked() {
		if !b.seatLockProvider.ValidateLock(booking.GetShow(), seat, user) {
			//	return err
		}
	}

	booking.ConfirmBooking()
}
