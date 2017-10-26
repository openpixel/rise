package interpolation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hil/ast"
)

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
