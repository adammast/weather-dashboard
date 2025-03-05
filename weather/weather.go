package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func GetWeather(city string, apiKey string) (WeatherData, error) {
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?appid=%s&units=metric&q=%s", apiKey, encodedCity)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error fetching weather: %v", err)
	}

	fmt.Println(resp)

	var data WeatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error decoding response: %v", err)
	}

	return data, nil
}
