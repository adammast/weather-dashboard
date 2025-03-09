package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
	"weather-dashboard/weather"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Use bufio to read user input for city
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

	// Get API key
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("API key is missing! Please set the WEATHER_API_KEY environment variable.")
	}

	var choice int
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Current Weather")
		fmt.Println("2. 5-Day Forecast")
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number (1 or 2).")
			// Clear the input buffer
			reader.ReadString('\n')
			continue
		}

		switch choice {
		case 1:
			fetchAndDisplayCurrentWeather(city, apiKey)
			return
		case 2:
			fetchAndDisplayForecast(city, apiKey)
			return
		default:
			fmt.Println("Invalid option. Please enter 1 or 2.")
		}
	}
}

// Fetch and display current weather
func fetchAndDisplayCurrentWeather(city, apiKey string) {
	weatherData, err := weather.GetWeather(city, apiKey)
	if err != nil {
		log.Fatalf("Error fetching weather: %v", err)
	}

	// Display current weather
	fmt.Println("\nWeather Data for", city)
	fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp)
	fmt.Printf("Humidity: %.2f%%\n", weatherData.Main.Humidity)
	fmt.Printf("Wind Speed: %.2f m/s\n", weatherData.Wind.Speed)
}

// Fetch and display 5-day forecast
func fetchAndDisplayForecast(city, apiKey string) {
	forecast, err := weather.GetForecast(city, apiKey)
	if err != nil {
		log.Fatalf("Error fetching forecast: %v", err)
	}

	// Map to store daily high/low temperatures and conditions
	forecastMap := make(map[string]struct {
		MinTemp   float64
		MaxTemp   float64
		Condition string
	})

	// Process forecast data
	for _, entry := range forecast.List {
		date := time.Unix(entry.Datetime, 0).Format("Mon Jan 02")

		if _, exists := forecastMap[date]; !exists {
			forecastMap[date] = struct {
				MinTemp   float64
				MaxTemp   float64
				Condition string
			}{
				MinTemp:   entry.Main.TempMin,
				MaxTemp:   entry.Main.TempMax,
				Condition: entry.Weather[0].Description,
			}
		} else {
			// Update min/max temps if needed
			current := forecastMap[date]
			if entry.Main.TempMin < current.MinTemp {
				current.MinTemp = entry.Main.TempMin
			}
			if entry.Main.TempMax > current.MaxTemp {
				current.MaxTemp = entry.Main.TempMax
			}
			forecastMap[date] = current
		}
	}

	// Extract dates from the map and sort them
	var sortedDates []string
	for date := range forecastMap {
		sortedDates = append(sortedDates, date)
	}
	sort.Slice(sortedDates, func(i, j int) bool {
		t1, _ := time.Parse("Mon Jan 02", sortedDates[i])
		t2, _ := time.Parse("Mon Jan 02", sortedDates[j])
		return t1.Before(t2)
	})

	// Display 5-day forecast in correct order
	fmt.Println("\n5-Day Forecast for", city)
	for _, date := range sortedDates {
		info := forecastMap[date]
		fmt.Printf("%s - High: %.1f°C, Low: %.1f°C, %s\n", date, info.MaxTemp, info.MinTemp, info.Condition)
	}
}
