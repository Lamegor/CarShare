package logger

import (
	"log"
	"os"
)

var (
	successLog *log.Logger
	errorLog   *log.Logger
)

func initLogs() {
	successFile, err := os.OpenFile("/var/log/success_log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("не удалось открыть файл success_log: %v", err)
	}

	errorFile, err := os.OpenFile("/var/log/error_log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("не удалось открыть файл error_log: %v", err)
	}

	successLog = log.New(successFile, "SUCCESS: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func logSuccess(message string) {
	successLog.Println(message)
}

func logError(message string) {
	errorLog.Println(message)
}
