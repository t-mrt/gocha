package main

import (
	"fmt"
	"io"

	"github.com/t-mrt/gocha"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	ExitCodeOK = iota
	ExitCodeError
)

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var app = kingpin.New("gocha", "Random strings generater based on a pattern.")

	var pattern = app.Arg("pattern", "Regular expression").Required().String()
	var num = app.Flag("number-of-lines", "Number of lines").Short('n').Int()

	var exitCode int
	kingpin.CommandLine.Writer(c.errStream).Terminate(func(i int) {
		exitCode = i
	})
	kingpin.MustParse(app.Parse(args[1:]))

	if exitCode == ExitCodeError {
		return ExitCodeError
	}

	err, g := gocha.New(*pattern)

	if err != nil {
		fmt.Fprintf(c.errStream, "gocha: %v\n", err.Error())
		return ExitCodeError
	}

	if *num <= 0 {
		*num = 1
	}

	for i := 0; i < *num; i++ {
		fmt.Fprintln(c.outStream, g.Gen())
	}

	return ExitCodeOK
}
