package repo

import (
	"github.com/tokopedia/test/bookmyshow/src/model"
	"time"
)

type ShowRepo struct {
	shows map[string]model.Show
}

func NewShowRepo() ShowRepo {
	return ShowRepo{shows: make(map[string]model.Show, 0)}
}

func (s ShowRepo) GetShow(showID string) model.Show {
	if _, ok := s.shows[showID]; !ok {
		// return error
	}

	return s.shows[showID]
}

func (s ShowRepo) GetShowsForScreen(screen model.Screen) []model.Show {
	shows := make([]model.Show, 0)
	for _, show := range s.shows {
		if show.GetScreen().GetScreenID() == screen.GetScreenID() {
			shows = append(shows, show)
		}
	}

	return shows
}

func (s ShowRepo) CreateShow(movie model.Movie, screen model.Screen, startTime time.Time, durationInSeconds int64) model.Show {
	if !s.checkIfShowCreationAllowed(screen, startTime, durationInSeconds) {
		//	return err
	}

	show := model.NewShow(movie, screen, startTime, durationInSeconds)
	s.shows[show.GetShowID()] = show

	return show
}

func (s ShowRepo) checkIfShowCreationAllowed(screen model.Screen, startTime time.Time, durationInSeconds int64) bool {
	return true
}


