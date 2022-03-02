package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var jadxPath string

var decompileCmd = &cobra.Command{
	Use:        "decompile",
	Aliases:    []string{"d"},
	Short:      "Decompile APK into Java code using jadx",
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"file", "output"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var out string
		if len(args) == 2 {
			out = args[1]
		} else {
			out = strings.TrimSuffix(args[0], ".apk")
		}

		fmt.Println("Decompiling APK using jadx")

		command := exec.Command("jadx", "-d", out, "--deobf", args[0])
		command.Stdout = os.Stdout
		err := command.Run()
		if err != nil {
			return err
		}

		fmt.Println("Decompiling complete, output directory: " + out)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(decompileCmd)
	decompileCmd.PersistentFlags().StringVarP(&jadxPath, "jadxPath", "p", "jadx", "Path to jadx binary")
}
