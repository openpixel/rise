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
		result      []variableConfig
		error       bool
	}{
		{
			"String parse",
			`variable "foo" { value = "bar"}`,
			[]variableConfig{
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
			[]variableConfig{
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
			[]variableConfig{
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
			[]variableConfig{
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

func TestPrepareVariables(t *testing.T) {
	testCases := []struct {
		description string
		config      []variableConfig
		result      map[string]ast.Variable
		error       bool
	}{
		{
			"Valid interpolation",
			[]variableConfig{
				{
					Name:  "foo",
					Value: `${lower("BAR")}`,
				},
			},
			map[string]ast.Variable{
				"var.foo": ast.Variable{
					Value: "${lower(\"BAR\")}",
					Type:  ast.TypeString,
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			vars := make(map[string]ast.Variable)
			config := &config{
				Variables: tc.config,
			}
			err := prepareVariables(vars, config)
			if err != nil != tc.error {
				t.Fatalf("unexpected error: %s", err)
			}
			if !reflect.DeepEqual(vars, tc.result) {
				t.Fatalf("wrong result\ngiven %#v\ngot: %#v\nwant: %#v", config, vars, tc.result)
			}
		})
	}
}

func TestLoadExtras(t *testing.T) {
	testCases := []struct {
		desc string
		args []string
		expectation map[string]ast.Variable
		error bool
	}{
		{
			"successfully load vars",
			[]string{`{"i": "foo"}`},
			map[string]ast.Variable{
				"var.i": {
					Value: "foo",
					Type: ast.TypeString,
				},
			},
			false,
		},
		{
			"invalid json",
			[]string{"s"},
			nil,
			true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := LoadExtras(tC.args)
			if err != nil != tC.error {
				t.Fatalf("Unexpected error: %s", err)
			}
			if !reflect.DeepEqual(result, tC.expectation) {
				t.Fatalf("Wrong results:\nGot: %#v\nWanted: %#v\n", result, tC.expectation)
			}
		})
	}
}

func TestLoadConfigFiles(t *testing.T) {
	testCases := []struct {
		description string
		filenames   []string
		extras		map[string]ast.Variable
		result      *Result
		error       bool
	}{
		{
			"config file inheritance",
			[]string{"testdata/var1.hcl", "testdata/var2.hcl"},
			nil,
			&Result{
				Variables: map[string]ast.Variable{
					"var.i": ast.Variable{
						Value: "6",
						Type:  ast.TypeString,
					},
					"var.j": ast.Variable{
						Value: "2",
						Type:  ast.TypeString,
					},
				},
				Templates: map[string]ast.Variable{
					"tmpl.basic": ast.Variable{
						Value: "this is a template",
						Type:  ast.TypeString,
					},
				},
			},
			false,
		},
		{
			"config with file template",
			[]string{"testdata/var3.hcl"},
			nil,
			&Result{
				Variables: map[string]ast.Variable{
					"var.j": ast.Variable{
						Value: "world",
						Type:  ast.TypeString,
					},
				},
				Templates: map[string]ast.Variable{
					"tmpl.advanced": ast.Variable{
						Value: "hello, ${var.j}",
						Type:  ast.TypeString,
					},
				},
			},
			false,
		},
		{
			"config with extras",
			[]string{"testdata/var3.hcl"},
			map[string]ast.Variable{
				"var.h": {
					Value: "foo",
					Type: ast.TypeString,
				},
			},
			&Result{
				Variables: map[string]ast.Variable{
					"var.j": ast.Variable{
						Value: "world",
						Type:  ast.TypeString,
					},
					"var.h": {
						Value: "foo",
						Type: ast.TypeString,
					},
				},
				Templates: map[string]ast.Variable{
					"tmpl.advanced": ast.Variable{
						Value: "hello, ${var.j}",
						Type:  ast.TypeString,
					},
				},
			},
			false,
		},
		{
			"Bad file should error",
			[]string{"bad"},
			nil,
			nil,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			res, err := LoadConfigFiles(tc.filenames, tc.extras)
			if err != nil != tc.error {
				t.Fatalf("unexpected error: %s", err)
			}
			if !reflect.DeepEqual(res, tc.result) {
				t.Fatalf("wrong result\ngiven %s\ngot: %#v\nwant: %#v", tc.filenames, res, tc.result)
			}
		})
	}
}
