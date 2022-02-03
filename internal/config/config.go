package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
)

// config is a variable file config definition
type config struct {
	Variables []variableConfig `hcl:"variable"`
	Templates []templateConfig `hcl:"template"`
}

// variableConfig defines the structure for our variable config sections
type variableConfig struct {
	Name  string      `hcl:",key"`
	Value interface{} `hcl:"value"`
}

// templateConfig defines the structure for our template config sections
type templateConfig struct {
	Name    string `hcl:",key"`
	Content string `hcl:"content"` // a string that contains a simple template
	File    string `hcl:"file"`    // a reference to a file that is relative to the config file
	Trim    bool   `hcl:"trim"`    // declares if whitespace should be trimmed from the file contents
}

// Result is the result of merging multiple config files
type Result struct {
	Variables map[string]ast.Variable
	Templates map[string]ast.Variable
}

// LoadExtras attemps to load any additional vars
func LoadExtras(args []string) (map[string]ast.Variable, error) {
	extras := make(map[string]ast.Variable)
	for _, extra := range args {
		extraResult := make(map[string]interface{})

		err := json.Unmarshal([]byte(extra), &extraResult)
		if err != nil {
			return nil, err
		}

		for k, v := range extraResult {
			extras[fmt.Sprintf("var.%s", k)], err = hil.InterfaceToVariable(v)
			if err != nil {
				return nil, err
			}
		}
	}
	return extras, nil
}

// LoadConfigFiles will load all config files and merge values into appropriate values
func LoadConfigFiles(configFiles []string, extras map[string]ast.Variable) (*Result, error) {
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
		err = prepareVariables(vars, config)
		if err != nil {
			return nil, err
		}

		templates, err = prepareTemplates(file, templates, config)
		if err != nil {
			return nil, err
		}
	}

	// apply extras to variable map
	for k, v := range extras {
		vars[k] = v
	}

	return &Result{
		Variables: vars,
		Templates: templates,
	}, nil
}

func prepareVariables(vars map[string]ast.Variable, config *config) error {
	for _, variable := range config.Variables {
		astVar, err := hil.InterfaceToVariable(variable.Value)
		if err != nil {
			return err
		}
		vars[fmt.Sprintf("var.%s", variable.Name)] = astVar
	}
	return nil
}

func prepareTemplates(baseFilePath string, templates map[string]ast.Variable, config *config) (map[string]ast.Variable, error) {
	for _, template := range config.Templates {
		var astVar ast.Variable
		var err error
		if template.File != "" {
			templateFilePath := filepath.Join(filepath.Dir(baseFilePath), template.File)
			var fileContents []byte
			fileContents, err = ioutil.ReadFile(templateFilePath)
			if err != nil {
				return nil, err
			}
			contents := string(fileContents)
			if template.Trim {
				contents = strings.TrimSpace(contents)
			}
			astVar, err = hil.InterfaceToVariable(contents)
		} else {
			astVar, err = hil.InterfaceToVariable(template.Content)
		}
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
func parseConfig(text string) (*config, error) {
	result := &config{}

	hclParseTree, err := hcl.Parse(text)
	if err != nil {
		return nil, err
	}

	if err := hcl.DecodeObject(result, hclParseTree); err != nil {
		return nil, err
	}
	return result, nil
}
