// Package stackdriver provides an error handler integration for Stackdriver.
package stackdriver

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"strings"

	"cloud.google.com/go/errorreporting"
	"emperror.dev/errors"
)

// Handler sends errors to Stackdriver.
type Handler struct {
	client *errorreporting.Client
}

// New returns a new Handler.
func New(client *errorreporting.Client) *Handler {
	return &Handler{
		client: client,
	}
}

// Handle sends an error to Stackdriver.
func (h *Handler) Handle(err error) {
	entry := errorreporting.Entry{
		Error: err,
	}

	var stackTracer interface{ StackTrace() errors.StackTrace }
	if errors.As(err, &stackTracer) {
		entry.Stack = h.formatStack(stackTracer.StackTrace())
	}

	h.client.Report(entry)
}

// Based on https://github.com/googleapis/google-cloud-go/issues/1084#issuecomment-474565019
// Stackdriver does not accept the pkg/errors stack trace format, so we have to format it manually.
func (h *Handler) formatStack(stackTrace errors.StackTrace) []byte {
	var buf bytes.Buffer

	// routine id and state aren't available in pure go, so we have to hard-code these
	buf.WriteString(fmt.Sprintf("goroutine 1 [running]:\n"))

	// format each frame of the stack to match runtime.Stack's format
	var lines []string
	for _, frame := range stackTrace {
		pc := uintptr(frame) - 1
		fn := runtime.FuncForPC(pc)

		if fn != nil {
			file, line := fn.FileLine(pc)
			lines = append(lines, fmt.Sprintf("%s()\n\t%s:%d +%#x", fn.Name(), file, line, fn.Entry()))
		}
	}

	buf.WriteString(strings.Join(lines, "\n"))

	return buf.Bytes()
}

// HandleContext sends an error to Stackdriver.
func (h *Handler) HandleContext(_ context.Context, err error) {
	h.Handle(err)
}

// Close calls Close on the underlying client.
func (h *Handler) Close() error {
	return h.client.Close()
}
