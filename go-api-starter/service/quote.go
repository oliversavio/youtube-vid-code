package service

import (
	"math/rand"
	"time"
)

type Quotter interface {
	GetQuote() (string, error)
}

func GetQuote(quotter Quotter) (string, error) {
	return quotter.GetQuote()
}

type SimpleQuote struct {
}

func (s *SimpleQuote) GetQuote() (string, error) {

	n := rand.Intn(500) // n will be between 0 and 10
	time.Sleep(time.Duration(n) * time.Millisecond)

	return "Hello World", nil
}
