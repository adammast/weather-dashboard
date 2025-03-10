package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// WeatherData represents the response for current weather
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

// ForecastData represents the response for the 5-day forecast
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
		Logger.Printf("[ERROR] Failed to fetch weather data for %s: %v\n", city, err)
		return WeatherData{}, fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		Logger.Printf("[ERROR] Received status code %d for %s\n", resp.StatusCode, city)
		return WeatherData{}, fmt.Errorf("received unexpected status code %d", resp.StatusCode)
	}

	var data WeatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		Logger.Printf("[ERROR] Error decoding JSON for %s: %v\n", city, err)
		return WeatherData{}, fmt.Errorf("error decoding JSON response: %w", err)
	}

	Logger.Printf("[INFO] Successfully fetched weather data for %s", city)
	return data, nil
}

// Fetch 5-day forecast
func GetForecast(city, unit, apiKey string) (ForecastData, error) {
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?appid=%s&q=%s&units=%s", apiKey, encodedCity, unit)

	resp, err := http.Get(url)
	if err != nil {
		Logger.Printf("[ERROR] Failed to fetch forecast data for %s: %v\n", city, err)
		return ForecastData{}, fmt.Errorf("failed to fetch forecast data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		Logger.Printf("[ERROR] Received status code %d for %s forecast\n", resp.StatusCode, city)
		return ForecastData{}, fmt.Errorf("received unexpected status code %d", resp.StatusCode)
	}

	var data ForecastData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		Logger.Printf("[ERROR] Error decoding JSON for forecast of %s: %v\n", city, err)
		return ForecastData{}, fmt.Errorf("error decoding JSON response: %v", err)
	}

	Logger.Printf("[INFO] Successfully fetched forecast data for %s", city)
	return data, nil
}
