package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"weather-dashboard/weather"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Use bufio to read user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a city name to get weather info for it: ")
	city, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	city = strings.TrimSpace(city)

	if city == "" {
		log.Fatal("City name cannot be empty.")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")

	if apiKey == "" {
		log.Fatal("API key is missing! Please set the WEATHER_API_KEY environment variable.")
	}

	weatherData, err := weather.GetWeather(city, apiKey)
	if err != nil {
		log.Fatalf("Error fetching weather: %v", err)
	}

	fmt.Println("Weather Data for", city)
	fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp)
	fmt.Printf("Humidity: %.2f%%\n", weatherData.Main.Humidity)
	fmt.Printf("Wind Speed: %.2f m/s\n", weatherData.Wind.Speed)
}
