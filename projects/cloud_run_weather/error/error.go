package apierror

import (
	"log"
	"net/http"
)

func GetInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	log.Println(http.StatusText(http.StatusInternalServerError), err)
}
