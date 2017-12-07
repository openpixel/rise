Installation
------------

There are multiple ways that rise can be installed.

Binaries
^^^^^^^^

You can find binaries for the latest release on the `releases <https://github.com/OpenPixel/rise/releases>`_ page. For example, to install on linux:

.. code-block:: sh

   $ wget https://github.com/openpixel/rise/releases/download/v0.0.3/rise_0.0.3_linux_amd64.tar.gz
   $ tar xvf rise_0.0.3_linux_amd64.tar.gz -C /usr/local/bin rise
   $ which rise
   /usr/local/bin/rise

Go Toolchain
^^^^^^^^^^^^

As rise was written in Go, it can be installed via the Go toolchain:

.. code-block:: sh

   $ go get -u github.com/openpixel/rise
