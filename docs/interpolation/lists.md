# Lists

## list

### Parameters

* *value* - `any (variadic)` - The values that make up a list

### Returns

* `list` - The list made up of the arguments provided

### Examples

```dart
${list("1", "2")} // ["1", "2"]
```

---

## concat

### Parameters

* *lists* - `list (variadic)` - The lists to concat

### Returns

* `list` - The resulting list

### Examples

```dart
${concat(list("1", "2"), list("3", "4"))} // ["1", "2", "3", "4"]
```

---

## unique

### Parameters

* *list* - `list` - The lists to extract unique values from

### Returns

* `list` - The resulting list with the unique items

### Examples

```dart
${unique(list("1", "2", "2", "1", "3"))} // ["3"]
```

---

## join

### Parameters

* *sep* - `string` - The separator between each value
* *values* - `list` - The list of values to join

### Returns

* `string` - The values joined by sep

### Examples

```dart
${join(",", ["foo", "bar"])} // "foo,bar"
```
