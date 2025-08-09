package env

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	magiskBusyBox = "/data/adb/magisk/busybox"
	apatchBusyBox = "/data/adb/ap/bin/busybox"
)

func (e env) getBusybox() string {
	cmd := exec.Command(e.suFile, "-v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatalf("failed check root provider")
	}

	lowerOutput := strings.ToLower(string(output))
	if strings.Contains(lowerOutput, "magisk") {
		return magiskBusyBox
	}
	return "/system/bin/toybox"
}
