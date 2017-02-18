package main

import (
	"fmt"
	"io"
	"os"

	"github.com/t-mrt/gocha"
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

	var pattern = kingpin.Arg("pattern", "Regular expression").Required().String()
	var num = kingpin.Flag("number-of-lines", "Number of lines").Short('n').Int()
	kingpin.Parse()

	err, g := gocha.New(*pattern)

	if err != nil {
		fmt.Fprintln(c.outStream, "gocha: error: Invalid regular expression")
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

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
