/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	lumcfg "github.com/tkdtn2013/lumina/internal/config"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set default configuration for Lumina",
	Run: func(cmd *cobra.Command, args []string) {
		model := args[0]

		cfg := &lumcfg.LuminaConfig{
			DefaultModel: model,
		}

		if err := lumcfg.SaveConfig(cfg); err != nil {
			fmt.Println("❌ Failed to save config:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Configuration saved.")
		fmt.Printf("📦 Default model       : %s\n", cfg.DefaultModel)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
