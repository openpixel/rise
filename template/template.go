package template

import (
	"errors"
	"strings"

	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
	"github.com/openpixel/rise/config"
	"github.com/openpixel/rise/interpolation"
)

// Template is a container for holding onto the ast Variables
type Template struct {
	vars      map[string]ast.Variable
	templates map[string]ast.Variable
}

// NewTemplate will prepare a template object for use
func NewTemplate(configResult *config.Result) (*Template, error) {
	return &Template{
		vars:      configResult.Variables,
		templates: configResult.Templates,
	}, nil
}

func (t *Template) buildConfig() *hil.EvalConfig {
	vars := make(map[string]ast.Variable)

	for k, v := range t.vars {
		vars[k] = v
	}
	for k, v := range t.templates {
		vars[k] = v
	}

	return &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			FuncMap: interpolation.CoreFunctions,
			VarMap:  vars,
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
	err := hil.Walk(&text, func(data *hil.WalkData) error {
		var resultErr error

		fn := func(n ast.Node) ast.Node {
			if resultErr != nil {
				return n
			}

			var varName string

			switch vn := n.(type) {
			case *ast.VariableAccess:
				varName = vn.Name
			case *ast.Index:
				if va, ok := vn.Target.(*ast.VariableAccess); ok {
					varName = va.Name
				}
				if va, ok := vn.Key.(*ast.VariableAccess); ok {
					varName = va.Name
				}
			default:
				return n
			}

			if strings.HasPrefix(varName, "tmpl.") {
				if t.templates[varName].Type == ast.TypeString {
					value := t.templates[varName].Value.(string)
					data.Replace = true
					data.ReplaceValue = value
				}
			}

			return n
		}

		data.Root.Accept(fn)

		if resultErr != nil {
			return resultErr
		}

		return nil
	})
	if err != nil {
		return hil.InvalidResult, err
	}

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
