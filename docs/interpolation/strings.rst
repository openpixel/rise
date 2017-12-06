Strings
=======

lower
-----

.. js:function:: lower(input)

   :param string input: The input to convert
   :return: The lowercase value of input

Example
^^^^^^^

.. code-block:: guess

   lower("FOO") // foo

upper
-----

.. js:function:: upper(input)

   :param string input: The input to convert
   :return: The uppercase value of input

Example
^^^^^^^

.. code-block:: guess

   upper("foo") // FOO

join
----

.. js:function:: join(sep, values)

   :param string sep: The separator between each value
   :param list values: The list of values to join
   :returns: The values joined by sep

Example
^^^^^^^

.. code-block:: guess

   join(",", ["foo", "bar"]) // "foo,bar"

split
-----

.. js:function:: split(value, sep)

   :param string value: The value to split
   :param string sep: The separator to use for splitting
   :returns: The list of split values

Example
^^^^^^^

.. code-block:: guess

   split("foo,bar", ",") // ["foo", "bar"]

replace
-------

.. js:function:: replace(value, original, new, count)

   :param string value: The value to run replace against
   :param string original: The original value to search for
   :param string new: The new value to replace occurences with
   :param int count: The number of times to replace
   :returns: The updated value

Example
^^^^^^^

.. code-block:: guess

   replace("!!!hello!!!", "!", "", -1) // "hello"

contains
--------

.. js:function:: contains(value, substr)

   :param string value: The value to search for substr
   :param string substr: The substring to search for
   :returns: If substr was found in value or not

Example
^^^^^^^

.. code-block:: guess

   contains("foo bar", "foo") // true
