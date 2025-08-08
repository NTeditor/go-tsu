package env

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (e env) NewShell(shell string) *exec.Cmd {
	suFile := fmt.Sprintf("%s/bin/su", e.termuxPrefix)
	envVars := strings.Join(e.genEnvVars(), " ")
	cmd := exec.Command(suFile, "-i", "-c",
		fmt.Sprintf("%s env -i %s %s", e.getBusybox(suFile), envVars, shell))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd
}
