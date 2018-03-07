# Strings

## lower

### Parameters

* *input* - `string` - The input to convert

### Returns

* `string` - The lowercase value of the input

### Examples

```
lower("FOO") // foo
```

---

## upper

### Parameters
* *input* - `string` - The input to convert

### Returns

* `string` - The uppercase value of input

### Examples

```
upper("foo") // FOO
```

---

## join

### Parameters

* *sep* - `string` - The separator between each value
* *values* - `list` - The list of values to join

### Returns

* `string` - The values joined by sep

### Examples

```
join(",", ["foo", "bar"]) // "foo,bar"
```

---

## split

### Parameters

* *value* - `string` - The value to split
* *sep* - `string` - The separator to use for splitting

### Returns

* `list` - The list of split values

### Examples

```
split("foo,bar", ",") // ["foo", "bar"]
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

```
replace("!!!hello!!!", "!", "", -1) // "hello"
```

---

## contains

### Parameters

* *value* - `string` - The value to search for substr
* *substr* - `string` - The substring to search for

### Returns

* `bool` - If substr was found in value or not

### Examples

```
contains("foo bar", "foo") // true
```
