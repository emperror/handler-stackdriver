package stackdriver_test

import (
	"context"

	"cloud.google.com/go/errorreporting"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"emperror.dev/handler/stackdriver"
)

func ExampleNew() {
	// Create the client
	ctx := context.Background()
	client, err := errorreporting.NewClient(ctx, "my-gcp-project", errorreporting.Config{
		ServiceName:    "myservice",
		ServiceVersion: "v1.0",
	})

	if err != nil {
		// TODO: handle error
		return
	}

	defer client.Close()

	// Create the handler
	_ = stackdriver.New(client)

	// Output:
}

func ExampleNew_withCredentials() {
	// Create the client
	ctx := context.Background()
	client, err := errorreporting.NewClient(
		ctx,
		"my-gcp-project",
		errorreporting.Config{
			ServiceName:    "myservice",
			ServiceVersion: "v1.0",
		},
		option.WithCredentials(&google.Credentials{}),
		// OR
		// option.WithCredentialsFile("path/to/google_credentials.json"),
	)

	if err != nil {
		// TODO: handle error
		return
	}

	defer client.Close()

	// Create the handler
	_ = stackdriver.New(client)

	// Output:
}
