package interpolation

import "testing"

func TestInterpolationFuncList(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Valid list",
			text:        `${list("foo", "bar")}`,
			expectation: []interface{}{"foo", "bar"},
		},
		{
			description: "Multi-type list",
			text:        `${list("foo", "${list("bar")}")}`,
			expectation: []interface{}{"foo", []interface{}{"bar"}},
		},
		{
			description: "Empty list",
			text:        `${list()}`,
			expectation: []interface{}{},
		},
	}

	listTestFunc := testInterpolationFunc("list", interpolationFuncList)

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			listTestFunc(t, tc)
		})
	}
}
