package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var apktoolPath string

var extractCmd = &cobra.Command{
	Use:        "extract",
	Aliases:    []string{"e", "x"},
	Short:      "Extract APK using apktool",
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"file", "output"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var out string
		if len(args) == 2 {
			out = args[1]
		} else {
			out = strings.TrimSuffix(args[0], ".apk")
		}

		fmt.Println("Extracting APK using apktool")

		command := exec.Command(apktoolPath, "d", args[0], "-o", out)
		command.Stdout = os.Stdout
		err := command.Run()
		if err != nil {
			return err
		}

		fmt.Println("Extraction complete, output directory: " + out)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	extractCmd.PersistentFlags().StringVarP(&apktoolPath, "apktoolPath", "p", "apktool", "Path to apktool binary")
}
