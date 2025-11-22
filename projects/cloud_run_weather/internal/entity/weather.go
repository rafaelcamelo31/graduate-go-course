package entity

type Weather struct {
	Current Current `json:"current"`
}

type Current struct {
	TempC float32 `json:"temp_c"`
	Tempf float32 `json:"temp_f"`
}

func NewWeather() *Weather {
	return &Weather{}
}
