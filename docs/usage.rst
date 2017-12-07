Usage
=====

CLI
---

You can see the usage documation for the CLI by running rise --help.

.. code-block:: bash

   $ rise --help
   A powerful text interpolation tool.

   Usage:
     rise [flags]

   Flags:
     -h, --help                  help for rise
     -i, --input string          The file to perform interpolation on
     -o, --output string         The file to output
     -V, --vars stringSlice      The files that contains the variables to be interpolated

Input (required)
^^^^^^^^^^^^^^^^

The input should be a string that references a file to run the interpolation against.

Output (optional)
^^^^^^^^^^^^^^^^^

The output is the location that the interpolated content should be written to. If not set, it will print to stdout.

Variable Files (optional)
^^^^^^^^^^^^^^^^^^^^^^^^^

The variable files should be in hcl compatible formats. See https://github.com/hashicorp/hcl for reference. Rise loads the files in the order they are supplied, so the latest reference of a variable will always be used. For example, if we had two files that looked like this

.. code-block:: guess
   :caption: vars.hcl
   :linenos:

   variable "i" {
     value = 6
   }

.. code-block:: guess
   :caption: vars2.hcl
   :linenos:

   variable "i" {
     value = 10
   }

And ran the following command

.. code-block:: bash

   $ rise ... --vars vars.hcl --vars vars2.hcl

The value of `i` would be `10`.

Basic Example
^^^^^^^^^^^^^

Look in the `examples <https://github.com/OpenPixel/rise/tree/master/examples>`_ directory for an example, including inheritance:

.. code-block:: bash

   $ rise -i ./examples/input.json -o ./examples/output.json --vars ./examples/vars.hcl --vars ./examples/vars2.hcl

API
---

rise can also be used within Go code.

.. code-block:: go

   import (
       //...
       "github.com/openpixel/rise/template"
       //..
   )

   vars := map[string]ast.Variable{}

   tmpl, err := template.NewTemplate(vars)
   // handle error

   input := `${lower("FOO")}`
   result, err := tmpl.Render(input)
   // handle error

   fmt.Printf("Value: %s", result.Value.(string)) // Value: foo
