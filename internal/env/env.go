package env

import (
	"fmt"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const androidBinPath = "/product/bin:/apex/com.android.runtime/bin:/apex/com.android.art/bin:/system_ext/bin:/system/bin:/system/xbin:/odm/bin:/vendor/bin:/vendor/xbin"

type env struct {
	termuxFS     string
	termuxPrefix string
	termuxTMP    string
	user         string
	suFile       string
}

func NewEnv(termuxFS string, termuxPrefix string, user string) *env {
	suFile := filepath.Join(termuxPrefix, "bin", "su")
	termuxTMP := filepath.Join(termuxPrefix, "tmp")
	log.WithFields(log.Fields{
		"termuxFS":     termuxFS,
		"termuxPrefix": termuxPrefix,
		"termuxTMP":    termuxTMP,
		"user":         user,
		"suFile":       suFile,
	}).Debugf("new env{}")
	return &env{
		termuxFS:     termuxFS,
		termuxPrefix: termuxPrefix,
		termuxTMP:    filepath.Join(termuxPrefix, "tmp"),
		user:         user,
		suFile: suFile,
	}
}

func (e env) genEnvVars() []string {
	home := filepath.Join(e.termuxFS, "users", e.user)
	termuxBinPath := filepath.Join(e.termuxPrefix, "bin")
	binPath := fmt.Sprintf("%s:%s", termuxBinPath, androidBinPath)
	envVars := []string{
		"HOME=" + home,
		"PATH=" + binPath,
		"TERM=" + "xterm-256color",
		"PREFIX=" + e.termuxPrefix,
		"PWD=" + home,
		"COLORTERM=" + "truecolor",
		"TMPDIR=" + e.termuxTMP,
		"TERMUX__HOME=" + home,
		"TERMUX__PREFIX=" + e.termuxPrefix,
		"TERMUX__ROOTFS_DIR=" + e.termuxFS,
	}
	log.WithFields(log.Fields{
		"envVars": envVars,
	}).Info("generated environment variables")
	return envVars
}
