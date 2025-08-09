package logger

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type StdoutHook struct {
	LogLevels []log.Level
	Formatter log.Formatter
}

func (hook *StdoutHook) Levels() []log.Level {
	if hook.LogLevels == nil {
		return log.AllLevels
	}
	return hook.LogLevels
}

func (hook *StdoutHook) Fire(entry *log.Entry) error {
	if hook.Formatter == nil {
		hook.Formatter = &log.TextFormatter{}
	}
	formatLog, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	fmt.Println(string(formatLog))
	return nil
}
