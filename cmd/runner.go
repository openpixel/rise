package cmd

import (
	"io"
	"os"

	"github.com/openpixel/rise/internal/config"
	"github.com/openpixel/rise/internal/template"
)

func prepareTemplate(configFiles []string, extraVars []string) (*template.Template, error) {
	extras, err := config.LoadExtras(extraVars)
	if err != nil {
		return nil, err
	}
	configResult, err := config.LoadConfigFiles(configFiles, extras)
	if err != nil {
		return nil, err
	}
	return template.NewTemplate(configResult)
}

// process accepts an input, output and config files and performs interpolation.
// If the output is empty, it writes to stdout
func process(inputFile, outputFile string, configFiles []string, extraVars []string) error {
	tmpl, err := prepareTemplate(configFiles, extraVars)
	if err != nil {
		return err
	}

	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}

	res, err := tmpl.Render(f)
	if err != nil {
		return err
	}

	var out io.Writer
	if outputFile != "" {
		out, err = os.Create(outputFile)
		if err != nil {
			return err
		}
	} else {
		out = os.Stdout
	}

	_, err = io.Copy(out, res)
	return err
}
