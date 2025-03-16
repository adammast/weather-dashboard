package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"weather-dashboard/weather"

	"github.com/joho/godotenv"
)

var logger = weather.Logger
var consoleLogger = weather.ConsoleLogger

func main() {
	logger.Println("[INFO] Weather dashboard application started.")

	// Try to load .env file (if it exists)
	_ = godotenv.Load()

	// Get API key from environment variables
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		logger.Println("[ERROR] API key is missing.")
		consoleLogger.Fatal("[ERROR] API key is missing! Please set the WEATHER_API_KEY environment variable or define it in a .env file.")
	}

	// Use bufio to read user input for city
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a city name to get weather info for it: ")
	city, err := reader.ReadString('\n')
	if err != nil {
		logger.Printf("[ERROR] Error reading input: %v\n", err)
		consoleLogger.Fatal("[ERROR] Error reading input:", err)
	}

	city = strings.TrimSpace(city)
	if city == "" {
		logger.Println("[ERROR] City name cannot be empty.")
		consoleLogger.Fatal("[ERROR] City name cannot be empty.")
	}

	logger.Printf("[INFO] User entered city: %s\n", city)

	var unitChoice int
	var unit string
	for {
		fmt.Println("Choose unit system:")
		fmt.Println("1. Metric (Celsius, m/s)")
		fmt.Println("2. Imperial (Fahrenheit, mph)")
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scan(&unitChoice)
		if err != nil || (unitChoice != 1 && unitChoice != 2) {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			reader.ReadString('\n')
			continue
		}

		if unitChoice == 1 {
			unit = "metric"
		} else {
			unit = "imperial"
		}
		logger.Printf("[INFO] User selected unit system: %s\n", unit)
		break
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
			logger.Printf("[INFO] User searched for the current weather for city: %s, Unit: %s\n", city, unit)
			fetchAndDisplayCurrentWeather(city, unit, apiKey)
			logger.Println("[INFO] Application exited successfully.")
			return
		case 2:
			logger.Printf("[INFO] User searched for the forecast for city: %s, Unit: %s\n", city, unit)
			fetchAndDisplayForecast(city, unit, apiKey)
			logger.Println("[INFO] Application exited successfully.")
			return
		default:
			fmt.Println("Invalid option. Please enter 1 or 2.")
		}
	}
}

// Fetch and display current weather
func fetchAndDisplayCurrentWeather(city, unit, apiKey string) {
	weatherData, err := weather.GetWeather(city, unit, apiKey)
	if err != nil {
		logger.Printf("[ERROR] Error fetching weather for %s: %v\n", city, err)
		consoleLogger.Fatalf("[ERROR] Error fetching weather: %v", err)
	}

	// Convert visibility from meters to km/miles
	visibility := float64(weatherData.Visibility) / 1000
	if unit == "imperial" {
		visibility = visibility * 0.621371 // Convert km to miles
	}

	// Display current weather
	fmt.Println("\nWeather Data for", city)
	fmt.Printf("Temperature: %.2f°%s\n", weatherData.Main.Temp, getTemperatureUnit(unit))
	fmt.Printf("Feels Like: %.2f°%s\n", weatherData.Main.FeelsLike, getTemperatureUnit(unit))
	fmt.Printf("Humidity: %.2f%%\n", weatherData.Main.Humidity)
	fmt.Printf("Visibility: %.2f %s\n", visibility, getVisibilityUnit(unit))
	fmt.Printf("Wind Speed: %.2f %s\n", weatherData.Wind.Speed, getWindSpeedUnit(unit))
	fmt.Printf("Sunrise: %s\n", formatTime(weatherData.Sys.Sunrise))
	fmt.Printf("Sunset: %s\n", formatTime(weatherData.Sys.Sunset))
}

// Fetch and display 5-day forecast
func fetchAndDisplayForecast(city, unit, apiKey string) {
	forecast, err := weather.GetForecast(city, unit, apiKey)
	if err != nil {
		logger.Printf("[ERROR] Error fetching forecast for %s: %v\n", city, err)
		consoleLogger.Fatalf("[ERROR] Error fetching forecast: %v", err)
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

	tempUnit := getTemperatureUnit(unit)

	// Display 5-day forecast in correct order
	fmt.Println("\n5-Day Forecast for", city)
	for _, date := range sortedDates {
		info := forecastMap[date]
		fmt.Printf("%s - High: %.1f°%s, Low: %.1f°%s, %s\n", date, info.MaxTemp, tempUnit, info.MinTemp, tempUnit, info.Condition)
	}
}

func getTemperatureUnit(unit string) string {
	if unit == "metric" {
		return "C"
	}
	return "F"
}

func getWindSpeedUnit(unit string) string {
	if unit == "metric" {
		return "m/s"
	}
	return "mph"
}

func getVisibilityUnit(unit string) string {
	if unit == "metric" {
		return "km"
	}
	return "miles"
}

func formatTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("15:04 PM MST")
}
