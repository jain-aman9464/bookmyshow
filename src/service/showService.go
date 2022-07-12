package service

import (
	"github.com/tokopedia/test/bookmyshow/src/repo"
	"time"
)

type ShowService struct {
	seatAvailabilityRepo repo.SeatsAvailabilityRepo
	theatreRepo          repo.TheatreRepo
	movieRepo            repo.MovieRepo
	showRepo             repo.ShowRepo
}

func NewShowService(seatAvailabilityRepo repo.SeatsAvailabilityRepo, theatreRepo repo.TheatreRepo, movieRepo repo.MovieRepo, showRepo repo.ShowRepo) ShowService {
	return ShowService{
		seatAvailabilityRepo: seatAvailabilityRepo,
		theatreRepo:          theatreRepo,
		movieRepo:            movieRepo,
		showRepo:             showRepo,
	}
}

func (s ShowService) CreateShow(movieID, screenID string, startTime time.Time, durationInSec int64) string {
	screen := s.theatreRepo.GetScreen(screenID)
	movie := s.movieRepo.GetMovie(movieID)

	return s.showRepo.CreateShow(movie, screen, startTime, durationInSec).GetShowID()
}

func (s ShowService) GetAvailableSeats(showID string) []string {
	var seats []string
	show := s.showRepo.GetShow(showID)
	availableSeats := s.seatAvailabilityRepo.GetAvailableSeats(show)
	for _, seat := range availableSeats {
		seats = append(seats, seat.GetSeatID())
	}

	return seats

}
