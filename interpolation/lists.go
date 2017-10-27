package interpolation

import "github.com/hashicorp/hil/ast"
import "github.com/hashicorp/hil"

// interpolationFuncList will accept a variable number of arguments and create a list
func interpolationFuncList() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{},
		ReturnType:   ast.TypeList,
		Variadic:     true,
		VariadicType: ast.TypeAny,
		Callback: func(inputs []interface{}) (interface{}, error) {
			result := make([]ast.Variable, 0, len(inputs))

			for _, input := range inputs {
				nativeVar, err := hil.InterfaceToVariable(input)
				if err != nil {
					return nil, err
				}
				result = append(result, nativeVar)
			}

			return result, nil
		},
	}
}
