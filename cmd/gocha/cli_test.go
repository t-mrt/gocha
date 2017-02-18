package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("ExitCodeOK", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split("gocha -n 1 '[a]{3}'", " ")
		status := cli.Run(args)

		if status != ExitCodeOK {
			t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
		}

		expected := fmt.Sprintf("")
		if errStream.String() != expected {
			t.Errorf("Output=%q, want %q", errStream.String(), expected)
		}

		expected = fmt.Sprintf("'aaa'\n")
		if outStream.String() != expected {
			t.Errorf("Output=%q, want %q", outStream.String(), expected)
		}
	})

	t.Run("ExitCodeError", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split("gocha -n 1 '.{100000}'", " ")
		status := cli.Run(args)

		if status != ExitCodeError {
			t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
		}

		expected := fmt.Sprintf("gocha: error parsing regexp: invalid repeat count: `{100000}`\n")
		if errStream.String() != expected {
			t.Errorf("Output=%q, want %q", errStream.String(), expected)
		}

		expected = fmt.Sprintf("")
		if outStream.String() != expected {
			t.Errorf("Output=%q, want %q", outStream.String(), expected)
		}
	})

	t.Run("RequirePattern", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split("gocha -n 1", " ")
		status := cli.Run(args)

		if status != ExitCodeError {
			t.Errorf("ExitStatus=%d, want %d", status, ExitCodeError)
		}

		expected := fmt.Sprintf("gocha.test: error: required argument 'pattern' not provided, try --help\n")
		if errStream.String() != expected {
			t.Errorf("Output=%q, want %q", errStream.String(), expected)
		}

		expected = fmt.Sprintf("")
		if outStream.String() != expected {
			t.Errorf("Output=%q, want %q", outStream.String(), expected)
		}
	})
}
