package template

import (
	"bytes"
	"io"
	"testing"

	"github.com/hashicorp/hil/ast"

	"github.com/openpixel/rise/internal/config"
)

func stringFromReader(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.String()
}

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

		result, err := tmpl.Render(bytes.NewBufferString(`${has(var.foo, "bar")} ${var.foo["bar"]}`))
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		if stringFromReader(result) != "true Bar" {
			t.Fatalf("Unexpected result: Expected %s, got %s", "true", stringFromReader(result))
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

		result, err := tmpl.Render(bytes.NewBufferString(`${tmpl.foo}`))
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		if stringFromReader(result) != "This is a template foo" {
			t.Fatalf("Unexpected result: Expected %s, got %s", "This is a template foo", stringFromReader(result))
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

		result, err := tmpl.Render(bytes.NewBufferString(`${lower("FOO")}`))
		if err != nil {
			t.Fatalf("Unexpected err: %s", err)
		}

		if stringFromReader(result) != "foo" {
			t.Fatalf("Unexpected result: Expected %s, got %s", "foo", stringFromReader(result))
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
		_, err = tmpl.Render(bytes.NewBufferString(`${tmpl.foo}`))
		if err == nil {
			t.Fatal("unexpected nil err")
		}
	})
}
