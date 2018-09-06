package interpolation

import (
	"fmt"
	"strings"

	"reflect"

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

// interpolationFuncUnique will extract the unique values and return a list
func interpolationFuncUnique() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeList},
		ReturnType: ast.TypeList,
		Callback: func(inputs []interface{}) (interface{}, error) {
			list := inputs[0].([]ast.Variable)
			result := make([]ast.Variable, 0, len(list))

			for _, val := range list {
				exists := false
				for _, r := range result {
					if reflect.DeepEqual(val, r) {
						exists = true
					}
				}
				if !exists {
					result = append(result, val)
				}
			}

			return result, nil
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
