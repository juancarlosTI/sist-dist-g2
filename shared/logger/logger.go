package logger

import (
	"log"
)

func Info(msg string, fields map[string]interface{}) {
	log.Printf("[INFO] %s | %+v\n", msg, fields)
}

func Error(msg string, fields map[string]interface{}) {
	log.Printf("[ERROR] %s | %+v\n", msg, fields)
}
