package template

import (
	"testing"

	"github.com/hashicorp/hil/ast"
	"github.com/openpixel/rise/internal/config"
)

func TestTemplate_Render(t *testing.T) {
	t.Run("Test render with variable passed", func(t *testing.T) {
		vars := map[string]ast.Variable{
			"var.foo": ast.Variable{
				Type: ast.TypeMap,
				Value: map[string]ast.Variable{
					"bar": ast.Variable{
						Type:  ast.TypeString,
						Value: "Bar",
					},
				},
			},
		}
		config := &config.Result{
			Variables: vars,
			Templates: map[string]ast.Variable{},
		}
		tmpl, err := NewTemplate(config)
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		result, err := tmpl.Render(`${has(var.foo, "bar")} ${var.foo["bar"]}`)
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		if result.Value.(string) != "true Bar" {
			t.Fatalf("Unexpected result: Expected %s, got %s", "true", result.Value.(string))
		}
	})

	t.Run("Test render with template", func(t *testing.T) {
		tmpls := map[string]ast.Variable{
			"tmpl.foo": ast.Variable{
				Type:  ast.TypeString,
				Value: "This is a template ${lower(\"FOO\")}",
			},
		}
		config := &config.Result{
			Variables: map[string]ast.Variable{},
			Templates: tmpls,
		}
		tmpl, err := NewTemplate(config)
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		result, err := tmpl.Render(`${tmpl.foo}`)
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		if result.Value.(string) != "This is a template foo" {
			t.Fatalf("Unexpected result: Expected %s, got %s", "This is a template foo", result.Value.(string))
		}
	})

	t.Run("Test render with nil variables", func(t *testing.T) {
		tmpl, err := NewTemplate(&config.Result{
			Variables: nil,
			Templates: nil,
		})
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		result, err := tmpl.Render(`${lower("FOO")}`)
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		if result.Value.(string) != "foo" {
			t.Fatalf("Unexpected result: Expected %s, got %s", "foo", result.Value.(string))
		}
	})

	t.Run("Test missing template", func(t *testing.T) {
		tmpl, err := NewTemplate(&config.Result{
			Variables: nil,
			Templates: nil,
		})
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}
		_, err = tmpl.Render(`${tmpl.foo}`)
		if err == nil {
			t.Fatal("unexpected nil err")
		}
	})
}
