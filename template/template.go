package template

import (
	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
	"github.com/openpixel/rise/interpolation"
	"github.com/openpixel/rise/variables"
)

// Template is a container for holding onto the ast Variables
type Template struct {
	vars map[string]ast.Variable
}

// NewTemplate will prepare a template object for use
func NewTemplate(varFiles *[]string) (*Template, error) {
	vars, err := variables.LoadVariableFiles(*varFiles)
	if err != nil {
		return nil, err
	}
	return &Template{
		vars: vars,
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

// Render will parse the provided text and interpolate the known variables/funcs
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
