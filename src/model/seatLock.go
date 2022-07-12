package model

import "time"

type SeatLock struct {
	seat             Seat
	show             Show
	timeOutInSeconds int64
	lockTime         time.Time
	lockedBy         string
}

func NewSeatLock(seat Seat, show Show, timeoutinsec int64, lockTime time.Time, lockedBy string) SeatLock {
	return SeatLock{
		seat:             seat,
		show:             show,
		timeOutInSeconds: timeoutinsec,
		lockTime:         lockTime,
		lockedBy:         lockedBy,
	}
}

func (s SeatLock) IsLockExpired() bool {
	lockInstant := s.lockTime.Add(time.Duration(s.timeOutInSeconds))
	currTime := time.Now()

	return lockInstant.Before(currTime)
}

func (s SeatLock) GetLockedBy() string {
	return s.lockedBy
}
