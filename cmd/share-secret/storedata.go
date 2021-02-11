package main

import (
	"time"

	"github.com/google/uuid"
)

type StoreData struct {
	Secret string
	Unit   string
	Time   int
}

func (s *StoreData) expirationDate() time.Duration {
	switch s.Unit {
	case "h":
		return time.Duration(s.Time) * time.Hour
	case "d":
		return time.Duration(s.Time*24) * time.Hour
	}
	return time.Duration(s.Time) * time.Minute
}

func (s *StoreData) GetKey() string {
	return uuid.New().String()
}
