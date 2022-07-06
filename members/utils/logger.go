package utils

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
}
