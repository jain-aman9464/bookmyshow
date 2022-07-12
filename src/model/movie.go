package model

type Movie struct {
	id   string
	name string
}

func NewMovie(name string) Movie {
	return Movie{
		name: name,
	}
}

func (m Movie) GetMovieID() string {
	return m.id
}
