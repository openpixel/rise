package cmd

import (
	"log"
	"fmt"

	"github.com/openpixel/rise/runner"
	"github.com/spf13/cobra"
)

const version = "v0.0.6"

var inputs string
var outputs string
var varFiles []string

func init() {
	RootCmd.PersistentFlags().StringVarP(&inputs, "input", "i", "", "The file to perform interpolation on")
	RootCmd.PersistentFlags().StringVarP(&outputs, "output", "o", "", "The file to output")
	RootCmd.PersistentFlags().StringSliceVarP(&varFiles, "vars", "V", []string{}, "The files that contains the variables to be interpolated")
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
		err := runner.Run(&inputs, &outputs, &varFiles)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of rise",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
