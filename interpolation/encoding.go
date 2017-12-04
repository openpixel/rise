package interpolation

import (
	"encoding/base64"

	"github.com/hashicorp/hil/ast"
)

// interpolationFuncBase64Encode will return the base64 encoding of the provided string
func interpolationFuncBase64Encode() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString},
		ReturnType: ast.TypeString,
		Callback: func(inputs []interface{}) (interface{}, error) {
			return base64.StdEncoding.EncodeToString([]byte(inputs[0].(string))), nil
		},
	}
}

// interpolationFuncBase64Decode will decode the provided base64 value to a string
func interpolationFuncBase64Decode() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString},
		ReturnType: ast.TypeString,
		Callback: func(inputs []interface{}) (interface{}, error) {
			result, err := base64.StdEncoding.DecodeString(inputs[0].(string))
			if err != nil {
				return nil, err
			}
			return string(result), nil
		},
	}
}
