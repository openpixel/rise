# Configuration Variables

Configuration files are made up of a number of different structures to assist with interpolation. The most common structure to be used are variables. Once declared, a variable can be referenced throughout interpolation and can even be referenced by other variables.

## Usage

The best way to explain the usage of a variable is through an example. In the below example, we will create a simple string variable and show how it is referenced within a file.

The variable in the configuration file:
```dart
variable "foo" {
  value = "bar"
}
```
Would be referenced like so:
```dart
${var.foo}
```

Note the `var` prefix that is placed on the name it was declared with. All variables are accessed through the `var` prefix.

There are a number of different data types that can be supplied. Please see the [Syntax](overview.md#Syntax) section for more details.

## Examples

### Basic String
```dart
variable "foo" {
  value = "bar"
}
```

### Multi-line String
```dart
variable "foo" {
  value = <<EOF
  Line 1
  Line 2
  EOF
}
```

### Lists
```dart
variable "foo" {
  value = ["Thing 1", "Thing 2"]
}
```

### Maps
```dart
variable "foo" {
  value = {
    "name": "John",
    "location": "Portland, OR"
  }
}
```

### Boolean
```dart
variable "foo" {
  value = false
}
```
