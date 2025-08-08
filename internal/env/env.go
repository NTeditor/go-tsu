package env

import "fmt"

const androidPaths = "/product/bin:/apex/com.android.runtime/bin:/apex/com.android.art/bin:/system_ext/bin:/system/bin:/system/xbin:/odm/bin:/vendor/bin:/vendor/xbin"

type env struct {
	termuxFS     string
	termuxPrefix string
	termuxTMP    string
	user         string
}

func NewEnv(termuxFS string, termuxPrefix string, user string) *env {
	return &env{
		termuxFS:     termuxFS,
		termuxPrefix: termuxPrefix,
		termuxTMP:    fmt.Sprintf("%s/tmp", termuxFS),
		user:         user,
	}
}

func (e env) genEnvVars() []string {
	home := fmt.Sprintf("%s/users/%s", e.termuxFS, e.user)
	path := fmt.Sprintf("%s/bin:%s", e.termuxPrefix, androidPaths)
	return []string{
		"HOME=" + home,
		"PATH=" + path,
		"TERM=" + "xterm-256color",
		"PREFIX=" + e.termuxPrefix,
		"PWD=" + home,
		"COLORTERM=" + "truecolor",
		"TMPDIR=" + e.termuxTMP,
		"TERMUX__HOME=" + home,
		"TERMUX__PREFIX=" + e.termuxPrefix,
		"TERMUX__ROOTFS_DIR=" + e.termuxFS,
	}
}
