package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	VIA_CEP_API_URL = "https://viacep.com.br/ws/01153000/json/"
)

type ViaCEPAPI struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	URL         string `json:"url"`
}

func NewViaCEPAPI() *ViaCEPAPI {
	return &ViaCEPAPI{
		URL: VIA_CEP_API_URL,
	}
}

func (v *ViaCEPAPI) ToJSON() {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(json))
}

func (v *ViaCEPAPI) GetViaCEPAPIAddress(w http.ResponseWriter, ch chan *ViaCEPAPI) error {
	req, err := http.NewRequest(http.MethodGet, VIA_CEP_API_URL, nil)
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

	err = json.Unmarshal(body, v)
	if err != nil {
		log.Println(err)
		return err
	}

	ch <- v

	return nil
}
