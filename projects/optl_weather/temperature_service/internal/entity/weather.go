package entity

type Weather struct {
	Current  *Current  `json:"current"`
	Location *Location `json:"location"`
}

type Location struct {
	Name string `json:"name"`
}

type Current struct {
	TempC float32 `json:"temp_c"`
	Tempf float32 `json:"temp_f"`
}

func NewWeather() *Weather {
	return &Weather{}
}
