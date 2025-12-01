package apierror

import (
	"log/slog"
	"net/http"
)

func InvalidCepError(w http.ResponseWriter) {
	slog.Error("invalid zip code", "status", http.StatusUnprocessableEntity)
	http.Error(w, "invalid zip code", http.StatusUnprocessableEntity)
}

func InternalServerError(w http.ResponseWriter) {
	slog.Error(http.StatusText(http.StatusInternalServerError), "status", http.StatusInternalServerError)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func BadRequestError(w http.ResponseWriter) {
	slog.Error(http.StatusText(http.StatusBadRequest), "status", http.StatusBadRequest)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
