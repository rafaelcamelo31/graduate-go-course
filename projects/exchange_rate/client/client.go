package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ExchangeRateResponse struct {
	USDBRL struct {
		Bid float64 `json:"bid,string"`
	} `json:"USDBRL"`
}

type BidFileFormat struct {
	USD float64 `json:"USD"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/exchange-rate", nil)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Response body", string(body))

	ex := &ExchangeRateResponse{}
	err = json.Unmarshal(body, &ex)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Response unmarshalled %+v\n", ex)

	file, err := os.Create("./client/rate.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	fileFormat := &BidFileFormat{
		USD: ex.USDBRL.Bid,
	}
	bytesData, err := json.Marshal(fileFormat)
	if err != nil {
		log.Println(err)
	}

	file.Write(bytesData)
	log.Println("Successfully saved exchange rate", fileFormat)
}
