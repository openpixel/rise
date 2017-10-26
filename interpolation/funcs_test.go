package interpolation

import (
	"testing"

	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
)

type functionTestCase struct {
	description string
	text        string
	expectation string
	parseError  bool
	evalError   bool
	vars        map[string]ast.Variable
}

func testInterpolationFunc(key string, interpolationFunc func() ast.Function) func(t *testing.T, testCase functionTestCase) {
	config := &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			FuncMap: map[string]ast.Function{
				key: interpolationFunc(),
			},
		},
	}

	return func(t *testing.T, testCase functionTestCase) {
		if testCase.vars == nil {
			config.GlobalScope.VarMap = map[string]ast.Variable{}
		} else {
			config.GlobalScope.VarMap = testCase.vars
		}
		tree, err := hil.Parse(testCase.text)
		if err != nil != testCase.parseError {
			t.Fatalf("Unexpected error: %s\n", err)
		}
		if testCase.parseError {
			return
		}

		actual, err := hil.Eval(tree, config)
		if err != nil != testCase.evalError {
			t.Fatalf("Unexpected error: %s\n", err)
		}
		if testCase.evalError {
			return
		}

		if actual.Value.(string) != testCase.expectation {
			t.Fatalf("wrong result\ngiven %s\ngot: %s\nwant: %s", testCase.text, actual, testCase.expectation)
		}
	}
}

func TestInterpolationFuncEnv(t *testing.T) {

}
