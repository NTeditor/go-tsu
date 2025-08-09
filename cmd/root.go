package cmd

import (
	"fmt"
	"os"

	"github.com/nteditor/go-tsu/internal/env"
	"github.com/spf13/cobra"
)

var (
	shell   string
	user    string
	command string
)

var version = "0.0.0"

var rootCmd = &cobra.Command{
	Use: "go-tsu",
	Run: func(cmd *cobra.Command, args []string) {
		termuxFS := "/data/data/com.termux/files"
		termuxPrefix := fmt.Sprintf("%s/usr", termuxFS)
		env := env.NewEnv(termuxFS, termuxPrefix, user)
		if command != "nil" {
			command := env.RunCommand(shell, command)
			if err := command.Run(); err != nil {
				panic(err)
			}
		} else {
			shell := env.NewShell(shell)
			if err := shell.Run(); err != nil {
				panic(err)
			}
		}
	},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Версия: %v\nЛицензия: GPL 3.0\n", version)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&shell, "shell", "s", os.Getenv("SHELL"), "")
	rootCmd.Flags().StringVarP(&user, "user", "u", "root", "")
	rootCmd.Flags().StringVarP(&command, "command", "c", "nil", "")
	rootCmd.DisableFlagParsing = false
	rootCmd.AddCommand(versionCmd)
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
