package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Struct to hold current weather data
type WeatherData struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Visibility int `json:"visibility"`
	Sys        struct {
		Sunrise int64 `json:"sunrise"`
		Sunset  int64 `json:"sunset"`
	} `json:"sys"`
}

// Struct to hold forecast data
type ForecastData struct {
	List []struct {
		Datetime int64 `json:"dt"`
		Main     struct {
			TempMin float64 `json:"temp_min"`
			TempMax float64 `json:"temp_max"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"list"`
}

// Fetch current weather
func GetWeather(city, unit, apiKey string) (WeatherData, error) {
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s&units=%s", apiKey, encodedCity, unit)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error fetching weather: %v", err)
	}
	defer resp.Body.Close()

	var data WeatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error decoding response: %v", err)
	}

	return data, nil
}

// Fetch 5-day forecast
func GetForecast(city, unit, apiKey string) (ForecastData, error) {
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?appid=%s&q=%s&units=%s", apiKey, encodedCity, unit)

	resp, err := http.Get(url)
	if err != nil {
		return ForecastData{}, fmt.Errorf("error fetching forecast: %v", err)
	}
	defer resp.Body.Close()

	var data ForecastData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return ForecastData{}, fmt.Errorf("error decoding response: %v", err)
	}

	return data, nil
}
