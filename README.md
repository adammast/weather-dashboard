# Weather Dashboard (Go)

## Description

This project is a simple weather dashboard built using Go. It fetches current weather data from the OpenWeatherMap API and displays it in a clean format.

## Features

- Fetch current weather information for any location.
- Display temperature, humidity, and wind speed.
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
- `encoding/json`: For parsing the JSON data returned by the API.
- `os`: For reading environment variables.
- `github.com/joho/godotenv`: For loading API keys from a .env file.

## License

This project is licensed under the MIT License.