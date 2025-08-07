package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const termuxFS = "/data/data/com.termux/files"

var (
	termuxPrefix = fmt.Sprintf("%s/usr", termuxFS)
	suPath       = fmt.Sprintf("%s/bin/su", termuxPrefix)
)

func main() {
	busyboxPath, err := getBusyBox()
	if err != nil {
		panic(err)
	}

	args := fmt.Sprintf("%s env -i %s %s", busyboxPath, getEnvVars(), getUserShell())
	cmd := exec.Command(suPath, "-i", "-c", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func getEnvVars() string {
	androidPaths := "/product/bin:/apex/com.android.runtime/bin:/apex/com.android.art/bin:/system_ext/bin:/system/bin:/system/xbin:/odm/bin:/vendor/bin:/vendor/xbin"

	home := fmt.Sprintf("%s/root", termuxFS)
	path := fmt.Sprintf("%s/bin:%s", termuxPrefix, androidPaths)
	tmpdir := fmt.Sprintf("%s/tmp", termuxPrefix)
	return strings.Join([]string{
		"HOME=" + home,
		"PATH=" + path,
		"TERM=" + "xterm-256color",
		"PREFIX=" + termuxPrefix,
		"TERMUX__ROOTFS_DIR=" + termuxFS,
		"TERMUX__PREFIX="+ termuxPrefix,
		"TERMUX__HOME=" + home,
		"PWD=" + home,
		"COLORTERM=" + "truecolor",
		"TMPDIR=" + tmpdir,
	}, " ")
}

func getUserShell() string {
	return os.Getenv("SHELL")
}

func getBusyBox() (string, error) {
	magiskBusyBox := "/data/adb/magisk/busybox"
	apatchBusyBox := "/data/adb/ap/bin/busybox"

	cmd := exec.Command(suPath, "-v")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lowerOutput := strings.ToLower(string(output))

	if strings.Contains(lowerOutput, "magisk") {
		return magiskBusyBox, nil
	} else if strings.Contains(lowerOutput, "apatch") {
		return apatchBusyBox, nil
	}

	return "", err
}
