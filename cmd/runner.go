package cmd

import (
	"io"
	"io/ioutil"
	"os"


	"github.com/openpixel/rise/internal/config"
	"github.com/openpixel/rise/internal/template"
)

// Run accepts an input, output and config files and performs interpolation.
// If the output is empty, it writes to stdout
func Run(inputFile, outputFile string, configFiles []string, extraVars []string) error {
	contents, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	extras, err := config.LoadExtras(extraVars)
	if err != nil {
		return err
	}

	configResult, err := config.LoadConfigFiles(configFiles, extras)
	if err != nil {
		return err
	}

	t, err := template.NewTemplate(configResult)
	if err != nil {
		return err
	}

	result, err := t.Render(string(contents))
	if err != nil {
		return err
	}

	if outputFile != "" {
		err = ioutil.WriteFile(outputFile, []byte(result.Value.(string)), 0644)
		if err != nil {
			return err
		}
	} else {
		_, err = io.WriteString(os.Stdout, result.Value.(string))
		if err != nil {
			return err
		}
	}

	return nil
}
