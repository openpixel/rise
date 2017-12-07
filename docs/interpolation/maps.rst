Maps
====

has
---

.. js:function:: has(map, key)

   :param map map: The map to perform the lookup on
   :param string key: The key to look for
   :return: A bool as to whether or not the key exists

Example
^^^^^^^

.. code-block:: guess

   has({"foo": "bar"}, "foo") // true

map
---

.. js:function:: map(key, value, ...)

   :param string key: The key for the key/value pair
   :param any value: The value for the key/value/pair
   :return: A map object consisting of the provided key/value pairs

Example
^^^^^^^

.. code-block:: guess

   map("foo", "bar") // {"foo": "bar"}
   map("foo", list(1, 2, 3), "bar", list(3, 2, 1)) // {"foo": [1, 2, 3], "bar": [3, 2, 1]}

keys
----

.. js:function:: keys(map)

   :param map map: The map to extract keys
   :return: The list of keys from the map

Example
^^^^^^^

.. code-block:: guess

   keys(map("foo", "bar")) // ["foo"]

merge
-----

.. js:function:: merge(map1, ...)

   :param map map1:  One or many maps to merge
   :return: The result of the merge

Example
^^^^^^^

.. code-block:: guess

   merge(map("foo", "bar"), map("foo", "bar2", "hello", "there")) // {"foo": "bar2", "hello": "there"}