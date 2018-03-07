# rise

rise is a powerful text interpolation CLI tool.

## Installation

There are multiple ways that rise can be installed.

### Binaries

You can find binaries for the latest release on the [releases](https://github.com/openpixel/rise/releases) page. For example, to install on linux:

```bash
$ export RISE_VERSION=X.X.X
$ wget https://github.com/openpixel/rise/releases/download/v${RISE_VERSION}/rise_${RISE_VERSION}_linux_amd64.tar.gz
$ tar xvf rise_${RISE_VERSION}_linux_amd64.tar.gz -C /usr/local/bin rise
$ which rise
/usr/local/bin/rise
```

### Go Toolchain

As rise was written in Go, it can be installed via the Go toolchain:

```bash
go get -u github.com/openpixel/rise
```
