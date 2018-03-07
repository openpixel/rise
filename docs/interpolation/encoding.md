# Encoding

## base64enc

### Parameters

* *item* - `string`- The value to encode

### Returns

* `string` - The base64 encoded representation of *item*

### Examples

```dart
${base64enc("foo")} // Zm9v
```

---

## base64dec

### Parameters

* *encoding* - `string`- A base64 encoded value

### Returns

* `string` - The decoded representation of *encoding*

### Examples

```dart
${base64dec("Zm9v")} // foo
```

---

## jsonencode

### Parameters

* *item* - `list/map/string`- The value to encode

### Returns

* `string` - The json representation of the value

### Examples

```dart
${jsonencode("foo")} // foo
```

```dart
${jsonencode(list("1", "2"))} // ["1", "2"]
```
