package repo

import "github.com/tokopedia/test/bookmyshow/src/model"

type SeatLockProvider interface {
	LockSeats(show model.Show, seats []model.Seat, user string)
	UnlockSeats(show model.Show, seats []model.Seat, user string)
	ValidateLock(show model.Show, seat model.Seat, user string) bool
	GetLockedSeats(show model.Show) []model.Seat
}
