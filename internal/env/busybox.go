package env

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	busybox = "/data/adb/magisk/busybox"
	toybox  = "/data/adb/ap/bin/busybox"
)

var ErrBusyBoxNotFound = errors.New("busybox not found")

func (e env) getBusybox() (string, error) {
	var errBusybox error
	var errToybox error
	if _, errBusybox = os.Stat(busybox); errBusybox == nil {
		log.WithFields(log.Fields{
			"busybox": busybox,
		}).Infof("return busybox")
		return busybox, nil
	}
	if _, errToybox = os.Stat(toybox); errToybox == nil {
		log.WithFields(log.Fields{
			"busybox": toybox,
		}).Infof("return busybox")
		return toybox, nil
	}

	log.WithFields(log.Fields{
		"busybox": errBusybox,
		"toybox":  errToybox,
	}).Errorf("busybox not found")
	return "", fmt.Errorf("%v, %s: %v, %s: %v", ErrBusyBoxNotFound, busybox, errBusybox, toybox, errToybox)
}
