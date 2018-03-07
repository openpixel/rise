# Configuration Templates

Configuration files are made up of a number of different structures to assist with interpolation. Along with [Variables](variables.md), they can also contain Templates. Templates provide modular content that can be interpolated.

## Usage

Let's create an example template to show how it can be used. Again, we will start with the most basic example. More advanced examples can be found in the Examples section of this page.

The template in the configuration file:
```dart
template "welcome" {
  content = "hello, ${name}"
}
```
Would be referenced like so:
```dart
${tmpl.welcome}
```

Note the `tmpl` prefix that is placed on the name it was declared with. All templates are accessed through the `tmpl` prefix.

Templates can also load the contents of a file for reference in interpolation. When providing the location of the file, ensure it is *relative to the configuration file location*.

## Examples

### Basic Template
```dart
template "welcome" {
  content = "hello, ${name}"
}
```

### Load File
```dart
template "welcome" {
  file = "./template.txt" // template.txt contains hello, ${name}
}
```

### Load File and Trim
```dart
template "welcome" {
  file = "./template.txt" // template.txt contains hello, ${name}
  trim = true
}
```
