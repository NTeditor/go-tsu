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

var Version = "1.0.0-rc1"

var rootCmd = &cobra.Command{
	Use: "go-tsu",
	Run: func(cmd *cobra.Command, args []string) {
		termuxFS := "/data/data/com.termux/files"
		termuxPrefix := filepath.Join(termuxFS, "usr")
		env := env.NewEnv(termuxFS, termuxPrefix, user)
		if command != "nil" {
			command := env.NewCommand(shell, command)
			if err := command.Run(); err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Errorf("failed execute command")
			}
		} else {
			shell := env.NewShell(shell)
			if err := shell.Run(); err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Errorf("failed execute shell")
			}
		}
	},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`--- go-tsu ---

Version: %s
License: GPL v3.0

Libraries used:
  • cobra: v1.9.1 (License: Apache 2.0)
  • logrus: v1.9.3 (License: MIT)

----------------------
`, Version)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&shell, "shell", "s", os.Getenv("SHELL"), "")
	rootCmd.Flags().StringVarP(&user, "user", "u", "root", "")
	rootCmd.Flags().StringVarP(&command, "command", "c", "nil", "")
	rootCmd.AddCommand(versionCmd)
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Errorf("failed to parse arguments")
	}
}

