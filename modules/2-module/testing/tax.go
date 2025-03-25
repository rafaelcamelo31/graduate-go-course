package main

import (
	"errors"
	"time"
)

type Repository interface {
	SaveTax(tax float64) error
}

/*
Simulate database operation using a repository and mock it in the test
*/
func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.SaveTax(tax)
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amount should be greater than 0")
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0, nil
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}

func CalculateTax2(amount float64) float64 {
	// Simulate a delay
	time.Sleep(time.Millisecond)
	if amount == 0 {
		return 0
	}
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
