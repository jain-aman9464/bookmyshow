package repo

import "github.com/tokopedia/test/bookmyshow/src/model"

type PaymentRepo struct {
	bookingFailures  map[string]int
	allowedRetries   int
	seatLockProvider SeatLockProvider
}

func NewPaymentRepo(allowedRetries int, lock SeatLockProvider) PaymentRepo {
	return PaymentRepo{
		allowedRetries:   allowedRetries,
		seatLockProvider: lock,
		bookingFailures:  make(map[string]int, 0),
	}
}

func (p PaymentRepo) ProcessPaymentFailed(booking model.Booking, user string) {
	if booking.GetUser() != user {
		//	return err
	}

	if _, ok := p.bookingFailures[booking.GetBoookingID()]; !ok {
		p.bookingFailures[booking.GetBoookingID()] = 0
	}

	currentFailureCount := p.bookingFailures[booking.GetBoookingID()]
	newFailureCount := currentFailureCount + 1

	if newFailureCount > p.allowedRetries {
		p.seatLockProvider.UnlockSeats(booking.GetShow(), booking.GetSeatsBooked(), booking.GetUser())
	}

}
