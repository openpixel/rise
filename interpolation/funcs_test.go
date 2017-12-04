package interpolation

import (
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
)

type functionTestCase struct {
	description string
	setup       func() error
	text        string
	expectation interface{}
	parseError  bool
	evalError   bool
	teardown    func() error
	vars        map[string]ast.Variable
}

type keyFuncs map[string]func() ast.Function

func testInterpolationFunc(funcMap keyFuncs) func(t *testing.T, testCase functionTestCase) {
	configFuncs := make(map[string]ast.Function)
	for key, interpolationFunc := range funcMap {
		configFuncs[key] = interpolationFunc()
	}
	config := &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			FuncMap: configFuncs,
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

		if _, ok := testCase.expectation.(string); ok {
			if actual.Value.(string) != testCase.expectation.(string) {
				t.Fatalf("wrong result\ngiven %s\ngot: %#v\nwant: %#v", testCase.text, actual.Value, testCase.expectation)
			}
		} else {
			if !reflect.DeepEqual(actual.Value, testCase.expectation) {
				t.Fatalf("wrong result\ngiven %s\ngot: %#v\nwant: %#v", testCase.text, actual.Value, testCase.expectation)
			}
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
		{
			description: "Empty env ternary",
			text:        `${env("FOO") != "" ? env("FOO") : "bar"}`,
			expectation: "bar",
		},
	}

	envTestFunc := testInterpolationFunc(keyFuncs{"env": interpolationFuncEnv})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			envTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncLength(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Length of string",
			text:        `${length("FOO")}`,
			expectation: "3",
		},
		{
			description: "Length of list",
			text:        `${length("${list("foo", "bar")}")}`,
			expectation: "2",
		},
		{
			description: "Length of map",
			text:        `${length("${map("foo", "bar")}")}`,
			expectation: "1",
		},
		{
			description: "Invalid type errors",
			text:        `${length(4)}`,
			evalError:   true,
		},
	}

	lengthTestFunc := testInterpolationFunc(keyFuncs{
		"length": interpolationFuncLength,
		"list":   interpolationFuncList,
		"map":    interpolationFuncMap,
	})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			lengthTestFunc(t, tc)
		})
	}
}
