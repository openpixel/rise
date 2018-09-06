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

func TestInterpolationFuncJSONEncode(t *testing.T) {
	testCases := []functionTestCase{
		{
			description: "String encoding",
			text:        `${jsonencode("foo")}`,
			expectation: `"foo"`,
		},
		{
			description: "List encoding",
			text:        `${jsonencode(list("foo", "bar"))}`,
			expectation: `["foo","bar"]`,
		},
		{
			description: "Map encoding",
			text:        `${jsonencode(map("foo", "bar"))}`,
			expectation: `{"foo":"bar"}`,
		},
		{
			description: "Nested encoding",
			text:        `${jsonencode(map("foo", list("this", "that")))}`,
			expectation: `{"foo":["this","that"]}`,
		},
	}

	jsonencodeFunc := testInterpolationFunc(keyFuncs{
		"jsonencode": interpolationFuncJSONEncode,
		"list":       interpolationFuncList,
		"map":        interpolationFuncMap,
	})

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			jsonencodeFunc(t, tc)
		})
	}
}
