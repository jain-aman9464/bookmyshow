package repo

import "github.com/tokopedia/test/bookmyshow/src/model"

type SeatsAvailabilityRepo struct {
	bookingRepo        BookingRepo
	seatLockerProvider SeatLockProvider
}

func NewSeatsAvailabilityRepo(repo BookingRepo, provider SeatLockProvider) SeatsAvailabilityRepo {
	return SeatsAvailabilityRepo{
		bookingRepo:        repo,
		seatLockerProvider: provider,
	}
}

func (s SeatsAvailabilityRepo) GetAvailableSeats(show model.Show) []model.Seat {
	allSeatsNap := make(map[model.Seat]bool, 0)
	seats := []model.Seat{}

	allSeats := show.GetScreen().GetSeats()
	unavailableSeats := s.getUnavailableSeats(show)

	for _, seat := range allSeats {
		allSeatsNap[seat] = true
	}

	for _, seat := range unavailableSeats {
		allSeatsNap[seat] = false
	}

	for seat, isAvailable := range allSeatsNap {
		if isAvailable {
			seats = append(seats, seat)
		}
	}

	return seats
}

func (s SeatsAvailabilityRepo) getUnavailableSeats(show model.Show) []model.Seat {
	unavailableSeats := s.bookingRepo.getBookedSeats(show)
	unavailableSeats = append(unavailableSeats, s.seatLockerProvider.GetLockedSeats(show)...)

	return unavailableSeats
}
