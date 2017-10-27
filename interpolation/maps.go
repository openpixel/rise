package interpolation

import (
	"fmt"

	"github.com/hashicorp/hil"

	"github.com/hashicorp/hil/ast"
)

// interpolationFuncHas returns if the key exists in the provided map
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

// interpolationFuncMap accepts a variable number of arguments in key/value pairs
// and converts them to a map
func interpolationFuncMap() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{},
		ReturnType:   ast.TypeMap,
		Variadic:     true,
		VariadicType: ast.TypeAny,
		Callback: func(inputs []interface{}) (interface{}, error) {
			result := make(map[string]ast.Variable)

			if len(inputs)%2 != 0 {
				return nil, fmt.Errorf("requires an even number of arguments, got %d", len(inputs))
			}

			for i := 0; i < len(inputs); i += 2 {
				key, ok := inputs[i].(string)
				if !ok {
					return nil, fmt.Errorf("argument %d represents a key in the map, but it is not a string", i+1)
				}
				val := inputs[i+1]
				nativeVar, err := hil.InterfaceToVariable(val)
				if err != nil {
					return nil, err
				}
				result[key] = nativeVar
			}

			return result, nil
		},
	}
}
