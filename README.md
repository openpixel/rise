[![Build Status](https://travis-ci.org/openpixel/rise.svg?branch=master)](https://travis-ci.org/openpixel/rise)
[![Go Report Card](https://goreportcard.com/badge/github.com/openpixel/rise)](https://goreportcard.com/report/github.com/openpixel/rise)
[![GoDoc](https://godoc.org/github.com/openpixel/rise?status.svg)](https://godoc.org/github.com/openpixel/rise)
[![Coverage Status](https://coveralls.io/repos/github/openpixel/rise/badge.svg?branch=master)](https://coveralls.io/github/openpixel/rise?branch=master)

# rise

Powerful text interpolation. Documentation can be found [here](https://openpixel.gitbooks.io/rise).

Note: rise is still in development and is subject to breaking changes until we reach our first major release.

## Installation

### Binaries

You can find binaries for the latest release on the [releases](https://github.com/OpenPixel/rise/releases) page.

### Go toolchain

```bash
$ go get -u github.com/openpixel/rise
```

## Quickstart

### CLI
You can see the usage documentation for the CLI by running `rise --help`.

```bash
$ rise --help
A powerful text interpolation tool.

Usage:
  rise [flags]
  rise [command]

Available Commands:
  help      Help about any command
  version   Print the version number of rise

Flags:
  -c, --config stringSlice    The files that define the configuration to use for interpolation
  -h, --help                  help for rise
  -i, --input string          The file to perform interpolation on
  -o, --output string         The file to output
```

### Config Files

The config files should be in hcl compatible formats. See https://github.com/hashicorp/hcl for reference. Rise loads the files using FIFO, meaning the last file to reference a key will take precedence. For example, if we had two files that looked like this:

vars.hcl
```hcl
variable "i" {
  value = 6
}
```

vars2.hcl
```hcl
variable "i" {
  value = 10
}
```

And ran the following command

```bash
$ rise ... --config vars.hcl --config vars2.hcl
```

The value of `i` would be `10`.

### Examples

Look in the [examples](https://github.com/OpenPixel/rise/tree/master/examples) directory for an example, including inheritance:

```bash
$ rise -i ./examples/input.json -o ./examples/output.json --config ./examples/vars.hcl --config ./examples/vars2.hcl
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
