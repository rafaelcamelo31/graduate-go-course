package apierror

import (
	"log/slog"
	"net/http"
)

func GetInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	slog.Error(http.StatusText(http.StatusInternalServerError), slog.Any("error", err))
}
