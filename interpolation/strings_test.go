package interpolation

import (
	"testing"
)

func TestInterpolationFuncLower(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Uppercase becomes lowercase",
			text:        `${lower("FOO")}`,
			expectation: "foo",
		},
		{
			description: "Lowercase stays lowercase",
			text:        `${lower("foo")}`,
			expectation: "foo",
		},
	}

	lowerTestFunc := testInterpolationFunc(keyFuncs{"lower": interpolationFuncLower})

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
		},
		{
			description: "Lowercase becomes uppercase",
			text:        `${upper("foo")}`,
			expectation: "FOO",
		},
	}

	lowerTestFunc := testInterpolationFunc(keyFuncs{"upper": interpolationFuncUpper})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			lowerTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncReplace(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Replace N occurrences",
			text:        `${replace("foo bar bar", " bar", "", -1)}`,
			expectation: "foo",
		},
		{
			description: "Replace 0 occurrences",
			text:        `${replace("foo bar bar", " bar", "", 0)}`,
			expectation: "foo bar bar",
		},
		{
			description: "Replace 1 occurrences",
			text:        `${replace("foo bar bar", " bar", "", 1)}`,
			expectation: "foo bar",
		},
	}

	replaceTestFunc := testInterpolationFunc(keyFuncs{"replace": interpolationFuncReplace})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			replaceTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncContains(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "String does contain portion",
			text:        `${contains("hello", "ell")}`,
			expectation: "true",
		},
		{
			description: "String does not contain portion",
			text:        `${contains("hello", "foo")}`,
			expectation: "false",
		},
	}

	containsTestFunc := testInterpolationFunc(keyFuncs{"contains": interpolationFuncContains})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			containsTestFunc(t, tc)
		})
	}
}

func TestInterpolationFuncSplit(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Splits into list",
			text:        `${split("foo,bar", ",")}`,
			expectation: []interface{}{"foo", "bar"},
		},
		{
			description: "Splits even when sep not found",
			text:        `${split("foo,bar", ":")}`,
			expectation: []interface{}{"foo,bar"},
		},
	}

	splitTestFunc := testInterpolationFunc(keyFuncs{"split": interpolationFuncSplit})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			splitTestFunc(t, tc)
		})
	}
}
