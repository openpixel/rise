package config

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
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
	Templates map[string]ast.Variable
}

// LoadConfigFiles will load all config files and merge values into appropriate values
func LoadConfigFiles(configFiles []string) (*Result, error) {
	vars := make(map[string]ast.Variable)
	templates := make(map[string]ast.Variable)
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

		templates, err = prepareTemplates(templates, config)
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
		astVar, err := hil.InterfaceToVariable(variable.Value)
		if err != nil {
			return err
		}
		vars[fmt.Sprintf("var.%s", variable.Name)] = astVar
	}
	return nil
}

func prepareTemplates(templates map[string]ast.Variable, config *Config) (map[string]ast.Variable, error) {
	for _, template := range config.Templates {
		astVar, err := hil.InterfaceToVariable(template.Content)
		if err != nil {
			return nil, err
		}
		if astVar.Type != ast.TypeString {
			return nil, fmt.Errorf("template %s content must be a string", template.Name)
		}
		templates[fmt.Sprintf("tmpl.%s", template.Name)] = astVar
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
