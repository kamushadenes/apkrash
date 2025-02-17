package cmd

import (
	"fmt"
	"github.com/kamushadenes/apkrash/apk"
	"github.com/spf13/cobra"
	"os"
)

var compareCmd = &cobra.Command{
	Use:        "compare",
	Aliases:    []string{"c"},
	Short:      "Compares two APKs or Manifests",
	Args:       cobra.ExactArgs(2),
	ArgAliases: []string{"file1", "file2"},
	RunE: func(cmd *cobra.Command, args []string) error {
		apk1, err := apk.ParseAPKInput(args[0], decompile, email, password)
		if err != nil {
			return err
		}
		defer os.RemoveAll(apk1.TmpDir)

		apk2, err := apk.ParseAPKInput(args[1], decompile, email, password)
		if err != nil {
			return err
		}
		defer os.RemoveAll(apk2.TmpDir)

		comparison := apk1.Compare(apk2)
		output, err := comparison.GetComparison(outputFormat, onlyDiffs, includeFiles)
		if err != nil {
			return err
		}

		fmt.Println(output)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
	compareCmd.PersistentFlags().BoolVarP(&includeFiles, "includeFiles", "f", false, "Include files in the output")
	compareCmd.PersistentFlags().BoolVarP(&decompile, "decompile", "l", false, "Decompile APK using jadx")
}
