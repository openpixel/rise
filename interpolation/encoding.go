package interpolation

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/hil"

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

// interpolationFuncJSONEncode will encode an arbitrary into its JSON representation
func interpolationFuncJSONEncode() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeAny},
		ReturnType: ast.TypeString,
		Callback: func(inputs []interface{}) (interface{}, error) {
			input := inputs[0]

			var encodeVal interface{}

			switch in := input.(type) {
			case []ast.Variable:
				//convert to slice
				inStrings := make([]string, len(in))

				for i, v := range in {
					if v.Type != ast.TypeString {
						variable, _ := hil.InterfaceToVariable(in)
						encodeVal, _ = hil.VariableToInterface(variable)

						jEnc, err := json.Marshal(encodeVal)
						if err != nil {
							return "", fmt.Errorf("failed to encode JSON data '%s'", encodeVal)
						}
						return string(jEnc), nil
					}
					inStrings[i] = v.Value.(string)
				}
				encodeVal = inStrings
			case map[string]ast.Variable:
				//convert to map
				mapStrings := make(map[string]string)
				for k, v := range in {
					if v.Type != ast.TypeString {
						variable, _ := hil.InterfaceToVariable(in)
						encodeVal, _ = hil.VariableToInterface(variable)

						jEnc, err := json.Marshal(encodeVal)
						if err != nil {
							return "", fmt.Errorf("failed to encode JSON data '%s'", encodeVal)
						}
						return string(jEnc), nil
					}
					mapStrings[k] = v.Value.(string)
				}
				encodeVal = mapStrings
			case string:
				encodeVal = in
			default:
				return nil, fmt.Errorf("unknown type for JSON encoding: %T", input)
			}
			b, err := json.Marshal(encodeVal)
			if err != nil {
				return nil, err
			}
			return string(b), nil
		},
	}
}
