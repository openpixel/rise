# Configuration Overview

When implementing more advanced interpolation, it is highly likely that you will need to utilize configuration files. Simply put, configuration files are containers for metadata such as variables and templates that can be referenced via interpolation.

## Syntax

Configuration files use [HashiCorp configuration language (HCL)](https://github.com/hashicorp/hcl) syntax. Please visit the HCL repository for detailed overview. Configuration files can be both hcl or json formatted.

For example, an example hcl configuration looks like this:
```dart
variable "name" {
  value = "John"
}

template "welcome" {
  content = "hello, ${name}"
}
```
The equivalent json structure:
```json
{
  "variable": {
    "name": {
      "value": "John"
    }
  },
  "template": {
    "welcome": {
      "content": "hello, ${name}"
    }
  }
}
```
