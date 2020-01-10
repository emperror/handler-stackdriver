# Stackdriver error handler

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emperror/handler-stackdriver/CI?style=flat-square)
[![Codecov](https://img.shields.io/codecov/c/github/emperror/handler-stackdriver?style=flat-square)](https://codecov.io/gh/emperror/handler-stackdriver)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/handler/stackdriver?style=flat-square)](https://goreportcard.com/report/emperror.dev/handler/stackdriver)
[![GolangCI](https://golangci.com/badges/github.com/emperror/handler-stackdriver.svg)](https://golangci.com/r/github.com/emperror/handler-stackdriver)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.12-61CFDD.svg?style=flat-square)](https://github.com/emperror/errors)
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
	"context"

	"cloud.google.com/go/errorreporting"
	"emperror.dev/handler/stackdriver"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func main() {
	// Create the client
	ctx := context.Background()
	client, err := errorreporting.NewClient(
		ctx,
		"my-gcp-project",
		errorreporting.Config{
			ServiceName:    "myservice",
			ServiceVersion: "v1.0",
		},
		option.WithCredentialsFile("path/to/google_credentials.json"),
	)
	if err != nil {
		// TODO: handle error
	}
	defer client.Close()

	// Create the handler
	_ = stackdriver.New(client)
}
```


## Development

When all coding and testing is done, please run the test suite:

```bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
