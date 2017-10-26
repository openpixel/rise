package interpolation

import (
	"os"
	"testing"

	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
)

type functionTestCase struct {
	description string
	setup       func() error
	text        string
	expectation string
	parseError  bool
	evalError   bool
	teardown    func() error
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
		if testCase.setup != nil {
			err := testCase.setup()
			if err != nil {
				t.Fatalf("Unexpected error: %s\n", err)
			}
		}
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

		if testCase.teardown != nil {
			err = testCase.teardown()
			if err != nil {
				t.Fatalf("Unexpected error: %s\n", err)
			}
		}
	}
}

func TestInterpolationFuncEnv(t *testing.T) {
	var old string
	testCases := []functionTestCase{
		{
			description: "Existing env works",
			setup: func() error {
				old = os.Getenv("FOO")
				os.Setenv("FOO", "Bar")
				return nil
			},
			text:        `${env("FOO")}`,
			expectation: "Bar",
			teardown: func() error {
				os.Setenv("FOO", old)
				return nil
			},
		},
		{
			description: "Empty env fails",
			setup: func() error {
				old = os.Getenv("FOO")
				os.Setenv("FOO", "")
				return nil
			},
			text:        `${env("FOO")}`,
			expectation: "",
			teardown: func() error {
				os.Setenv("FOO", old)
				return nil
			},
		},
		{
			description: "Empty argument fails",
			text:        `${env("")}`,
			expectation: "",
			evalError:   true,
		},
	}

	envTestFunc := testInterpolationFunc("env", interpolationFuncEnv)

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			envTestFunc(t, tc)
		})
	}
}
