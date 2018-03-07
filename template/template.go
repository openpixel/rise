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

	vf := visitorFn{
		config:    config,
		templates: t.templates,
	}
	tree = tree.Accept(vf.fn)
	if vf.resultErr != nil {
		return hil.InvalidResult, resultErr
	}

	result, err := hil.Eval(tree, config)
	if err != nil {
		return hil.InvalidResult, err
	}

	return result, nil
}

type visitorFn struct {
	resultErr error
	config    *hil.EvalConfig
	templates map[string]ast.Variable
}

func (vf visitorFn) fn(n ast.Node) ast.Node {
	if vf.resultErr != nil {
		return n
	}
	switch vn := n.(type) {
	case *ast.VariableAccess:
		// Found a variable reference that needs to be processed
		n = vf.processVariable(vn)
	case *ast.Index:
		// We found an index reference. In this case we need to check
		// both the target and key for VariableAccess usage.
		if va, ok := vn.Target.(*ast.VariableAccess); ok {
			vn.Target = vf.processVariable(va)
		}
		if va, ok := vn.Key.(*ast.VariableAccess); ok {
			vn.Key = vf.processVariable(va)
		}
	default:
		return n
	}
	return n
}

func (vf visitorFn) processVariable(va *ast.VariableAccess) ast.Node {
	name := va.Name
	// When the variable accessor is a template, we need to process it
	if strings.HasPrefix(name, "tmpl.") {
		resultN, err := vf.processTemplateNode(va, name)
		if err != nil {
			vf.resultErr = err
			return va
		}
		resultN = vf.fn(resultN)
		return resultN
	}

	return va
}

func (vf visitorFn) processTemplateNode(original ast.Node, name string) (replacement ast.Node, err error) {
	template := vf.templates[name]

	switch template.Type {
	case ast.TypeString:
		var root ast.Node
		var newValue hil.EvaluationResult
		value := template.Value.(string)
		root, err = hil.Parse(value)
		if err != nil {
			break
		}
		newValue, err = hil.Eval(root, vf.config)
		if err != nil {
			break
		}
		replacement, err = hil.ParseWithPosition(newValue.Value.(string), original.Pos())
		if err != nil {
			break
		}
		return replacement, nil
	}

	return original, err
}
