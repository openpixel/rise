package interpolation

import (
	"fmt"
	"sort"

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

// interpolationFuncKeys returns the keys of the provided map sorted in dictionary order
func interpolationFuncKeys() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeMap},
		ReturnType: ast.TypeList,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			mapInput := inputs[0].(map[string]ast.Variable)
			keys := make([]string, 0, len(mapInput)+1)
			result := make([]ast.Variable, 0, len(mapInput)+1)
			for key := range mapInput {
				keys = append(keys, key)
			}
			sort.Strings(keys)

			for _, key := range keys {
				nativeKey, err := hil.InterfaceToVariable(key)
				if err != nil {
					return nil, err
				}
				result = append(result, nativeKey)
			}

			return result, nil
		},
	}
}

// interpolationFuncValues extracts the values from a map
func interpolationFuncValues() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeMap},
		ReturnType: ast.TypeList,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			mapInput := inputs[0].(map[string]ast.Variable)
			values := make([]interface{}, 0, len(mapInput)+1)
			result := make([]ast.Variable, 0, len(mapInput)+1)
			for _, val := range mapInput {
				values = append(values, val)
			}

			for _, val := range values {
				nativeValue, err := hil.InterfaceToVariable(val)
				if err != nil {
					return nil, err
				}
				result = append(result, nativeValue)
			}

			return result, nil
		},
	}
}

// interpolationFuncMerge will merge multiple maps into a single map. Last reference of a key will always win.
func interpolationFuncMerge() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{ast.TypeMap},
		ReturnType:   ast.TypeMap,
		Variadic:     true,
		VariadicType: ast.TypeMap,
		Callback: func(inputs []interface{}) (interface{}, error) {
			result := make(map[string]ast.Variable)

			for _, input := range inputs {
				for k, v := range input.(map[string]ast.Variable) {
					result[k] = v
				}
			}

			return result, nil
		},
	}
}

// interpolationFuncPick will pick the values of the provided keys, and create a new map
func interpolationFuncPick() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{ast.TypeMap},
		ReturnType:   ast.TypeMap,
		Variadic:     true,
		VariadicType: ast.TypeString,
		Callback: func(inputs []interface{}) (interface{}, error) {
			inMap := inputs[0].(map[string]ast.Variable)
			result := make(map[string]ast.Variable)

			for i := 1; i < len(inputs); i++ {
				key := inputs[i].(string)

				if val, ok := inMap[key]; ok {
					result[key] = val
				}
			}

			return result, nil
		},
	}
}

// interpolationFuncOmit will return a map that has omitted the keys provided
func interpolationFuncOmit() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{ast.TypeMap},
		ReturnType:   ast.TypeMap,
		Variadic:     true,
		VariadicType: ast.TypeString,
		Callback: func(inputs []interface{}) (interface{}, error) {
			inMap := inputs[0].(map[string]ast.Variable)

			for i := 1; i < len(inputs); i++ {
				key := inputs[i].(string)
				delete(inMap, key)
			}

			return inMap, nil
		},
	}
}
