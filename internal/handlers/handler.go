package handlers

import (
	"ctogram/internal/store"
)

type Cities struct {
	Store *store.Store
}

func NewCities(s *store.Store) *Cities {
	return &Cities{
		Store: s,
	}
}
