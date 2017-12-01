package interpolation

import (
	"testing"

	"github.com/hashicorp/hil/ast"
)

func TestInterpolationFuncList(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Valid list",
			text:        `${list("foo", "bar")}`,
			expectation: []interface{}{"foo", "bar"},
		},
		{
			description: "Multi-type list",
			text:        `${list("foo", "${list("bar")}", "${map("flip", "flop")}")}`,
			expectation: []interface{}{"foo", []interface{}{"bar"}, map[string]interface{}{"flip": "flop"}},
		},
		{
			description: "Empty list",
			text:        `${list()}`,
			expectation: []interface{}{},
		},
	}

	listTestFunc := testInterpolationFunc(keyFuncs{"list": interpolationFuncList, "map": interpolationFuncMap})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			listTestFunc(t, tc)
		})
	}
}

func TestInterpoloationFuncConcat(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Two lists combine",
			text:        `${concat(foo, bar)}`,
			expectation: []interface{}{"test1", "test3", "test2"},
			vars: map[string]ast.Variable{
				"foo": ast.Variable{
					Type: ast.TypeList,
					Value: []ast.Variable{
						{
							Type:  ast.TypeString,
							Value: "test1",
						},
						{
							Type:  ast.TypeString,
							Value: "test3",
						},
					},
				},
				"bar": ast.Variable{
					Type: ast.TypeList,
					Value: []ast.Variable{
						{
							Type:  ast.TypeString,
							Value: "test2",
						},
					},
				},
			},
		},
	}

	concatTestFunc := testInterpolationFunc(keyFuncs{"concat": interpolationFuncConcat})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			concatTestFunc(t, tc)
		})
	}
}
