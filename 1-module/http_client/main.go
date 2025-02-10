package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	// client := http.Client{Timeout: time.Second}
	// resp, err := client.Get("https://google.com")
	// if err != nil {
	// 	panic(err)
	// }
	customClient := http.Client{Timeout: time.Second}
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")
	resp, err := customClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
