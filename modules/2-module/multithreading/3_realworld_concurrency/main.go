package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

/*
Understanding Mutex and atomic operations

Load test using apache bench
ab -n 10000 -c 100 http://localhost:3000/

Dedect race condition
go run -race .
*/
func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write(fmt.Appendf(nil, "You are a visitor number %d", number))
	})
	http.ListenAndServe(":3000", nil)
}
