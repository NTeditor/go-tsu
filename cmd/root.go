package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nteditor/go-tsu/internal/env"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	shell   string
	user    string
	command string
)

var (
	Version   = "1.0.0"
	BuildType = "user"
)

var rootCmd = &cobra.Command{
	Use: "go-tsu",
	Run: func(cmd *cobra.Command, args []string) {
		termuxFS := "/data/data/com.termux/files"
		termuxPrefix := filepath.Join(termuxFS, "usr")
		env := env.NewEnv(termuxFS, termuxPrefix, user)
		if command != "nil" {
			command, err := env.NewCommand(shell, command)
			if err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Fatalf("failed execute command")
			}

			if err := command.Run(); err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Fatalf("failed execute command")
			}
		} else {
			shell, err := env.NewShell(shell)
			if err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Fatalf("failed execute shell")
			}

			if err := shell.Run(); err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Fatalf("failed execute shell")
			}
		}
	},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`---------------- go-tsu ----------------

Version: %s
Build: %s
License: GPL v3.0

Libraries used:
  • cobra: v1.9.1 (License: Apache 2.0)
  • logrus: v1.9.3 (License: MIT)

-----------------------------------------
`, Version, BuildType)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&shell, "shell", "s", os.Getenv("SHELL"), "Run SHELL instead of the current shell")
	rootCmd.Flags().StringVarP(&user, "user", "u", "root", "Run shell as USER")
	rootCmd.Flags().StringVarP(&command, "command", "c", "nil", "Pass COMMAND to the invoked shell")
	rootCmd.AddCommand(versionCmd)
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Debugf("failed to parse arguments")
	}
}
