package infra

import (
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

type LoggerAdapter struct {
	logger *logrus.Logger
}

func NewLoggerAdapter() *LoggerAdapter {
	logger := logrus.New()
	logger.SetFormatter(&ecslogrus.Formatter{})
	return &LoggerAdapter{logger: logger}
}

func (l *LoggerAdapter) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LoggerAdapter) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LoggerAdapter) Print(args ...interface{}) {
	l.logger.Print(args...)
}
