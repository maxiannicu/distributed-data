package utils

import (
	"log"
	"fmt"
	"os"
)

func NewLogger(prefix string) *log.Logger {
	formattedPrefix := fmt.Sprintf("[%s] ", prefix)
	return log.New(os.Stdout, formattedPrefix, log.LstdFlags|log.Lmicroseconds)
}