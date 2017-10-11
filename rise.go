package rise

import (
	"io/ioutil"

	"github.com/hashicorp/hil"
)

func buildConfig() *hil.EvalConfig {
	return &hil.EvalConfig{}
}

// Run will run
func Run(inputFile, outputFile *string, varFiles *[]string) error {
	contents, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return err
	}

	tree, err := hil.Parse(string(contents))
	if err != nil {
		return err
	}

	result, err := hil.Eval(tree, buildConfig())
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(*outputFile, []byte(result.Value.(string)), 0644)
	if err != nil {
		return err
	}

	return nil
}
