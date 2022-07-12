package repo

import (
	"github.com/tokopedia/test/bookmyshow/src/model"
	"time"
)

type InMemorySeatLockProvider struct {
	locks       map[string]map[model.Seat]model.SeatLock
	lockTimeout int64
}

func NewInMemorySeatLockProvider(timeout int64) InMemorySeatLockProvider {
	return InMemorySeatLockProvider{
		locks:       make(map[string]map[model.Seat]model.SeatLock, 0),
		lockTimeout: timeout,
	}
}

func (i InMemorySeatLockProvider) LockSeats(show model.Show, seats []model.Seat, user string) {
	for _, seat := range seats {
		if i.isSeatLocked(show, seat) {
			//	return err
		}
	}

	for _, seat := range seats {
		i.lockSeat(show, seat, user, i.lockTimeout)
	}
}

func (i InMemorySeatLockProvider) UnlockSeats(show model.Show, seats []model.Seat, user string) {
	for _, seat := range seats {
		if i.ValidateLock(show, seat, user) {
			i.unlockSeat(show, seat)
		}
	}
}

func (i InMemorySeatLockProvider) ValidateLock(show model.Show, seat model.Seat, user string) bool {
	if i.isSeatLocked(show, seat) && i.locks[show.GetShowID()][seat].GetLockedBy() == user {
		return true
	}

	return false
}

func (i InMemorySeatLockProvider) GetLockedSeats(show model.Show) []model.Seat {
	lockedSeats := []model.Seat{}

	if _, ok := i.locks[show.GetShowID()]; !ok {
		//	return empty
	}

	for seat, _ := range i.locks[show.GetShowID()] {
		if i.isSeatLocked(show, seat) {
			lockedSeats = append(lockedSeats, seat)
		}
	}

	return lockedSeats
}

func (i InMemorySeatLockProvider) isSeatLocked(show model.Show, seat model.Seat) bool {
	if _, ok := i.locks[show.GetShowID()]; !ok {
		return false
	}

	if _, ok := i.locks[show.GetShowID()][seat]; !ok {
		return false
	}

	if i.locks[show.GetShowID()][seat].IsLockExpired() {
		return false
	}

	return true
}

func (i InMemorySeatLockProvider) lockSeat(show model.Show, seat model.Seat, user string, timeoutInSec int64) {
	if _, ok := i.locks[show.GetShowID()]; !ok {
		i.locks[show.GetShowID()] = make(map[model.Seat]model.SeatLock, 0)
	}

	lock := model.NewSeatLock(seat, show, timeoutInSec, time.Now(), user)

	mp := map[model.Seat]model.SeatLock{
		seat: lock,
	}

	i.locks[show.GetShowID()] = mp
}

func (i InMemorySeatLockProvider) unlockSeat(show model.Show, seat model.Seat) {
	if _, ok := i.locks[show.GetShowID()]; !ok {
		return
	}

	delete(i.locks[show.GetShowID()], seat)
}
