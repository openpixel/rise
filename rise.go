package rise

import (
	"io"
	"io/ioutil"
	"os"
)

// Run will run
func Run(inputFile, outputFile *string, varFiles *[]string) error {
	contents, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return err
	}

	t, err := newTemplate(varFiles)
	if err != nil {
		return err
	}

	result, err := t.render(string(contents))
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
