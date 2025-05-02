/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [model]",
	Short: "Run a local LLM model using Ollama",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]

		// 1. ollama 설치 여부 확인
		if _, err := exec.LookPath("ollama"); err != nil {
			fmt.Fprintln(os.Stderr, "❌ Ollama is not installed or not found in $PATH.")
			fmt.Fprintln(os.Stderr, "👉 Please install it from: https://ollama.com/download")
			os.Exit(1)
		}

		fmt.Printf("🧠 Launching model: %s\n\n", modelName)

		// 2. ollama run 실행
		command := exec.Command("ollama", "run", modelName)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin

		if err := command.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error running model: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
