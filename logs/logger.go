package logs

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	level     string
	logToFile bool
	file      *os.File
}

const (
	InfoLevel    = "info"
	WarningLevel = "warning"
	ErrorLevel   = "error"
	DebugLevel   = "debug"
)

func NewLogger(level string, logToFile bool, filePath string) (*Logger, error) {
	var file *os.File
	if logToFile {
		var err error
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return nil, err
		}
	}
	return &Logger{
		level:     level,
		logToFile: logToFile,
		file:      file,
	}, nil
}

// log will write a log entry with level, message, and time
func (l *Logger) log(level string, message string) {
	if l.shouldLog(level) {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		logMessage := fmt.Sprintf("%s - %s: %s", timestamp, level, message)
		if l.logToFile {
			fmt.Fprintln(l.file, logMessage)
		}
		fmt.Println(logMessage)
	}
}

// shouldLog determines if the current log level should be logged
func (l *Logger) shouldLog(level string) bool {
	levels := map[string]int{
		InfoLevel:    1,
		WarningLevel: 2,
		ErrorLevel:   3,
		DebugLevel:   4,
	}
	return levels[level] >= levels[l.level]
}

// Info logs an info message
func (l *Logger) Info(message string) {
	l.log(InfoLevel, message)
}

// Warning logs a warning message
func (l *Logger) Warning(message string) {
	l.log(WarningLevel, message)
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.log(ErrorLevel, message)
}

// Debug logs a debug message
func (l *Logger) Debug(message string) {
	l.log(DebugLevel, message)
}

// Close closes the log file if it was opened
func (l *Logger) Close() {
	if l.logToFile && l.file != nil {
		l.file.Close()
	}
}
