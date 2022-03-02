package cmd

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/kamushadenes/apkrash/apk"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:        "analyze",
	Aliases:    []string{"a"},
	Short:      "Analyze an APK or Manifest",
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"file"},
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}
		mtype := mimetype.Detect(file)

		var apk apk.APK

		switch strings.Split(mtype.String(), ";")[0] {
		case "text/xml":
			err = apk.ParseManifest(file)
			if err != nil {
				return err
			}
		case "application/zip", "application/jar":
			apk.Filename = args[0]
			err = apk.ParseZIP()
			if err != nil {
				return err
			}
			if decompile {
				err = apk.Decompile()
				if err != nil {
					return err
				}
				err = apk.ParseSources()
				if err != nil {
					return err
				}
			}
		}

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
