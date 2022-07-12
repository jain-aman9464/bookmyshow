package service

import "github.com/tokopedia/test/bookmyshow/src/repo"

type MovieService struct {
	movieRepo repo.MovieRepo
}

func NewMovieService(movieRepo repo.MovieRepo) MovieService {
	return MovieService{movieRepo: movieRepo}
}

func (m MovieService) CreateMovie(movieName string) string {
	return m.movieRepo.CreateMovie(movieName).GetMovieID()
}
