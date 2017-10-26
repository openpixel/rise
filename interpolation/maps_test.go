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
			expectation: "false",
			vars: map[string]ast.Variable{
				"foo": ast.Variable{
					Type: ast.TypeMap,
					Value: map[string]ast.Variable{
						"foo": ast.Variable{
							Type:  ast.TypeString,
							Value: "Bar",
						},
					},
				},
			},
		},
	}

	hasTestFunc := testInterpolationFunc("has", interpolationFuncHas)

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			hasTestFunc(t, tc)
		})
	}
}
