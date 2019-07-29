# Stackdriver error handler

[![CircleCI](https://circleci.com/gh/emperror/handler-stackdriver.svg?style=svg)](https://circleci.com/gh/emperror/handler-stackdriver)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/handler/stackdriver?style=flat-square)](https://goreportcard.com/report/emperror.dev/handler/stackdriver)
[![GolangCI](https://golangci.com/badges/github.com/emperror/handler-stackdriver.svg)](https://golangci.com/r/github.com/emperror/handler-stackdriver)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/emperror.dev/handler/stackdriver)

**Error handler integration for [Stackdriver](https://cloud.google.com/stackdriver/)**


## Installation

```bash
go get emperror.dev/handler/stackdriver
```


## Usage

```go
package main

import (
	stackdriverhandler "emperror.dev/handler/stackdriver"
)

func main() {
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
