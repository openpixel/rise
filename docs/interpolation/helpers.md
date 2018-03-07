# Helpers

## env

### Parameters

* *key* - `string`- The environment key to get

### Returns

* `string` - The value of the environment variable, or "" if it doesn't exist

### Examples

```dart
${env("FOO")} // foo
```

---

## length

### Parameters

* *object* - `string/list/map` - The object to return the length of

### Returns

* `int` - The length of the object:
  * If object is a string, it is the number of characters.
  * If object is a list, it is the number of items in the list.
  * If object is a map, it is the number of keys in the map.

### Examples

```dart
${length("hello")} // 5
${length(["foo", "bar"])} // 2
${length({"foo": "bar"})} // 1
```
