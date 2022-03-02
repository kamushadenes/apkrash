package cmd

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/kamushadenes/apkrash/apk"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var compareCmd = &cobra.Command{
	Use:        "compare",
	Aliases:    []string{"c"},
	Short:      "Compares two APKs or Manifests",
	Args:       cobra.ExactArgs(2),
	ArgAliases: []string{"file1", "file2"},
	RunE: func(cmd *cobra.Command, args []string) error {
		file1, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}
		file2, err := os.ReadFile(args[1])
		if err != nil {
			return err
		}
		mtype1 := mimetype.Detect(file1)
		mtype2 := mimetype.Detect(file2)

		var apk1, apk2 apk.APK

		switch strings.Split(mtype1.String(), ";")[0] {
		case "text/xml":
			err = apk1.ParseManifest(file1)
			if err != nil {
				return err
			}
		case "application/zip", "application/jar":
			apk1.Filename = args[0]
			err = apk1.ParseZIP()
			if err != nil {
				return err
			}
			if decompile {
				err = apk1.Decompile()
				if err != nil {
					return err
				}
				err = apk1.ParseSources()
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("unsupported file type: %s", mtype1.String())
		}

		switch strings.Split(mtype2.String(), ";")[0] {
		case "text/xml":
			err = apk2.ParseManifest(file2)
			if err != nil {
				return err
			}
		case "application/zip", "application/jar":
			apk2.Filename = args[1]
			err = apk2.ParseZIP()
			if err != nil {
				return err
			}
			if decompile {
				err = apk2.Decompile()
				if err != nil {
					return err
				}
				err = apk2.ParseSources()
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("unsupported file type: %s", mtype2.String())
		}

		comparison := apk1.Compare(&apk2)
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
