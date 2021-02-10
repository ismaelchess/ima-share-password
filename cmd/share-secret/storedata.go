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
	return time.Duration(time.Minute + 2)
}

func (s *StoreData) GetKey() string {
	return uuid.New().String()
}
