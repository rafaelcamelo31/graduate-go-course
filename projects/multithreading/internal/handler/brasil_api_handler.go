package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	BRASIL_API_URL = "https://brasilapi.com.br/api/cep/v1/01153000"
)

type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
	URL          string `json:"url"`
}

func NewBrasilAPI() *BrasilAPI {
	return &BrasilAPI{
		URL: BRASIL_API_URL,
	}
}

func (b *BrasilAPI) ToJSON() {
	json, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(json))
}

func (b *BrasilAPI) GetBrasilAPIAddress(w http.ResponseWriter, ch chan *BrasilAPI) error {
	req, err := http.NewRequest(http.MethodGet, BRASIL_API_URL, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(body, b)
	if err != nil {
		log.Println(err)
		return err
	}

	ch <- b

	return nil
}
