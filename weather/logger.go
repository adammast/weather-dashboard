package weather

import (
	"log"
	"os"
)

// Logger writes to a file only
var Logger *log.Logger

// ConsoleLogger writes to the terminal only
var ConsoleLogger *log.Logger

func init() {
	file, err := os.OpenFile("weather.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}

	Logger = log.New(file, "", log.Ldate|log.Ltime)
	ConsoleLogger = log.New(os.Stderr, "", log.LstdFlags)
}
