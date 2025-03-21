package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

type ExchangeRate struct {
	USDBRL map[string]any `json:"usdbrl"`
}

func exchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request received")

	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	resp := fetchExchangeRate(ctx, w)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
	}

	ex := ExchangeRate{}
	err = json.Unmarshal(body, &ex)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error unmarshalling response body", http.StatusInternalServerError)
	}
	log.Println("Response unmarshalled", ex)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ex)
}

func fetchExchangeRate(ctx context.Context, w http.ResponseWriter) *http.Response {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		w.Write([]byte("Error creating request"))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		w.Write([]byte("Error fetching exchange rate"))
	}
	defer resp.Body.Close()

	return resp
}
