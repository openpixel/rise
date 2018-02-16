package template

import (
	"errors"
	"fmt"
	"strings"

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

func NewVariable(name string) (interface{}, error) {
	if strings.HasPrefix(name, "var.") {
		return "Var", nil
	} else if strings.HasPrefix(name, "tmpl") {
		return "Template", nil
	}

	return nil, errors.New("Unknown variable access")
}

// Render will parse the provided text and interpolate the known variables/functions
func (t *Template) Render(text string) (hil.EvaluationResult, error) {
	tree, err := hil.Parse(text)
	if err != nil {
		return hil.InvalidResult, err
	}

	hil.Walk(text, func(data *hil.WalkData) error {
		var resultErr error

		fn := func(n ast.Node) ast.Node {
			if resultErr != nil {
				return n
			}

			switch vn := n.(type) {
			case *ast.VariableAccess:
				v, err := NewVariable(vn.Name)
				if err != nil {
					resultErr = err
					return n
				}
				fmt.Println(v)
			case *ast.Index:
				if va, ok := vn.Target.(*ast.VariableAccess); ok {
					v, err := NewVariable(va.Name)
					if err != nil {
						resultErr = err
						return n
					}
					fmt.Println(v)
				}
				if va, ok := vn.Key.(*ast.VariableAccess); ok {
					v, err := NewVariable(va.Name)
					if err != nil {
						resultErr = err
						return n
					}
					fmt.Println(v)
				}
			default:
				return n
			}
			return n
		}

		data.Root.Accept(fn)

		if resultErr != nil {
			return resultErr
		}

		return nil
	})

	result, err := hil.Eval(tree, t.buildConfig())
	if err != nil {
		return hil.InvalidResult, err
	}

	return result, nil
}
