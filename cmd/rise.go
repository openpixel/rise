package cmd

import (
	"log"

	"github.com/openpixel/rise/runner"
	"github.com/spf13/cobra"
)

var inputs string
var outputs string
var varFiles []string

func init() {
	RootCmd.PersistentFlags().StringVarP(&inputs, "input", "i", "", "The file to perform substitution on")
	RootCmd.PersistentFlags().StringVarP(&outputs, "output", "o", "", "The file to output")
	RootCmd.PersistentFlags().StringSliceVar(&varFiles, "varFile", []string{}, "The files that contains the variables to be substituted")
}

// RootCmd is the root command for the entire cli
var RootCmd = &cobra.Command{
	Use:   "rise",
	Short: "Rise is an awesome template substitution tool.",
	Long:  `A simple template substitution tool.`,
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