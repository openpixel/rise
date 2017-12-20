package interpolation

import (
	"testing"

	"github.com/hashicorp/hil/ast"
)

func TestInterpolationFuncHas(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Key exists",
			text:        `${has(foo, "bar")}`,
			expectation: "true",
			vars: map[string]ast.Variable{
				"foo": ast.Variable{
					Type: ast.TypeMap,
					Value: map[string]ast.Variable{
						"bar": ast.Variable{
							Type:  ast.TypeString,
							Value: "Bar",
						},
					},
				},
			},
		},
		{
			description: "Key does not exist",
			text:        `${has(foo, "bar")}`,
			expectation: "false",
			vars: map[string]ast.Variable{
				"foo": ast.Variable{
					Type:  ast.TypeMap,
					Value: map[string]ast.Variable{},
				},
			},
		},
	}

	hasTestFunc := testInterpolationFunc(keyFuncs{"has": interpolationFuncHas})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			hasTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncMap(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Map is created",
			text:        `${map("foo", "bar")}`,
			expectation: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			description: "Map with different types",
			text:        `${map("foo", "bar", "key", map("flip", "flop"))}`,
			expectation: map[string]interface{}{
				"foo": "bar",
				"key": map[string]interface{}{
					"flip": "flop",
				},
			},
		},
		{
			description: "Odd argument count fails",
			text:        `${map("foo")}`,
			evalError:   true,
		},
	}

	mapTestFunc := testInterpolationFunc(keyFuncs{"map": interpolationFuncMap})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mapTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncKeys(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Returns the keys",
			text:        `${keys(foo)}`,
			expectation: []interface{}{"bar", "bar2"},
			vars: map[string]ast.Variable{
				"foo": ast.Variable{
					Type: ast.TypeMap,
					Value: map[string]ast.Variable{
						"bar2": ast.Variable{
							Type:  ast.TypeString,
							Value: "other2",
						},
						"bar": ast.Variable{
							Type:  ast.TypeString,
							Value: "other",
						},
					},
				},
			},
		},
	}

	keysTestFunc := testInterpolationFunc(keyFuncs{"keys": interpolationFuncKeys})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			keysTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncMerge(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Merges two maps",
			text:        `${merge(map("flip", "flop", "foo", "bar"), map("foo", "override"))}`,
			expectation: map[string]interface{}{
				"foo":  "override",
				"flip": "flop",
			},
		},
		{
			description: "Single map returns",
			text:        `${merge(map("foo", "bar"))}`,
			expectation: map[string]interface{}{
				"foo": "bar",
			},
		},
	}

	mergeTestFunc := testInterpolationFunc(keyFuncs{"merge": interpolationFuncMerge, "map": interpolationFuncMap})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mergeTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncPick(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Picks correctly",
			text:        `${pick(map("foo", "bar", "this", "that"), "foo")}`,
			expectation: map[string]interface{}{
				"foo": "bar",
			},
		},
	}

	pickTestFunc := testInterpolationFunc(keyFuncs{
		"pick": interpolationFuncPick,
		"map":  interpolationFuncMap,
	})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			pickTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncOmit(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Omit correctly",
			text:        `${omit(map("foo", "bar", "this", "that"), "foo")}`,
			expectation: map[string]interface{}{
				"this": "that",
			},
		},
	}

	pickTestFunc := testInterpolationFunc(keyFuncs{
		"omit": interpolationFuncOmit,
		"map":  interpolationFuncMap,
	})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			pickTestFunc(t, tc)
		})
	}
}
