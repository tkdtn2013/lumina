/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check running model processes",
	Run: func(cmd *cobra.Command, args []string) {
		usr, _ := user.Current()
		pidFile := filepath.Join(usr.HomeDir, ".lumina", "pid", "serve.pid")

		data, err := os.ReadFile(pidFile)
		if err != nil {
			fmt.Println("❌ Ollama server not running (no PID found).")
			return
		}

		pid := string(data)
		fmt.Printf("📦 Ollama server running (PID %s)\n", pid)

		// API 응답 테스트
		resp, err := http.Get("http://localhost:11434")
		if err != nil {
			fmt.Println("⚠️  API server not responding.")
		} else {
			fmt.Printf("✅ API server reachable. Status: %s\n", resp.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
