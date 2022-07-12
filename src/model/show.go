package model

import (
	"github.com/google/uuid"
	"time"
)

type Show struct {
	id                string
	movie             Movie
	screen            Screen
	startTime         time.Time
	durationInSeconds int64
}

func NewShow(movie Movie, screen Screen, startTime time.Time, durationInSeconds int64) Show {
	return Show{
		id:                uuid.New().String(),
		movie:             movie,
		screen:            screen,
		startTime:         startTime,
		durationInSeconds: durationInSeconds,
	}
}

func (s Show) GetShowID() string {
	return s.id
}

func (s Show) GetScreen() Screen {
	return s.screen
}
