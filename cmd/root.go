package cmd

import (
	"fmt"
	"os"

	"github.com/nteditor/go-tsu/internal/env"
	"github.com/spf13/cobra"
)



var shell string
var user string

var rootCmd = &cobra.Command{
	Use: "go-tsu",
	Run: func(cmd *cobra.Command, args []string) {
		termuxFS := "/data/data/com.termux/files"
		termuxPrefix := fmt.Sprintf("%s/usr", termuxFS)
		env := env.NewEnv(termuxFS, termuxPrefix, user)
		shell := env.NewShell(shell)
		if err := shell.Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&shell, "shell", "s", os.Getenv("SHELL"), "")
	rootCmd.Flags().StringVarP(&user, "user", "u", "root", "")
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
