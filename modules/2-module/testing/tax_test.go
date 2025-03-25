package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
   Coverage:
    go test -coverprofile=coverage.out

   Show coverage in html:
    go tool cover -html=coverage.out
*/

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result, _ := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type taxTest struct {
		amount, expected float64
	}

	tests := []taxTest{
		{500, 5.0},
		{1500, 10.0},
		{25000, 20.0},
		{0, 0},
	}

	for _, v := range tests {
		result, _ := CalculateTax(v.amount)
		if result != v.expected {
			t.Errorf("Expected %f but got %f", v.expected, result)
		}
	}
}

/*
   Benchmark:
   go test -bench=. -run=^#
   go test -bench=. -count 5 -run=^# -benchtime=5s -benchmem
*/

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)
	}
}

/*
	Fuzzing:
	go test -fuzz=. -fuzztime=5s -run=^#
*/

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0, 25000.0, 25001.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result, _ := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})
}

/*
	Testify assertions
*/

func TestCalculateTaxTestify(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.NotNil(t, err)
	assert.Error(t, err, "amount should be greater than 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")
}

/*
	Testify assertions with mock
*/

func TestCalculateTaxAndSave(t *testing.T) {
	repository := new(TaxRepositoryMock)
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("amount should be greater than 0"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.NotNil(t, err)

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
