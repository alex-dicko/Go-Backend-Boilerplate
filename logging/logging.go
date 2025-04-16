package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Warning
	Error
	Success
	Info
)

type Logger struct {
	Name string	
}

func (l *Logger) Log(level LogLevel, message string) {
	output := fmt.Sprintf("[%s]%s - %s \n", l.Name, verboseName(level), message)
	fmt.Printf(output)
	output = "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + output
	err := writeToFile(l, output)
	if err != nil {
		panic(err)
	}
}

func writeToFile(l *Logger, output string) error {
	logDir := "logs/" + l.Name                 // This is used as folder name
	logFilePath := filepath.Join(logDir, "log.txt")

	// Ensure the folder exists (creates folder if it doesn't exist)
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	// Open the log file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	// Write the log output
	_, err = file.Write([]byte(output))
	if err != nil {
		return fmt.Errorf("failed to write to log file: %w", err)
	}

	return nil
}

func verboseName(level LogLevel) string {
	if level == Error{
		return "[ERROR]"
	}
	if level == Warning {
		return "[WARNING]"
	}
	if level == Debug	 {
		return "[ERROR]"
	}
	if level == Success {
		return "[SUCCESS]"
	}
	if level == Info {
		return "[INFO]"
	}

	return "[NONE]"
}

func InitLogger(name string) *Logger {
	return &Logger{Name: name}
}

