package interpolation

import (
	"fmt"

	"github.com/hashicorp/hil/ast"
)

// interpolationFuncList will accept a variable number of arguments and create a list
func interpolationFuncList() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{},
		ReturnType:   ast.TypeList,
		Variadic:     true,
		VariadicType: ast.TypeAny,
		Callback: func(inputs []interface{}) (interface{}, error) {
			var result []ast.Variable

			for i, input := range inputs {
				switch v := input.(type) {
				case string:
					result = append(result, ast.Variable{Type: ast.TypeString, Value: v})
				case []ast.Variable:
					result = append(result, ast.Variable{Type: ast.TypeList, Value: v})
				case map[string]ast.Variable:
					result = append(result, ast.Variable{Type: ast.TypeMap, Value: v})
				default:
					return nil, fmt.Errorf("unexpected type %T for argument %d in list", v, i)
				}
			}

			return result, nil
		},
	}
}

// interpolationFuncConcat will concat multiple lists into a single list
func interpolationFuncConcat() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{ast.TypeList},
		ReturnType:   ast.TypeList,
		Variadic:     true,
		VariadicType: ast.TypeList,
		Callback: func(inputs []interface{}) (interface{}, error) {
			var result []ast.Variable
			for _, input := range inputs {
				for _, v := range input.([]ast.Variable) {
					switch v.Type {
					case ast.TypeString:
						result = append(result, v)
					case ast.TypeList:
						result = append(result, v)
					case ast.TypeMap:
						result = append(result, v)
					default:
						return nil, fmt.Errorf("concat() does not support lists of %s", v.Type.Printable())
					}
				}
			}

			return result, nil
		},
	}
}
