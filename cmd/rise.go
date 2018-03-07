package cmd

import (
	"fmt"
	"log"

	"github.com/openpixel/rise/runner"
	"github.com/spf13/cobra"
)

const version = "v0.0.6"

var inputs string
var outputs string
var configFiles []string

func init() {
	RootCmd.PersistentFlags().StringVarP(&inputs, "input", "i", "", "The file to perform interpolation on")
	RootCmd.PersistentFlags().StringVarP(&outputs, "output", "o", "", "The file to output")
	RootCmd.PersistentFlags().StringSliceVarP(&configFiles, "config", "c", []string{}, "The files that define the configuration to use for interpolation")
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
		err := runner.Run(&inputs, &outputs, &configFiles)
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
