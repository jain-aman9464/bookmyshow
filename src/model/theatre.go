package model

import (
	"github.com/google/uuid"
)

type Theatre struct {
	id, name string
	screens  []Screen
}

func NewTheatre(name string) Theatre {
	return Theatre{
		id:      uuid.New().String(),
		name:    name,
		screens: make([]Screen, 0),
	}
}

func (t Theatre) GetheatreID() string {
	return t.id
}

func (t *Theatre) AddScreen(screen Screen) {
	t.screens = append(t.screens, screen)
}
