package utils

import (
	"log"
	"os"
)

// Реализация стандартного логгера.
type Logger struct {
	*log.Logger
}

// NewLogger — конструктор для создания нового логгера.
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime), // Логирование в stdout с датой и временем
	}
}

// Info — логирование информационных сообщений.
func (l *Logger) Info(msg string) {
	l.Println("INFO:", msg)
}

// Error — логирование ошибок.
func (l *Logger) Error(msg string) {
	l.Println("ERROR:", msg)
}

// Debug — логирование для отладки.
func (l *Logger) Debug(msg string) {
	l.Println("DEBUG:", msg)
}
