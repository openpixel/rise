package interpolation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hil/ast"
)

// interpolationFuncLower converts a string to be all in lowercase
func interpolationFuncLower() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			return strings.ToLower(inputs[0].(string)), nil
		},
	}
}

// interpolationFuncUpper converts a string to be all in uppercase
func interpolationFuncUpper() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			return strings.ToUpper(inputs[0].(string)), nil
		},
	}
}

// interpolationFuncJoin will join together a list of values with the provided separator
func interpolationFuncJoin() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString, ast.TypeList},
		ReturnType: ast.TypeString,
		Callback: func(inputs []interface{}) (interface{}, error) {
			var list []string

			for _, arg := range inputs[1].([]ast.Variable) {
				if arg.Type != ast.TypeString {
					return nil, fmt.Errorf(
						"only works on string lists, this list contains elements of %s",
						arg.Type.Printable(),
					)
				}
				list = append(list, arg.Value.(string))
			}

			return strings.Join(list, inputs[0].(string)), nil
		},
	}
}
