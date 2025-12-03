package entity

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CityTestSuite struct {
	suite.Suite
}

func (suite *CityTestSuite) TestIsAllDigits_VariousCEPs() {
	testCases := []struct {
		name     string
		cep      string
		expected bool
	}{
		{"valid numeric", "13083970", true},
		{"with letter", "1308397a", false},
		{"with dash", "13083-970", false},
		{"with space", "13083 970", false},
		{"all letters", "abcdefgh", false},
		{"special chars", "!@#$%^&*", false},
		{"mixed", "1234ab56", false},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			city := NewCity(tc.cep)
			result := city.IsAllDigits()
			suite.Equal(tc.expected, result, "CEP: %s", tc.cep)
		})
	}
}

func (suite *CityTestSuite) TestIsEightDigits_VariousLengths() {
	testCases := []struct {
		name     string
		cep      string
		expected bool
	}{
		{"6 digits", "123456", false},
		{"7 digits", "1234567", false},
		{"8 digits", "12345678", true},
		{"9 digits", "123456789", false},
		{"10 digits", "1234567890", false},
		{"empty", "", false},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			city := NewCity(tc.cep)
			result := city.IsEightDigits()
			suite.Equal(tc.expected, result, "CEP length: %d", len(tc.cep))
		})
	}
}

func TestCityTestSuite(t *testing.T) {
	suite.Run(t, new(CityTestSuite))
}
