package interpolation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hil"
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

// interpolationFuncSplit will split a string into substrings separated by a separator string.
func interpolationFuncSplit() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString, ast.TypeString},
		ReturnType: ast.TypeList,
		Callback: func(inputs []interface{}) (interface{}, error) {
			var result []ast.Variable
			val := inputs[0].(string)
			sep := inputs[1].(string)
			for _, sub := range strings.Split(val, sep) {
				nativeVal, err := hil.InterfaceToVariable(sub)
				if err != nil {
					return nil, err
				}
				result = append(result, nativeVal)
			}
			return result, nil
		},
	}
}

// interpolationFuncReplace replaces the occurrences of a value on the provided string with another value.
// The number of occurrences to replace is the last argument to the function.
func interpolationFuncReplace() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString, ast.TypeString, ast.TypeString, ast.TypeInt},
		ReturnType: ast.TypeString,
		Callback: func(inputs []interface{}) (interface{}, error) {
			input := inputs[0].(string)
			search := inputs[1].(string)
			replace := inputs[2].(string)
			count := inputs[3].(int)

			result := strings.Replace(input, search, replace, count)
			return result, nil
		},
	}
}

// interpolationFuncContains will check if a string contains the portion provided
func interpolationFuncContains() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString, ast.TypeString},
		ReturnType: ast.TypeBool,
		Callback: func(inputs []interface{}) (interface{}, error) {
			val := inputs[0].(string)
			portion := inputs[1].(string)
			return strings.Contains(val, portion), nil
		},
	}
}
