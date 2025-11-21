package entity

import "unicode"

type ViaCEP struct {
	Cep  string `json:"cep"`
	City string `json:"localidade"`
}

func NewViaCEP(cep string) *ViaCEP {
	return &ViaCEP{
		Cep: cep,
	}
}

func (v *ViaCEP) IsAllDigits() bool {
	for _, s := range v.Cep {
		if !unicode.IsDigit(s) {
			return false
		}
	}
	return true
}

func (v *ViaCEP) IsEightDigits() bool {
	cepLength := len(v.Cep)
	return cepLength != 8
}
