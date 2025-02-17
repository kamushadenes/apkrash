package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var dex2jarPath string

var jarCmd = &cobra.Command{
	Use:        "jar",
	Aliases:    []string{"j"},
	Short:      "Convert APK to JAR using dex2jar",
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"file", "output"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var out string
		if len(args) == 2 {
			out = args[1]
		} else {
			out = args[0] + ".jar"
		}

		fmt.Println("Converting APK to JAR using dex2jar")

		command := exec.Command(dex2jarPath, args[0], "-o", out)
		command.Stdout = os.Stdout
		err := command.Run()
		if err != nil {
			return err
		}

		fmt.Println("Conversion complete, output file: " + out)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(jarCmd)

	jarCmd.PersistentFlags().StringVarP(&dex2jarPath, "dex2jar", "p", "d2j-dex2jar", "Path to dex2jar binary")
}
