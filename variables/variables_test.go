package variables

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	testCases := []struct {
		input  string
		result []VariableConfig
		error  bool
	}{
		{
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
	}

	for _, tc := range testCases {
		config, err := parseConfig(tc.input)
		if err != nil != tc.error {
			t.Fatalf("unexpected error: %s", err)
		}
		if !reflect.DeepEqual(config.Variables, tc.result) {
			t.Fatalf("wrong result\ngiven %s\ngot: %#v\nwant: %#v", tc.input, config, tc.result)
		}
	}
}
