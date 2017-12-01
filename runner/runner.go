package runner

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/openpixel/rise/template"
	"github.com/openpixel/rise/variables"
)

// Run will run
func Run(inputFile, outputFile *string, varFiles *[]string) error {
	contents, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return err
	}

	vars, err := variables.LoadVariableFiles(*varFiles)
	if err != nil {
		return err
	}

	t, err := template.NewTemplate(vars)
	if err != nil {
		return err
	}

	result, err := t.Render(string(contents))
	if err != nil {
		return err
	}

	if *outputFile != "" {
		err = ioutil.WriteFile(*outputFile, []byte(result.Value.(string)), 0644)
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
