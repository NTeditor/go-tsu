package env

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (e env) NewShell(shell string) (*exec.Cmd, error) {
	return e.NewCommand(shell, "")
}

func (e env) NewCommand(shell string, command string) (*exec.Cmd, error) {
	envVars := strings.Join(e.genEnvVars(), " ")
	busyBox, err := e.getBusybox()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Errorf("failed get busyBox")
		return nil, fmt.Errorf("failed get busyBox: %v", err)
	}

	cmd := exec.Command(e.suFile, e.user, "-i", "-c",
		fmt.Sprintf("%s env -i %s %s -c %s", busyBox, envVars, shell, command),
	)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	log.WithFields(log.Fields{
		"envVars": envVars,
		"cmd":     cmd,
		"busyBox": busyBox,
	}).Debugf("command created")
	return cmd, nil
}
