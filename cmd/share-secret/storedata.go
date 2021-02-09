package main

import (
	"time"

	"github.com/google/uuid"
)

type StoreData struct {
	Secret     string
	Unit       string
	Time       int64
	Registered *time.Time
	Expire     *time.Time
}

func (s *StoreData) GetExpirationDate() *time.Time {
	return nil
}

func (s *StoreData) GetShareData() string {

	return ""
}
func (s *StoreData) GetKey() string {
	return uuid.New().String()
}
