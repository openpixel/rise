package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const version = "v0.2.0"

var inputs string
var outputs string
var configFiles []string
var extraVars []string

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVarP(&inputs, "input", "i", "", "The file to perform interpolation on")
	flags.StringVarP(&outputs, "output", "o", "", "The file to output")
	flags.StringSliceVarP(&configFiles, "config", "c", []string{}, "The files that define the configuration to use for interpolation")
	flags.StringArrayVar(&extraVars, "var", []string{}, "Additional variables to apply. These always take priority.")

	RootCmd.AddCommand(versionCmd)
}

// RootCmd is the root command for the entire cli
var RootCmd = &cobra.Command{
	Use:   "rise",
	Short: "Rise is a powerful text interpolation tool.",
	Long:  `A powerful text interpolation tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inputs == "" {
			log.Fatal("Must have an input")
		}
		err := process(inputs, outputs, configFiles, extraVars)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of rise",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
