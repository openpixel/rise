package config

import (
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

func TestLoadConfigFiles(t *testing.T) {
	testCases := []struct {
		description string
		filenames   []string
		result      map[string]ast.Variable
		error       bool
	}{
		{
			"Variable file inheritance",
			[]string{"testdata/var1.hcl", "testdata/var2.hcl"},
			map[string]ast.Variable{
				"i": ast.Variable{
					Value: "6",
					Type:  ast.TypeString,
				},
				"j": ast.Variable{
					Value: "2",
					Type:  ast.TypeString,
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
			res, err := LoadConfigFiles(tc.filenames)
			if err != nil != tc.error {
				t.Fatalf("unexpected error: %s", err)
			}
			if !reflect.DeepEqual(res, tc.result) {
				t.Fatalf("wrong result\ngiven %s\ngot: %#v\nwant: %#v", tc.filenames, res, tc.result)
			}
		})
	}
}
