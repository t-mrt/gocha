package main

import (
	"flag"
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

	var n int
	var p string

	flags := flag.NewFlagSet("gocha", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	flag.IntVar(&n, "n", 1, "number of lines")
	flag.StringVar(&p, "p", "", "regexp pattern")
	flag.Parse()

	err, g := gocha.New(p)
	if err != nil {
		return ExitCodeError
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(c.outStream, g.Gen())
	}

	return ExitCodeOK
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
