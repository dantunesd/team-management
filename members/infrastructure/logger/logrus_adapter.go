package logger

import (
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

type LogrusAdapter struct {
	logger *logrus.Logger
}

func NewLogrusAdapter() *LogrusAdapter {
	logger := logrus.New()
	logger.SetFormatter(&ecslogrus.Formatter{})
	return &LogrusAdapter{logger: logger}
}

func (l *LogrusAdapter) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogrusAdapter) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogrusAdapter) Print(args ...interface{}) {
	l.logger.Print(args...)
}
