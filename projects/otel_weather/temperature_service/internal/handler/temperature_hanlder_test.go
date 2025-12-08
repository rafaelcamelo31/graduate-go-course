package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/temperature_service/config"
	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/temperature_service/constant"
	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/temperature_service/internal/entity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockViaCepAdapter struct {
	mock.Mock
}

func (m *MockViaCepAdapter) GetCityByCep(ctx context.Context, cep string) (*entity.City, error) {
	args := m.Called(ctx, cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.City), args.Error(1)
}

type MockWeatherApiAdapter struct {
	mock.Mock
}

func (m *MockWeatherApiAdapter) GetWeather(ctx context.Context, city string) (*entity.Weather, error) {
	args := m.Called(ctx, city)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Weather), args.Error(1)
}

type TemperatureHandlerTestSuite struct {
	suite.Suite
	mockViacep  *MockViaCepAdapter
	mockWeather *MockWeatherApiAdapter
	cfg         *config.WeatherConfig
	handler     *Handler
}

func (suite *TemperatureHandlerTestSuite) SetupTest() {
	suite.mockViacep = new(MockViaCepAdapter)
	suite.mockWeather = new(MockWeatherApiAdapter)
	suite.cfg = &config.WeatherConfig{}
	suite.handler = NewHandler(suite.mockViacep, suite.mockWeather, suite.cfg)
}

func (suite *TemperatureHandlerTestSuite) TearDownTest() {
	suite.mockViacep.AssertExpectations(suite.T())
	suite.mockWeather.AssertExpectations(suite.T())
}

func (suite *TemperatureHandlerTestSuite) TestGetTemperature_MissingCEP() {
	req := httptest.NewRequest("GET", "/api/temperature", nil)
	w := httptest.NewRecorder()

	suite.handler.GetTemperature(w, req)

	suite.Equal(w.Code, http.StatusBadRequest)
	suite.Contains(w.Body.String(), constant.MISSING_CEP_QUERY)
	suite.mockViacep.AssertNotCalled(suite.T(), "GetCityByCep")
	suite.mockWeather.AssertNotCalled(suite.T(), "GetWeather")
}

func (suite *TemperatureHandlerTestSuite) TestGetTemperature_CepIsNotEightDigits() {
	req := httptest.NewRequest("GET", "/api/temperature?cep=123456789", nil)
	w := httptest.NewRecorder()

	suite.handler.GetTemperature(w, req)

	suite.Equal(w.Code, http.StatusUnprocessableEntity)
	suite.Contains(w.Body.String(), constant.INVALID_ZIPCODE)
	suite.mockViacep.AssertNotCalled(suite.T(), "GetCityByCep")
	suite.mockWeather.AssertNotCalled(suite.T(), "GetWeather")
}

func (suite *TemperatureHandlerTestSuite) TestGetTemperature_NonNumeric() {
	req := httptest.NewRequest("GET", "/api/temperature?cep=999abc88", nil)
	w := httptest.NewRecorder()

	suite.handler.GetTemperature(w, req)

	suite.Equal(w.Code, http.StatusUnprocessableEntity)
	suite.Contains(w.Body.String(), constant.INVALID_ZIPCODE)
	suite.mockViacep.AssertNotCalled(suite.T(), "GetCityByCep")
	suite.mockWeather.AssertNotCalled(suite.T(), "GetWeather")
}

func (suite *TemperatureHandlerTestSuite) TestGetTemperature_CityNotFound() {
	suite.mockViacep.On("GetCityByCep", mock.Anything, "99999999").Return(nil, nil).Once()

	req := httptest.NewRequest("GET", "/api/temperature?cep=99999999", nil)
	w := httptest.NewRecorder()

	suite.handler.GetTemperature(w, req)

	suite.Equal(w.Code, http.StatusNotFound)
	suite.mockViacep.AssertCalled(suite.T(), "GetCityByCep", mock.Anything, "99999999")
	suite.mockWeather.AssertNotCalled(suite.T(), "GetWeather")
}

