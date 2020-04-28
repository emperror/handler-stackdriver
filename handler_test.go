package stackdriver

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/errorreporting"
	"emperror.dev/errors"
)

func TestHandler(t *testing.T) {
	projectID := os.Getenv("GOOGLE_PROJECT")
	if projectID == "" {
		t.Skip("project ID not specified")
	}

	ctx := context.Background()
	client, err := errorreporting.NewClient(ctx, projectID, errorreporting.Config{
		ServiceName:    "handler-test",
		ServiceVersion: "v1.0",
		OnError: func(err error) {
			t.Fatal(err)
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := client.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	handler := New(client)

	err = errors.New("error")

	handler.Handle(err)
}
