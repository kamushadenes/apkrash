package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"os"

	"github.com/spf13/cobra"
)

var outputFormat string
var useColor bool
var onlyDiffs bool
var includeFiles bool
var decompile bool
var email string
var password string

var rootCmd = &cobra.Command{
	Use:   "apkrash",
	Short: "Android APK security analysis toolkit",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if outputFormat != "text" && outputFormat != "json" && outputFormat != "json_pretty" && outputFormat != "table" {
			return fmt.Errorf("invalid output format: %s", outputFormat)
		}

		color.NoColor = !useColor

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "format", "o", "text", "Output format, one of text, json, json_pretty, table")
	rootCmd.PersistentFlags().BoolVarP(&useColor, "color", "c", false, "Output with color (only valid for text mode)")
	rootCmd.PersistentFlags().BoolVarP(&onlyDiffs, "onlyDiffs", "d", false, "Output only diffs (only valid for text mode)")

	rootCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "Email to use for downloading APKs from Google Play")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "w", "", "Password to use for downloading APKs from Google Play")
}