func (suite *TemperatureHandlerTestSuite) TestGetTemperature_ViaCepError() {
	suite.mockViacep.On("GetCityByCep", mock.Anything, "99999999").Return(nil, fmt.Errorf("%s", http.StatusText(http.StatusInternalServerError))).Once()

	req := httptest.NewRequest("GET", "/api/temperature?cep=99999999", nil)
	w := httptest.NewRecorder()

	suite.handler.GetTemperature(w, req)

	suite.Equal(w.Code, http.StatusInternalServerError)
	suite.mockViacep.AssertCalled(suite.T(), "GetCityByCep", mock.Anything, "99999999")
	suite.mockWeather.AssertNotCalled(suite.T(), "GetWeather")
}

func (suite *TemperatureHandlerTestSuite) TestGetTemperature_WeatherApiError() {
	city := entity.NewCity("13083970")
	city.Name = "Campinas"
	suite.mockViacep.On("GetCityByCep", mock.Anything, "13083970").Return(city, nil).Once()
	suite.mockWeather.On("GetWeather", mock.Anything, city.Name).Return(nil, fmt.Errorf("%s", http.StatusText(http.StatusInternalServerError))).Once()

	req := httptest.NewRequest("GET", "/api/temperature?cep=13083970", nil)
	w := httptest.NewRecorder()

	suite.handler.GetTemperature(w, req)

	suite.Equal(w.Code, http.StatusInternalServerError)
	suite.mockViacep.AssertCalled(suite.T(), "GetCityByCep", mock.Anything, "13083970")
	suite.mockWeather.AssertCalled(suite.T(), "GetWeather", mock.Anything, city.Name)
}

func (suite *TemperatureHandlerTestSuite) TestGetTemperature_WeatherNotFound() {
	city := &entity.City{
		Cep:  "00000000",
		Name: "abcdefg",
	}
	suite.mockViacep.On("GetCityByCep", mock.Anything, city.Cep).Return(city, nil).Once()
	suite.mockWeather.On("GetWeather", mock.Anything, city.Name).Return(nil, nil).Once()

	req := httptest.NewRequest("GET", "/api/temperature?cep=00000000", nil)
	w := httptest.NewRecorder()

	suite.handler.GetTemperature(w, req)

	suite.Equal(w.Code, http.StatusNotFound)
	suite.mockViacep.AssertCalled(suite.T(), "GetCityByCep", mock.Anything, city.Cep)
	suite.mockWeather.AssertCalled(suite.T(), "GetWeather", mock.Anything, city.Name)
}

func (suite *TemperatureHandlerTestSuite) TestGetTemperature_Success() {
	city := &entity.City{
		Cep:  "13083970",
		Name: "Campinas",
	}
	weather := &entity.Weather{
		Current: &entity.Current{
			TempC: 25.5,
			Tempf: 77.9,
		},
	}
	suite.mockViacep.On("GetCityByCep", mock.Anything, city.Cep).Return(city, nil).Once()
	suite.mockWeather.On("GetWeather", mock.Anything, city.Name).Return(weather, nil).Once()

	req := httptest.NewRequest("GET", "/api/temperature?cep=13083970", nil)
	w := httptest.NewRecorder()

	suite.handler.GetTemperature(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockViacep.AssertCalled(suite.T(), "GetCityByCep", mock.Anything, city.Cep)
	suite.mockWeather.AssertCalled(suite.T(), "GetWeather", mock.Anything, city.Name)

	temperature := &entity.Temperature{}
	json.Unmarshal(w.Body.Bytes(), temperature)
	suite.Equal(float32(25.5), temperature.TempC)
	suite.Equal(float32(77.9), temperature.TempF)
	suite.Equal(float32(298.5), temperature.TempK)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TemperatureHandlerTestSuite))
}
