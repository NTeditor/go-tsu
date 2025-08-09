package main

import (
	"io"
	"os"
	"runtime"

	"github.com/nteditor/go-tsu/cmd"
	"github.com/nteditor/go-tsu/internal/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
	close := initLogger()
	defer close()
	cmd.Exec()
}

func initLogger() func() error {
	log.AddHook(&logger.StdoutHook{
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
		Formatter: &log.TextFormatter{
			ForceColors: true,
			DisableTimestamp: true,
			DisableLevelTruncation: true,
		},
	})

	log.SetOutput(io.Discard)

	if runtime.GOOS != "android" {
		log.WithFields(log.Fields{
			"runtime.GOOS": runtime.GOOS,
		}).Fatalf("runtime.GOOS not equals android")
	}

	logFile := "/data/data/com.termux/files/usr/var/log/go-tsu.json"

	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Fatalf("failed create file: %s", logFile)
	}

	log.AddHook(&logger.FileHook{
		LogLevels: log.AllLevels,
		Formatter: &log.JSONFormatter{
			PrettyPrint: true,
		},
		File: file,
	})
	return file.Close
}

