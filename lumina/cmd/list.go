/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available LLM models installed via Ollama",
	Run: func(cmd *cobra.Command, args []string) {

		if _, err := exec.LookPath("ollama"); err != nil {
			fmt.Fprintln(os.Stderr, "❌ Ollama is not installed or not found in $PATH.")
			fmt.Fprintln(os.Stderr, "👉 Please install it from: https://ollama.com/download")
			os.Exit(1)
		}

		out, err := exec.Command("ollama", "list").Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to get model list: %v\n", err)
			os.Exit(1)
		}

		lines := strings.Split(string(out), "\n")
		if len(lines) <= 1 {
			fmt.Println("🤖 No models found.")
			return
		}

		fmt.Println("🧠 Installed models:")
		for _, line := range lines[1:] {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				fmt.Printf("- %s\n", parts[0])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
