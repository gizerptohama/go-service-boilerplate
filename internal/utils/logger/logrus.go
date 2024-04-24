package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type loggerLogrus struct {
	log *logrus.Logger
}

func (logger *loggerLogrus) Errorf(format string, args ...any) {
	logger.log.Errorf(format, args...)
}
func (logger *loggerLogrus) Error(args ...any) {
	logger.log.Error(args...)
}
func (logger *loggerLogrus) Fatalf(format string, args ...any) {
	logger.log.Fatalf(format, args...)
}
func (logger *loggerLogrus) Fatal(args ...any) {
	logger.log.Fatal(args...)
}
func (logger *loggerLogrus) Warnf(format string, args ...any) {
	logger.log.Warnf(format, args...)
}
func (logger *loggerLogrus) Debugf(format string, args ...any) {
	logger.log.Debugf(format, args...)
}
func (logger *loggerLogrus) Debug(args ...any) {
	logger.log.Debug(args...)
}
func (logger *loggerLogrus) Infof(format string, args ...any) {
	logger.log.Infof(format, args...)
}
func (logger *loggerLogrus) Info(args ...any) {
	logger.log.Info(args...)
}
func (logger *loggerLogrus) WithField(key string, value any) (entry Logger) {
	entry = &LoggerEntry{logger.log.WithField(key, value)}
	return
}
func (logger *loggerLogrus) WithFields(args map[string]any) (entry Logger) {
	entry = &LoggerEntry{logger.log.WithFields(args)}
	return
}
func (logger *loggerLogrus) Printf(format string, args ...any) {
	logger.log.Printf(format, args...)
}
func (logger *loggerLogrus) Print(args ...any) {
	logger.log.Print(args...)
}
func (logger *loggerLogrus) Println(args ...any) {
	logger.log.Println(args...)
}

func NewLogger(writeToFile bool) (logger *loggerLogrus, logFile *os.File) {
	fileName := "app.log"
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		ForceColors:     true,
	})

	writers := []io.Writer{
		os.Stdout,
	}

	if writeToFile {
		logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return logger, nil
		}

		writers = append(writers, logFile)
	}

	multiWriter := io.MultiWriter(writers...)
	log.SetOutput(multiWriter)
	logger = &loggerLogrus{log}

	return
}
