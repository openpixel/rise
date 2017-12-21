package interpolation

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/hil/ast"
)

// CoreFunctions are the custom functions for interpolation
var CoreFunctions = map[string]ast.Function{
	"lower":      interpolationFuncLower(),
	"upper":      interpolationFuncUpper(),
	"env":        interpolationFuncEnv(),
	"join":       interpolationFuncJoin(),
	"has":        interpolationFuncHas(),
	"map":        interpolationFuncMap(),
	"keys":       interpolationFuncKeys(),
	"list":       interpolationFuncList(),
	"concat":     interpolationFuncConcat(),
	"replace":    interpolationFuncReplace(),
	"max":        interpolationFuncMax(),
	"min":        interpolationFuncMin(),
	"contains":   interpolationFuncContains(),
	"split":      interpolationFuncSplit(),
	"length":     interpolationFuncLength(),
	"jsonencode": interpolationFuncJSONEncode(),
	"pick":       interpolationFuncPick(),
	"omit":       interpolationFuncOmit(),
	"unique":     interpolationFuncUnique(),
}

// interpolationFuncEnv will extract a variable out of the env
func interpolationFuncEnv() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			varName := inputs[0].(string)
			if varName == "" {
				return "", errors.New("must provide a variable name")
			}
			return os.Getenv(varName), nil
		},
	}
}

// interpolationFuncLength will determine the length of the input
// if the input is a list, it will count the length of the items
// if the input is a map, it will count the keys
// if the input is a string, it will count the characters
func interpolationFuncLength() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeAny},
		ReturnType: ast.TypeInt,
		Callback: func(inputs []interface{}) (interface{}, error) {
			switch input := inputs[0].(type) {
			case string:
				return len(input), nil
			case []ast.Variable:
				return len(input), nil
			case map[string]ast.Variable:
				return len(input), nil
			default:
				return nil, fmt.Errorf("must provide either a list, map or string, got %T", input)
			}
		},
	}
}
