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
	USDBRL CurrencyInfo `json:"usdbrl"`
}

type CurrencyInfo struct {
	Code      string  `json:"code"`
	Codein    string  `json:"codein"`
	Name      string  `json:"name"`
	High      float64 `json:"high,string"`
	Low       float64 `json:"low,string"`
	VarBid    float64 `json:"varBid,string"`
	PctChange float64 `json:"pctChange,string"`
	Bid       float64 `json:"bid,string"`
	Ask       float64 `json:"ask,string"`
	Timestamp string  `json:"timestamp"`
	CreatedAt string  `json:"createdAt"`
}

func NewExchangeRate() *ExchangeRate {
	return &ExchangeRate{
		USDBRL: CurrencyInfo{
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}
}

func (bh *BaseHandler) exchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request received")

	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	resp := fetchExchangeRate(ctx, w)
	if resp == nil {
		log.Println("Response is nil")
		http.Error(w, "Error fetching exchange rate", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	ex := NewExchangeRate()
	err = json.Unmarshal(body, &ex)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error unmarshalling response body", http.StatusInternalServerError)
		return
	}
	log.Println("Response unmarshalled", ex)

	err = insertExchangeRate(ctx, bh.DB, ex)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error inserting exchange rate", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ex)
}

func fetchExchangeRate(ctx context.Context, w http.ResponseWriter) *http.Response {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return nil
	}

	return resp
}
