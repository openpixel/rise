package variables

import (
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
)

// Config is a variable file config definition
type Config struct {
	Variables []VariableConfig `hcl:"variable"`
}

// VariableConfig defines the structure for our variable config files
type VariableConfig struct {
	Name  string      `hcl:",key"`
	Value interface{} `hcl:"value"`
}

// LoadVariableFiles will load all variables files and merge it into a single
// variable map
func LoadVariableFiles(varFiles []string) (map[string]ast.Variable, error) {
	vars := make(map[string]ast.Variable)
	for _, file := range varFiles {
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		config, err := parseConfig(string(contents))
		if err != nil {
			return nil, err
		}
		for _, variable := range config.Variables {
			err := hil.Walk(&variable.Value, func(data *hil.WalkData) error {
				result, err := hil.Eval(data.Root, &hil.EvalConfig{
					GlobalScope: &ast.BasicScope{
						FuncMap: map[string]ast.Function{
							"env": ast.Function{
								ArgTypes:   []ast.Type{ast.TypeString},
								ReturnType: ast.TypeString,
								Variadic:   false,
								Callback: func(inputs []interface{}) (interface{}, error) {
									input := inputs[0].(string)
									return os.Getenv(input), nil
								},
							},
						},
					},
				})
				if err != nil {
					return err
				}
				data.Replace = true
				data.ReplaceValue = result.Value.(string)
				return nil
			})
			astVar, err := hil.InterfaceToVariable(variable.Value)
			if err != nil {
				return nil, err
			}
			vars[variable.Name] = astVar
		}
	}

	return vars, nil
}

// parseConfig will parse the text into variable definitions
func parseConfig(text string) (*Config, error) {
	result := &Config{}

	hclParseTree, err := hcl.Parse(text)
	if err != nil {
		return nil, err
	}

	if err := hcl.DecodeObject(result, hclParseTree); err != nil {
		return nil, err
	}
	return result, nil
}
