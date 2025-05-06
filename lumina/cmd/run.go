/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	lumcfg "github.com/tkdtn2013/lumina/internal/config"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [model]",
	Short: "Run a local LLM model using Ollama",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 모델명 결정
		var modelName string

		if len(args) == 1 {
			modelName = args[0]
		} else {
			cfg, err := lumcfg.LoadConfig()
			if err != nil || cfg.DefaultModel == "" {
				fmt.Println("❌ No model specified and no default_model set in config.")
				os.Exit(1)
			}

			modelName = cfg.DefaultModel
			fmt.Printf("📦 Using default model from config: %s\n", modelName)
		}

		// 2. ollama 설치 여부 확인
		if _, err := exec.LookPath("ollama"); err != nil {
			fmt.Fprintln(os.Stderr, "❌ Ollama is not installed or not found in $PATH.")
			fmt.Fprintln(os.Stderr, "👉 Please install it from: https://ollama.com/download")
			os.Exit(1)
		}

		// 3. 모델 실행
		fmt.Printf("🚀 Running model '%s' with Ollama...\n", modelName)
		run := exec.Command("ollama", "run", modelName)
		run.Stdout = os.Stdout
		run.Stderr = os.Stderr
		run.Stdin = os.Stdin

		if err := run.Run(); err != nil {
			fmt.Println("❌ Failed to run model:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
