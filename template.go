package rise

import (
	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
	"github.com/openpixel/rise/interpolation"
	"github.com/openpixel/rise/variables"
)

type template struct {
	vars map[string]ast.Variable
}

func newTemplate(varFiles *[]string) (*template, error) {
	vars, err := variables.LoadVariableFiles(*varFiles)
	if err != nil {
		return nil, err
	}
	return &template{
		vars: vars,
	}, nil
}

func (t *template) buildConfig() *hil.EvalConfig {
	return &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			FuncMap: interpolation.CoreFunctions,
			VarMap:  t.vars,
		},
	}
}

func (t *template) render(text string) (hil.EvaluationResult, error) {
	tree, err := hil.Parse(text)
	if err != nil {
		return hil.InvalidResult, err
	}

	if err != nil {
		return hil.InvalidResult, err
	}
	result, err := hil.Eval(tree, t.buildConfig())
	if err != nil {
		return hil.InvalidResult, err
	}

	return result, nil
}
