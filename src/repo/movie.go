package repo

import "github.com/tokopedia/test/bookmyshow/src/model"

type MovieRepo struct {
	movies map[string]model.Movie
}

func NewMovieRepo() MovieRepo {
	return MovieRepo{
		movies: make(map[string]model.Movie),
	}
}

func (m MovieRepo) GetMovie(movieID string) model.Movie {
	if _, ok := m.movies[movieID]; !ok {
		// return error
	}

	return m.movies[movieID]
}

func (m MovieRepo) CreateMovie(movieName string) model.Movie {
	movie := model.NewMovie(movieName)
	m.movies[movie.GetMovieID()] = movie

	return movie
}
