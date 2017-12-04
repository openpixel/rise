package interpolation

import "testing"

func TestInterpolationFuncBase64Encode(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Valid list",
			text:        `${base64encode("foo")}`,
			expectation: "Zm9v",
		},
	}

	base64EncodeFunc := testInterpolationFunc(keyFuncs{"base64encode": interpolationFuncBase64Encode})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			base64EncodeFunc(t, tc)
		})
	}
}

func TestInterpolationFuncBase64Decode(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "Valid list",
			text:        `${base64decode("Zm9v")}`,
			expectation: "foo",
		},
	}

	base64DecodeFunc := testInterpolationFunc(keyFuncs{"base64decode": interpolationFuncBase64Decode})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			base64DecodeFunc(t, tc)
		})
	}
}
