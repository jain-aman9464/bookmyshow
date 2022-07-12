package repo

import "github.com/tokopedia/test/bookmyshow/src/model"

type TheatreRepo struct {
	theatres map[string]model.Theatre
	screens  map[string]model.Screen
	seats    map[string]model.Seat
}

func NewTheatreRepo() TheatreRepo {
	return TheatreRepo{
		theatres: make(map[string]model.Theatre),
		screens:  make(map[string]model.Screen),
		seats:    make(map[string]model.Seat),
	}
}

func (t TheatreRepo) GetSeat(seatID string) model.Seat {
	if _, ok := t.seats[seatID]; !ok {
		//	return error
	}

	return t.seats[seatID]
}

func (t TheatreRepo) GetTheatre(theatreID string) model.Theatre {
	if _, ok := t.theatres[theatreID]; !ok {
		//	return error
	}

	return t.theatres[theatreID]
}

func (t TheatreRepo) GetScreen(screenID string) model.Screen {
	if _, ok := t.screens[screenID]; !ok {
		//	return error
	}

	return t.screens[screenID]
}

func (t TheatreRepo) CreateTheatre(theatreName string) model.Theatre {
	theatre := model.NewTheatre(theatreName)
	t.theatres[theatre.GetheatreID()] = theatre

	return theatre
}

func (t TheatreRepo) CreateScreen(screenName string, theatre model.Theatre) model.Screen {
	screen := model.NewScreen(screenName, theatre)
	t.screens[screen.GetScreenID()] = screen

	return screen
}

func (t TheatreRepo) CreateSeatInScreen(rowNum int, seatNum int, screen model.Screen) model.Seat {
	seat := model.NewSeat(rowNum, seatNum)
	t.seats[seat.GetSeatID()] = seat
	screen.Addseat(seat)
	return seat
}

func (t TheatreRepo) CreateScreenInTheatre(screenName string, theatre model.Theatre) model.Screen {
	screen := model.NewScreen(screenName, theatre)
	theatre.AddScreen(screen)

	return screen
}
