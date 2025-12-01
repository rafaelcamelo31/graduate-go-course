package entity

import "unicode"

type City struct {
	Cep  string
	Name string
}

func NewCity(cep string) *City {
	return &City{
		Cep: cep,
	}
}

func (v *City) IsAllDigits() bool {
	for _, s := range v.Cep {
		if !unicode.IsDigit(s) {
			return false
		}
	}
	return true
}

func (v *City) IsEightDigits() bool {
	cepLength := len(v.Cep)
	return cepLength == 8
}
