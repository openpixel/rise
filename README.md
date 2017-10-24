# rise

Powerful text interpolation.

## Installation

```
$ go get -u github.com/openpixel/rise
```

## Usage

You can see the usage documation for the CLI by running `rise --help`.

```
$ rise --help
A simple template substitution tool.

Usage:
  rise [flags]

Flags:
  -h, --help                  help for rise
  -i, --input string          The file to perform substitution on
  -o, --output string         The file to output
      --varFile stringSlice   The files that contains the variables to be substituted
```

### Input (required)

The input should be a string that references a file to run the interpolation against.

### Output (optional)

The output is the location that the interpolated content should be written to. If not set, it will print to stdout.

### Variable Files (optional)

The variable files should be in hcl compatible formats. See https://github.com/hashicorp/hcl for reference. Rise loads the files in the order they are supplied, so the latest reference of a variable will always be used. For example, if we had two files that looked like this

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

## Basic Example

Look in the [examples](https://github.com/OpenPixel/rise/tree/master/examples) directory for an example, including inheritance:

```
$ rise -i ./examples/basic.tmpl -o ./examples/basic.txt --varFile ./examples/vars.hcl --varFile ./examples/vars2.hcl
```

## Interpolation Methods

- lower
    - Convert the provided argument to lowercase
- upper
    - Convert the provided argument to uppercase
- env
    - Find the provided environment variable