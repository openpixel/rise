package interpolation

import (
	"github.com/hashicorp/hil/ast"
)

func interpolationFuncHas() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeMap, ast.TypeString},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			mapInput := inputs[0].(map[string]ast.Variable)
			_, ok := mapInput[inputs[1].(string)]
			if ok {
				return "true", nil
			}

			return "false", nil
		},
	}
}
