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
- Logs application info and errors to a separate file.
- Supports running in a Docker container.

## Setup

### Prerequisites
- Install [Go](https://go.dev/doc/install) if you haven't already.
- Install [Docker](https://docs.docker.com/get-docker/) if you want to run the application in a container.

### Installation & Running (Locally)
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

### Running with Docker
1. Build the Docker image:
   ```bash
   docker build -t weather-dashboard .
   ```
2. Run the container with your API key:
   ```bash
   docker run --rm -it --env WEATHER_API_KEY=your_api_key weather-dashboard
   ```

## Libraries Used

- `net/http`: For making HTTP requests to the weather API.
- `net/url`: For encoding city names in API requests.
- `encoding/json`: For parsing the JSON data returned by the API.
- `os`: For reading environment variables.
- `github.com/joho/godotenv`: For loading API keys from a .env file.
- `sort`: For sorting forecast data correctly.
- `time`: For handling and formatting date/time values.
- `log`: For handling logging to a file and to the console.

## License

This project is licensed under the MIT License.