# Maps

## has

### Parameters

* *map* - `map` - The map to perform the lookup on
* *key* - `string` - The key to lookup

### Returns

* `bool` - A bool as to whether or not the key exists

### Examples

```
has({"foo": "bar"}, "foo") // true
```

## map

### Parameters

* *key/value* - `any (variadic)` - The key or value to be added to the map. It should alternate between key and value with the first parameter being a key

### Returns

* `map` - A map consisting of the provided key/value pairs

### Examples

```
map("foo", "bar") // {"foo": "bar"}
map("foo", list(1, 2, 3), "bar", list(3, 2, 1)) // {"foo": [1, 2, 3], "bar": [3, 2, 1]}
```

## keys

### Parameters

* *map* - `map` - The map to perform extraction on

### Returns

* `list` - The list of keys from the map

### Examples

```
keys(map("foo", "bar")) // ["foo"]
```

## merge

### Parameters
* *map* - `map (variadic)` - One or many maps to merge

### Returns

* `map` - The result of the merge

### Examples

```
merge(map("foo", "bar"), map("foo", "bar2", "hello", "there")) // {"foo": "bar2", "hello": "there"}
```
