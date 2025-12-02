package utils

import "log"

var logger *log.Logger

func GetLogger() *log.Logger {
	if logger == nil {
		logger = log.Default()
	}

	return logger
}
