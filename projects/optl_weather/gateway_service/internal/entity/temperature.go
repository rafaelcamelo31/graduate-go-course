package entity

type Temperature struct {
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
	TempK float32 `json:"temp_k"`
	City  string  `json:"city"`
}

func NewTemperature(tempC, tempF float32) *Temperature {
	return &Temperature{
		TempC: tempC,
		TempF: tempF,
		TempK: tempC + 273,
	}
}
