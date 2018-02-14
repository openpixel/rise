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
	Name    string `hcl:",key"`
	Content string `hcl:"content"`
	Count   int    `hcl:"count"`
}

// Result is the result of merging multiple config files
type Result struct {
	Variables map[string]ast.Variable
	Templates map[string]string
}

// LoadConfigFiles will load all config files and merge values into appropriate values
func LoadConfigFiles(configFiles []string) (*Result, error) {
	vars := make(map[string]ast.Variable)
	templates := make(map[string]string)
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

		templates, err = prepareTemplates(vars, config)
		if err != nil {
			return nil, err
		}
	}

	result := &Result{
		Variables: vars,
		Templates: templates,
	}
	return result, nil
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

func prepareTemplates(vars map[string]ast.Variable, config *Config) (map[string]string, error) {
	templates := make(map[string]string)
	evalConfig := &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			FuncMap: interpolation.CoreFunctions,
			VarMap:  vars,
		},
	}

	for _, template := range config.Templates {
		tree, err := hil.Parse(template.Content)
		if err != nil {
			return nil, err
		}

		result, err := hil.Eval(tree, evalConfig)
		if err != nil {
			return nil, err
		}

		templates[template.Name] = result.Value.(string)
	}

	return templates, nil
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
