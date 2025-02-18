package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
)

var (
	Logger *slog.Logger
	once   sync.Once
)

func InitLogger() {
	once.Do(func() {
		logDir := "logs"
		if _, err := os.Stat(logDir); os.IsNotExist(err) {
			err := os.Mkdir(logDir, os.ModePerm)
			if err != nil {
				fmt.Printf("Failed to create logs directory: %v\n", err)
				return
			}
		}

		// Generate log file name with date-time stamp
		currentTime := time.Now().Format("2006-01-02_15-04-05")
		logFileName := filepath.Join(logDir, fmt.Sprintf("logfile_%s.log", currentTime))

		// Set up lumberjack logger for log rotation
		logFile := &lumberjack.Logger{
			Filename:   logFileName,
			MaxSize:    10, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   // days
			Compress:   true, // disabled by default
		}

		// Create a TextHandler with the file as the output
		handler := slog.NewTextHandler(logFile, nil)

		// Create a new logger with the handler
		Logger = slog.New(handler)

		Logger.Info("Logger initialised successfully")
	})
}
