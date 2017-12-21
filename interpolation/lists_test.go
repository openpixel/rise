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
			text:        `${list("foo", list("bar"), map("flip", "flop"))}`,
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

func TestInterpolationFuncConcat(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Two lists combine",
			text:        `${concat(foo, bar)}`,
			expectation: []interface{}{"test1", "test3", "test2", []interface{}{"test4"}, map[string]interface{}{"test5": "test6"}},
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
						{
							Type: ast.TypeList,
							Value: []ast.Variable{
								{
									Type:  ast.TypeString,
									Value: "test4",
								},
							},
						},
						{
							Type: ast.TypeMap,
							Value: map[string]ast.Variable{
								"test5": ast.Variable{
									Type:  ast.TypeString,
									Value: "test6",
								},
							},
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

func TestInterpolationFuncUnique(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Extracts unique values",
			text:        `${unique(list("1", "2", "1"))}`,
			expectation: []interface{}{"1", "2"},
		},
		{
			description: "Extract unique complex values",
			text:        `${unique(list(list("1", "2", "3"), list("1", "2"), list("1", "2", "3")))}`,
			expectation: []interface{}{
				[]interface{}{"1", "2", "3"},
				[]interface{}{"1", "2"},
			},
		},
	}

	uniqueTestFunc := testInterpolationFunc(keyFuncs{
		"unique": interpolationFuncUnique,
		"list":   interpolationFuncList,
	})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			uniqueTestFunc(t, tc)
		})
	}
}
