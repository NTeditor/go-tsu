package env

import (
	"strings"
	"os/exec"
)

const (
	magiskBusyBox = "/data/adb/magisk/busybox"
	apatchBusyBox = "/data/adb/ap/bin/busybox"
)

func (e env) getBusybox(suFile string) string {
	cmd := exec.Command(suFile, "-v")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	lowerOutput := strings.ToLower(string(output))
	if strings.Contains(lowerOutput, "magisk") {
		return magiskBusyBox
	}
	panic("Root provider not suppport")
}
