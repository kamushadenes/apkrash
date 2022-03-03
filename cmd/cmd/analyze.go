package cmd

import (
	"fmt"
	"github.com/kamushadenes/apkrash/apk"
	"github.com/spf13/cobra"
	"os"
)

var analyzeCmd = &cobra.Command{
	Use:        "analyze",
	Aliases:    []string{"a"},
	Short:      "Analyze an APK or Manifest",
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"file"},
	RunE: func(cmd *cobra.Command, args []string) error {
		apk, err := apk.ParseAPKInput(args[0], decompile, email, password)
		if err != nil {
			return err
		}
		defer os.RemoveAll(apk.TmpDir)

		output, err := apk.GetAnalysis(outputFormat, includeFiles)
		if err != nil {
			return err
		}

		fmt.Println(output)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.PersistentFlags().BoolVarP(&includeFiles, "includeFiles", "f", false, "Include files in the output")
	analyzeCmd.PersistentFlags().BoolVarP(&decompile, "decompile", "l", false, "Decompile APK using jadx")
}
