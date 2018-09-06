package interpolation

import (
	"errors"
	"math"

	"github.com/hashicorp/hil/ast"
)

// interpolationFuncMax returns the largest value in a series of floats
func interpolationFuncMax() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{},
		ReturnType:   ast.TypeFloat,
		Variadic:     true,
		VariadicType: ast.TypeFloat,
		Callback: func(inputs []interface{}) (interface{}, error) {
			if len(inputs) < 2 {
				return nil, errors.New("max function requires at least 2 inputs")
			}
			max := inputs[0].(float64)
			for _, input := range inputs[1:] {
				max = math.Max(max, input.(float64))
			}
			return max, nil
		},
	}
}

// interpolationFuncMin returns the smallest value in a series of floats
func interpolationFuncMin() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{},
		ReturnType:   ast.TypeFloat,
		Variadic:     true,
		VariadicType: ast.TypeFloat,
		Callback: func(inputs []interface{}) (interface{}, error) {
			if len(inputs) < 2 {
				return nil, errors.New("min function requires at least 2 inputs")
			}
			min := inputs[0].(float64)
			for _, input := range inputs[1:] {
				min = math.Min(min, input.(float64))
			}
			return min, nil
		},
	}
}

// InterpolationFuncAvg calculates the average of the given numbers
func interpolationFuncAvg() ast.Function {
	return ast.Function{
		ArgTypes:     []ast.Type{},
		ReturnType:   ast.TypeFloat,
		Variadic:     true,
		VariadicType: ast.TypeFloat,
		Callback: func(inputs []interface{}) (interface{}, error) {
			inputLen := len(inputs)
			if inputLen == 0 {
				return nil, errors.New("avg function requires at least 1 inputs")
			}
			total := float64(0)
			for _, input := range inputs {
				total += input.(float64)
			}
			return total / float64(inputLen), nil
		},
	}
}
