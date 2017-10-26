package interpolation

import (
	"errors"
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
		ArgTypes:     []ast.Type{ast.TypeString},
		ReturnType:   ast.TypeString,
		Variadic:     true,
		VariadicType: ast.TypeList,
		Callback: func(inputs []interface{}) (interface{}, error) {
			if len(inputs) < 2 {
				return nil, errors.New("must have 2 arguments to join")
			}
			var list []string

			for _, arg := range inputs[1:] {
				for _, part := range arg.([]ast.Variable) {
					if part.Type != ast.TypeString {
						return nil, fmt.Errorf(
							"only works on string lists, this list contains elements of %s",
							part.Type.Printable(),
						)
					}
					list = append(list, part.Value.(string))
				}
			}

			return strings.Join(list, inputs[0].(string)), nil
		},
	}
}
