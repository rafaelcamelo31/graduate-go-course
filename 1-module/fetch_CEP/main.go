package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
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
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		data := &ViaCEP{}
		err = json.Unmarshal(res, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse de resposta: %v\n", err)
		}

		file, err := os.Create("cep.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		bytesData, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}

		file.Write(bytesData)
		fmt.Println("Successfully saved CEP", data.Cep)
	}
}
