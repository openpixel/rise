package interpolation

import "testing"

func TestInterpolationFuncLower(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Uppercase becomes lowercase",
			text:        `${lower("FOO")}`,
			expectation: "foo",
			error:       false,
		},
		{
			description: "Lowercase stays lowercase",
			text:        `${lower("foo")}`,
			expectation: "foo",
			error:       false,
		},
	}

	lowerTestFunc := testInterpolationFunc("lower", interpolationFuncLower)

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			lowerTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncUpper(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Uppercase stays uppercase",
			text:        `${upper("FOO")}`,
			expectation: "FOO",
			error:       false,
		},
		{
			description: "Lowercase becomes uppercase",
			text:        `${upper("foo")}`,
			expectation: "FOO",
			error:       false,
		},
	}

	lowerTestFunc := testInterpolationFunc("upper", interpolationFuncUpper)

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			lowerTestFunc(t, tc)
		})
	}
}
