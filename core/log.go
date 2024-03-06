package core

import (
	. "github.com/sirupsen/logrus"
)

func InitLogger() *Logger {

	logger := New()
	return logger
}
