package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARN: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// TODO: Create non-formatted logs
func (logger *Logger) Debug(message ...interface{}) {
	logger.debug.Println(message...)
}

func (logger *Logger) Info(message ...interface{}) {
	logger.info.Println(message...)
}

func (logger *Logger) Warn(message ...interface{}) {
	logger.warning.Println(message...)
}

func (logger *Logger) Err(message ...interface{}) {
	logger.err.Println(message...)
}

// TODO: Create format enabled logs
func (logger *Logger) Debugf(format string, message ...interface{}) {
	logger.debug.Printf(format, message...)
}

func (logger *Logger) Infof(format string, message ...interface{}) {
	logger.info.Printf(format, message...)
}

func (logger *Logger) Warnf(format string, message ...interface{}) {
	logger.warning.Printf(format, message...)
}

func (logger *Logger) Errf(format string, message ...interface{}) {
	logger.err.Printf(format, message...)
}
