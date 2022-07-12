package service

import "github.com/tokopedia/test/bookmyshow/src/repo"

type TheatreService struct {
	theatreRepo repo.TheatreRepo
}

func NewTheatreService(theatreRepo repo.TheatreRepo) TheatreService {
	return TheatreService{theatreRepo: theatreRepo}
}

func (t TheatreService) CreateTheatre(theatreName string) string {
	return t.theatreRepo.CreateTheatre(theatreName).GetheatreID()
}

func (t TheatreService) CreateScreenInTheatre(screenName, theatreID string) string {
	theatre := t.theatreRepo.GetTheatre(theatreID)
	return t.theatreRepo.CreateScreenInTheatre(screenName, theatre).GetScreenID()
}

func (t TheatreService) CreatSeatInScreen(rowNum, seatNum int, screenID string) string {
	screen := t.theatreRepo.GetScreen(screenID)
	return t.theatreRepo.CreateSeatInScreen(rowNum, seatNum, screen).GetSeatID()

}
