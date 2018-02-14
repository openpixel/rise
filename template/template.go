package template

import (
	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
	"github.com/openpixel/rise/config"
	"github.com/openpixel/rise/interpolation"
)

// Template is a container for holding onto the ast Variables
type Template struct {
	vars      map[string]ast.Variable
	templates map[string]string
}

// NewTemplate will prepare a template object for use
func NewTemplate(configResult *config.Result) (*Template, error) {
	return &Template{
		vars:      configResult.Variables,
		templates: configResult.Templates,
	}, nil
}

func (t *Template) buildConfig() *hil.EvalConfig {
	return &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			FuncMap: interpolation.CoreFunctions,
			VarMap:  t.vars,
		},
	}
}

// Render will parse the provided text and interpolate the known variables/functions
func (t *Template) Render(text string) (hil.EvaluationResult, error) {
	tree, err := hil.Parse(text)
	if err != nil {
		return hil.InvalidResult, err
	}

	result, err := hil.Eval(tree, t.buildConfig())
	if err != nil {
		return hil.InvalidResult, err
	}

	return result, nil
}
