Common
======

env
---

.. js:function:: env(key)

   :param string key: The environment key to get
   :return: The value of the environment variable, or "" if it doesn't exist

Example
^^^^^^^

.. code-block:: guess

   env("FOO") // foo

length
------

.. js:function:: length(object)

   :param string/list/map object: The object to return the length of
   :return: The length of the object.
            If object is a string, it is the number of characters.
            If object is a list, it is the number of items in the list.
            If object is a map, it is the number of keys in the map.

Example
^^^^^^^

.. code-block:: guess

   length("hello") // 5
   length(["foo", "bar"]) // 2
   length({"foo": "bar"}) // 1