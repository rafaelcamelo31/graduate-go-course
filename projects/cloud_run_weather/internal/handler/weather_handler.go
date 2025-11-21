package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/config"
	"github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/constant"
	apierror "github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/error"
	"github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/internal/entity"
)

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	weatherConfig, _ := config.LoadConfig()

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		http.Error(w, constant.MISSING_CEP_QUERY, http.StatusBadRequest)
		slog.Error(constant.MISSING_CEP_QUERY, "code", http.StatusBadRequest)
		return
	}

	cep := entity.NewViaCEP(cepParam)
	if cep.IsEightDigits() || !cep.IsAllDigits() {
		http.Error(w, constant.INVALID_ZIPCODE, http.StatusUnprocessableEntity)
		slog.Error(constant.INVALID_ZIPCODE, "code", http.StatusUnprocessableEntity)
		return
	}

	resp, _ := http.Get("https://viacep.com.br/ws/" + cepParam + "/json/")
	body, _ := io.ReadAll(resp.Body)
	if strings.Contains(string(body), "erro") {
		http.Error(w, constant.CANNOT_FIND_CEP, http.StatusNotFound)
		slog.Error(constant.CANNOT_FIND_CEP, "code", http.StatusNotFound)
		return
	}
	defer resp.Body.Close()

	err := json.Unmarshal(body, cep)
	if err != nil {
		apierror.GetInternalServerError(w, err)
		return
	}
	slog.Info("successfully unmarshaled viacep", "CEP", cep)

	weatherURL := "https://api.weatherapi.com/v1/current.json"
	params := url.Values{}
	params.Add("key", weatherConfig.ApiKey)
	params.Add("q", cep.City)
	resp, err = http.Get(weatherURL + "?" + params.Encode())
	if err != nil {
		apierror.GetInternalServerError(w, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		slog.Error("WeatherAPI returned status", "code", resp.StatusCode, "body", string(body))
		http.Error(w, "Error from WeatherAPI", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		apierror.GetInternalServerError(w, err)
		return
	}

	wapi := entity.NewWeatherAPI()
	err = json.Unmarshal(body, wapi)
	if err != nil {
		apierror.GetInternalServerError(w, err)
		return
	}
	slog.Info("successfully unmarshaled weather api", "data", wapi)

	temp := entity.NewTemperature(wapi.Current.TempC, wapi.Current.Tempf)
	slog.Info("Temperature in", "city", cep.City, "temperature", fmt.Sprintf("%.1f", temp.TempC))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temp)
}
