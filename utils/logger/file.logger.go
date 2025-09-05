package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	logger zerolog.Logger
	mu     sync.Mutex
)

func init() {
	SetupLogger()
}

func SetupLogger() {
	mu.Lock()
	defer mu.Unlock()

	filename := fmt.Sprintf("logs/%v.log", "auth")

	// Set up lumberjack logger for log rotation with compression
	logFile := &lumberjack.Logger{
		Filename: filename, // Use filename with date
		MaxSize:  200,      // Maximum size in megabytes before rotation
		//MaxBackups: 10,     // Maximum number of old log files to retain, use this to omit old log files
		Compress:  true, // Compress rotated log files
		LocalTime: true, // Use local time for file timestamps
	}

	logger = zerolog.New(logFile).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{
		Out:        logFile,
		NoColor:    true,
		TimeFormat: "2006-01-02 15:04:05",
		FormatLevel: func(i interface{}) string {
			if ll, ok := i.(string); ok {
				return strings.ToUpper(ll)
			}
			return ""
		},
	})
	fmt.Println("Logger reinitialized with filename:", logFile.Filename)
}

func LogErrorWithDepth(data interface{}) {
	defer func() {
		if r := recover(); r != nil {
			// Silent recovery - doesn't affect the main request
			logger.Error().Msg("There was an error happened in logger.")
			fmt.Printf("Recovered in logger: %v\n", r)
		}
	}()
	errorData := data.(map[string]interface{})
	err := errorData["error"].(error)
	depth := errorData["depth"].(int)
	message := errorData["message"].(string)

	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "unknown"
		line = 0
	}

	customMsg := fmt.Sprintf("File: %s, Line: %d, ErrorMessage: \"%s\", ErrorDetails: \"%v\"", file, line, message, err)
	logger.Error().Msg(customMsg)
}

func LogInfo(handle string, msg string) {
	defer func() {
		if r := recover(); r != nil {
			// Silent recovery - doesn't affect the main request
			logger.Error().Msg("There was an error happened in logger.")
			fmt.Printf("Recovered in logger: %v\n", r)
		}
	}()

	customMsg := fmt.Sprintf("Service: %s, Handle: \"%s\", InfoDetail: \"%v\"", os.Getenv("LOG_FILE"), handle, msg)
	logger.Info().Msg(customMsg)
}
