package template

import (
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

// Render will parse the provided text and interpolate the known variables/functions
func (t *Template) Render(text string) (hil.EvaluationResult, error) {
	var resultErr error

	config := t.buildConfig()

	tree, err := hil.Parse(text)
	if err != nil {
		return hil.InvalidResult, err
	}

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
			resultN, err := processTemplateNode(n, t.templates[varName], config)
			if err != nil {
				resultErr = err
				return n
			}
			return resultN
		}

		return n
	}

	tree.Accept(fn)

	if resultErr != nil {
		return hil.InvalidResult, resultErr
	}

	result, err := hil.Eval(tree, config)
	if err != nil {
		return hil.InvalidResult, err
	}

	return result, nil
}

func processTemplateNode(original ast.Node, template ast.Variable, config *hil.EvalConfig) (new ast.Node, err error) {
	switch template.Type {
	case ast.TypeString:
		var root ast.Node
		var newValue hil.EvaluationResult
		value := template.Value.(string)
		root, err = hil.Parse(value)
		if err != nil {
			break
		}
		newValue, err = hil.Eval(root, config)
		if err != nil {
			break
		}
		new, err = hil.ParseWithPosition(newValue.Value.(string), original.Pos())
		if err != nil {
			break
		}
		return new, nil
	}

	return original, err
}
