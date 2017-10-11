package main

import (
	"log"

	"github.com/openpixel/rise"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	inputFile  = kingpin.Flag("input", "The template to run substitutions on.").Short('i').Required().String()
	outputFile = kingpin.Flag("output", "The location to output the final contents to.").Short('o').String()
	varFiles   = kingpin.Flag("var-file", "A variable file to load vars from").Strings()
)

func main() {
	kingpin.Parse()

	err := rise.Run(inputFile, outputFile, varFiles)
	if err != nil {
		log.Fatal(err)
	}
}
