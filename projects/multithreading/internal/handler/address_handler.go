package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func GetAddress(w http.ResponseWriter, r *http.Request) {
	braChann := make(chan *BrasilAPI)
	viaChann := make(chan *ViaCEPAPI)

	bra := NewBrasilAPI()
	go bra.GetBrasilAPIAddress(w, braChann)

	via := NewViaCEPAPI()
	go via.GetViaCEPAPIAddress(w, viaChann)

	json.NewEncoder(w).Encode("Loading...")

	for {
		select {
		case <-braChann:
			bra.ToJSON()
			return

		case <-viaChann:
			via.ToJSON()
			return

		case <-time.After(time.Second):
			log.Println("Timeout")
			return
		}
	}
}
