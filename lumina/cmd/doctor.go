/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	lumcfg "github.com/tkdtn2013/lumina/internal/config"

	"github.com/spf13/cobra"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose Lumina system environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🩺 Lumina Doctor Report")

		// 1. Ollama 설치 확인
		_, err := exec.LookPath("ollama")
		if err != nil {
			fmt.Println("❌ Ollama is NOT installed.")
		} else {
			fmt.Println("✅ Ollama is installed.")
		}

		// 2. 설치된 모델 목록
		out, err := exec.Command("ollama", "list").Output()
		if err != nil || len(strings.TrimSpace(string(out))) == 0 {
			fmt.Println("⚠️  No models found.")
		} else {
			fmt.Println("✅ Models installed:")
			lines := strings.Split(string(out), "\n")
			for _, line := range lines[1:] {
				parts := strings.Fields(line)
				if len(parts) > 0 {
					fmt.Printf("   - %s\n", parts[0])
				}
			}
		}

		// 3. 사용자 설정 확인
		fmt.Println("🔧 Checking config file...")
		cfg, err := lumcfg.LoadConfig()

		if err != nil {
			fmt.Println("❌ Config file not found or unreadable.")
			return
		}

		if cfg.DefaultModel != "" {
			fmt.Printf("✅ Default model set to: %s\n", cfg.DefaultModel)
		} else {
			fmt.Println("⚠️  Default model is not set.")
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doctorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doctorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
