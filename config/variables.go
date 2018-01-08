package config

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
	"github.com/openpixel/rise/interpolation"
)

// Config is a variable file config definition
type Config struct {
	Variables []VariableConfig `hcl:"variable"`
	Templates []TemplateConfig `hcl:"template"`
}

// VariableConfig defines the structure for our variable config sections
type VariableConfig struct {
	Name  string      `hcl:",key"`
	Value interface{} `hcl:"value"`
}

// TemplateConfig defines the structure for our template config sections
type TemplateConfig struct {
	Name     string `hcl:",key"`
	Content  string `hcl:"content"`
	Filename string `hcl:"file"`
	Count    int    `hcl:"count"`
}

// LoadConfigFiles will load all config files and merge it into a single
// variable map
func LoadConfigFiles(configFiles []string) (map[string]ast.Variable, error) {
	vars := make(map[string]ast.Variable)
	for _, file := range configFiles {
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		config, err := parseConfig(string(contents))
		if err != nil {
			return nil, err
		}
		err = interpolateVariables(vars, config)
		if err != nil {
			return nil, err
		}
	}

	return vars, nil
}

func interpolateVariables(vars map[string]ast.Variable, config *Config) error {
	for _, variable := range config.Variables {
		err := hil.Walk(&variable.Value, func(data *hil.WalkData) error {
			result, err := hil.Eval(data.Root, &hil.EvalConfig{
				GlobalScope: &ast.BasicScope{
					FuncMap: interpolation.CoreFunctions,
				},
			})
			if err != nil {
				return err
			}
			data.Replace = true
			data.ReplaceValue = result.Value.(string)
			return nil
		})
		if err != nil {
			return err
		}
		astVar, err := hil.InterfaceToVariable(variable.Value)
		if err != nil {
			return err
		}
		vars[variable.Name] = astVar
	}
	return nil
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
