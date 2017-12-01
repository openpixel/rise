package interpolation

import (
	"errors"
	"os"

	"github.com/hashicorp/hil/ast"
)

// CoreFunctions are the custom funtions for interpolation
var CoreFunctions = map[string]ast.Function{
	"lower":   interpolationFuncLower(),
	"upper":   interpolationFuncUpper(),
	"env":     interpolationFuncEnv(),
	"join":    interpolationFuncJoin(),
	"has":     interpolationFuncHas(),
	"map":     interpolationFuncMap(),
	"keys":    interpolationFuncKeys(),
	"list":    interpolationFuncList(),
	"concat":  interpolationFuncConcat(),
	"replace": interpolationFuncReplace(),
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
				return "", errors.New("Must provide a variable name")
			}
			return os.Getenv(varName), nil
		},
	}
}
