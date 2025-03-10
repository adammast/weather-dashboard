# Weather Dashboard (Go)

## Description

This project is a simple weather dashboard built using Go. It fetches weather data from the OpenWeatherMap API and displays it in a clean command line interface.

## Features

- Fetch current weather information for any location.
- Supports both metric and imperial units.
- Fetch and display a 5-day weather forecast with daily high/low temperatures and conditions.
- Display temperature, humidity, visibility, wind speed, and sunrise/sunset times.
- Supports city names with spaces (e.g., "New York City").
- Uses environment variables for API key security.

## Setup

### Prerequisites
- Install [Go](https://go.dev/doc/install) if you haven't already.

### Installation & Running
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/weather-dashboard.git
   cd weather-dashboard
   ```
2. Create a `.env` file in the root directory and add your OpenWeatherMap API key:
   ```ini
   WEATHER_API_KEY=your_api_key_here
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## Libraries Used

- `net/http`: For making HTTP requests to the weather API.
- `net/url`: For encoding city names in API requests.
- `encoding/json`: For parsing the JSON data returned by the API.
- `os`: For reading environment variables.
- `github.com/joho/godotenv`: For loading API keys from a .env file.
- `sort`: For sorting forecast data correctly.
- `time`: For handling and formatting date/time values.

## License

This project is licensed under the MIT License.