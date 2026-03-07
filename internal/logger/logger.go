package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(message string)
	Error(message string)
}

type StdLogger struct {
	logger *log.Logger
}

func NewStdLogger() *StdLogger {
	return &StdLogger{
		logger: log.New(os.Stdout, "[API] ", log.LstdFlags),
	}
}

func (l *StdLogger) Info(message string) {
	l.logger.Printf("INFO: %s", message)
}

func (l *StdLogger) Error(message string) {
	l.logger.Printf("ERROR: %s", message)
}
