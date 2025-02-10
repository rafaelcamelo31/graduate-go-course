package main

import (
	"net/http"
	"sync"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{"Blog"})

	/*
	   Add two goroutines to wait for the two servers to finish.
	*/
	wg := sync.WaitGroup{}
	wg.Add(2)
	/*
	   Must use goroutine to run both servers on different threads. This way, the first server will not block the second one.
	*/
	go func() {
		defer wg.Done()
		http.ListenAndServe(":8080", mux)
	}()

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", HomeHandler2)

	go func() {
		defer wg.Done()
		http.ListenAndServe(":8081", mux2)
	}()

	wg.Wait()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func HomeHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World! 2"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
