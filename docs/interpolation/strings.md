# Strings

## lower

### Parameters

* *input* - `string` - The input to convert

### Returns

* `string` - The lowercase value of the input

### Examples

```dart
${lower("FOO")} // foo
```

---

## upper

### Parameters
* *input* - `string` - The input to convert

### Returns

* `string` - The uppercase value of input

### Examples

```dart
${upper("foo")} // FOO
```

---

## split

### Parameters

* *value* - `string` - The value to split
* *sep* - `string` - The separator to use for splitting

### Returns

* `list` - The list of split values

### Examples

```dart
${split("foo,bar", ",")} // ["foo", "bar"]
```

---

## replace

### Parameters

* *value* - `string` - The value to run replace against
* *original* - `string` - The original value to search for
* *new* - `string` - The new value to replace occurrences with
* *count* - `int` - The number of times to replace

### Returns

- `string` - The updated value

### Examples

```dart
${replace("!!!hello!!!", "!", "", -1)} // "hello"
```

---

## contains

### Parameters

* *value* - `string` - The value to search for substr
* *substr* - `string` - The substring to search for

### Returns

* `bool` - If substr was found in value or not

### Examples

```dart
${contains("foo bar", "foo")} // true
```
