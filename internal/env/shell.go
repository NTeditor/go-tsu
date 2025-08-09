package env

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (e env) NewShell(shell string) *exec.Cmd {
	envVars := strings.Join(e.genEnvVars(), " ")
	cmd := exec.Command(e.suFile, "-i", "-c",
		fmt.Sprintf("%s env -i %s %s", e.getBusybox(), envVars, shell))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	log.WithFields(log.Fields{
		"envVars": envVars,
		"cmd": cmd,
	}).Debugf("shell created")
	return cmd
}

func (e env) NewCommand(shell string, command string) *exec.Cmd {
	envVars := strings.Join(e.genEnvVars(), " ")
	cmd := exec.Command(e.suFile, "-i", "-c",
		fmt.Sprintf("%s env -i %s %s -c %s", e.getBusybox(), envVars, shell, command))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	log.WithFields(log.Fields{
		"envVars": envVars,
		"cmd": cmd,
	}).Debugf("command created")
	return cmd
}
