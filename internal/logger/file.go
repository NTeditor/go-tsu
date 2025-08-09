package logger

import (
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)





type FileHook struct {
	LogLevels []log.Level
	Formatter log.Formatter
	File io.Writer
}

func (hook *FileHook) Levels() []log.Level {
	if hook.LogLevels == nil {
		return log.AllLevels
	}
	return hook.LogLevels
}

func (hook *FileHook) Fire(entry *log.Entry) error {
	if hook.Formatter == nil {
		hook.Formatter = &log.JSONFormatter{}
	}
	formatLog, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	if hook.File == nil {
		return fmt.Errorf("FileHook.File not set")
	}
	_, err = hook.File.Write(formatLog)
	return err
}

