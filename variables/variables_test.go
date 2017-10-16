package variables

import (
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/hil/ast"
)

func TestParseConfig(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		result      []VariableConfig
		error       bool
	}{
		{
			"String parse",
			`variable "foo" { value = "bar"}`,
			[]VariableConfig{
				{
					Name:  "foo",
					Value: "bar",
				},
			},
			false,
		},
		{
			"List parse",
			`variable "foo" { value = ["bar"]}`,
			[]VariableConfig{
				{
					Name:  "foo",
					Value: []interface{}{"bar"},
				},
			},
			false,
		},
		{
			"Map parse",
			`variable "foo" { value = { "bar" = "zoo" } }`,
			[]VariableConfig{
				{
					Name: "foo",
					Value: []map[string]interface{}{
						map[string]interface{}{
							"bar": "zoo",
						},
					},
				},
			},
			false,
		},
		{
			"Int parse",
			`variable "foo" { value = 6}`,
			[]VariableConfig{
				{
					Name:  "foo",
					Value: 6,
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			config, err := parseConfig(tc.input)
			if err != nil != tc.error {
				t.Fatalf("unexpected error: %s", err)
			}
			if !reflect.DeepEqual(config.Variables, tc.result) {
				t.Fatalf("wrong result\ngiven %s\ngot: %#v\nwant: %#v", tc.input, config, tc.result)
			}
		})
	}
}

func TestInterpolateVariables(t *testing.T) {
	testCases := []struct {
		description string
		config      []VariableConfig
		result      map[string]ast.Variable
		error       bool
	}{
		{
			"Valid interpolation",
			[]VariableConfig{
				{
					Name:  "foo",
					Value: `${lower("BAR")}`,
				},
			},
			map[string]ast.Variable{
				"foo": ast.Variable{
					Value: "bar",
					Type:  ast.TypeString,
				},
			},
			false,
		},
		{
			"Invalid interpolation should error",
			[]VariableConfig{
				{
					Name:  "foo",
					Value: `${lower(a, b)}`,
				},
			},
			map[string]ast.Variable{},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			vars := make(map[string]ast.Variable)
			config := &Config{
				Variables: tc.config,
			}
			err := interpolateVariables(vars, config)
			if err != nil != tc.error {
				t.Fatalf("unexpected error: %s", err)
			}
			if !reflect.DeepEqual(vars, tc.result) {
				t.Fatalf("wrong result\ngiven %s\ngot: %#v\nwant: %#v", config, vars, tc.result)
			}
		})
	}
}

func TestLoadVariableFiles(t *testing.T) {
	jValOrig := os.Getenv("J_VAL")
	defer func() {
		err := os.Setenv("J_VAL", jValOrig)
		if err != nil {
			t.Fatal(err)
		}
	}()
	os.Setenv("J_VAL", "2")
	testCases := []struct {
		description string
		filenames   []string
		result      map[string]ast.Variable
		error       bool
	}{
		{
			"Variable file inheritance",
			[]string{"../examples/vars.hcl", "../examples/vars2.hcl"},
			map[string]ast.Variable{
				"i": ast.Variable{
					Value: "10",
					Type:  ast.TypeString,
				},
				"j": ast.Variable{
					Value: "2",
					Type:  ast.TypeString,
				},
				"k": ast.Variable{
					Value: map[string]ast.Variable{
						"t": ast.Variable{
							Value: "z",
							Type:  ast.TypeString,
						},
					},
					Type: ast.TypeMap,
				},
				"h": ast.Variable{
					Value: []ast.Variable{
						{
							Value: "Foo",
							Type:  ast.TypeString,
						},
					},
					Type: ast.TypeList,
				},
			},
			false,
		},
		{
			"Bad file should error",
			[]string{"bad"},
			map[string]ast.Variable(nil),
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			res, err := LoadVariableFiles(tc.filenames)
			if err != nil != tc.error {
				t.Fatalf("unexpected error: %s", err)
			}
			if !reflect.DeepEqual(res, tc.result) {
				t.Fatalf("wrong result\ngiven %s\ngot: %#v\nwant: %#v", tc.filenames, res, tc.result)
			}
		})
	}
}
