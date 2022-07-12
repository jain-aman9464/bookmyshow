package service

import (
	"github.com/tokopedia/test/bookmyshow/src/model"
	"github.com/tokopedia/test/bookmyshow/src/repo"
)

type BookingService struct {
	bookingRepo repo.BookingRepo
	showRepo    repo.ShowRepo
	theatreRepo repo.TheatreRepo
}

func NewBookingService(bookingRepo repo.BookingRepo, showRepo repo.ShowRepo, theatreRepo repo.TheatreRepo) BookingService {
	return BookingService{
		bookingRepo: bookingRepo,
		showRepo:    showRepo,
		theatreRepo: theatreRepo,
	}
}

func (b BookingService) CreateBooking(userID, showID string, seatIDs []string) string {
	seatsFromTheatre := []model.Seat{}
	show := b.showRepo.GetShow(showID)

	for _, seatID := range seatIDs {
		seatsFromTheatre = append(seatsFromTheatre, b.theatreRepo.GetSeat(seatID))
	}

	return b.bookingRepo.CreateBooking(userID, show, seatsFromTheatre).GetBoookingID()
}
