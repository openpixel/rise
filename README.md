[![Build Status](https://travis-ci.org/OpenPixel/rise.svg?branch=master)](https://travis-ci.org/OpenPixel/rise)
[![Go Report Card](https://goreportcard.com/badge/github.com/openpixel/rise)](https://goreportcard.com/report/github.com/openpixel/rise)

# rise

Powerful text interpolation. Documentation can be found here: 

## Installation

### Binaries

You can find binaries for the latest release on the [releases](https://github.com/OpenPixel/rise/releases) page.

### Go toolchain

```
$ go get -u github.com/openpixel/rise
```

## Quickstart

### CLI
You can see the usage documation for the CLI by running `rise --help`.

```
$ rise --help
A powerful text interpolation tool.

Usage:
  rise [flags]

Flags:
  -h, --help                  help for rise
  -i, --input string          The file to perform interpolation on
  -o, --output string         The file to output
      --varFile stringSlice   The files that contains the variables to be interpolated
```

### Variable Files

The variable files should be in hcl compatible formats. See https://github.com/hashicorp/hcl for reference. Rise loads the files in the order they are supplied, so the latest reference of a variable will always be used. For example, if we had two files that looked like this:

vars.hcl
```
variable "i" {
  value = 6
}
```

vars2.hcl
```
variable "i" {
  value = 10
}
```

And ran the following command

```
$ rise ... --varFile vars.hcl --varFile vars2.hcl
```

The value of `i` would be `10`.

### Examples

Look in the [examples](https://github.com/OpenPixel/rise/tree/master/examples) directory for an example, including inheritance:

```
$ rise -i ./examples/input.json -o ./examples/output.json --varFile ./examples/vars.hcl --varFile ./examples/vars2.hcl
```

## Coming Soon

- More interpolation methods
- Deeper documentation with examples for interpolation methods
- More configuration CLI arguments
  - Support for directories as inputs/outputs
  - Support for globs (eg: /tmp/*.json)
  - Support for var overrides at cli level (eg: --var "foo=bar")

## Inspiration

- [hashicorp/hil](https://github.com/hashicorp/hil) - Used to perform interpolation
- [hashicorp/hcl](https://github.com/hashicorp/hcl) - Used as a configuration syntax for variables
- [hashicorp/terraform](https://github.com/hashicorp/terraform) - Inspiration for the tool. A number of the interpolation functions have been extracted directly from terraform.
