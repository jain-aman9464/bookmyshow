package model

type BookingStatus int

const (
	Created BookingStatus = iota
	Confirmed
	Expired
)
