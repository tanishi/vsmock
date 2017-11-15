package main

import (
	"flag"
	"io"
	"os"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("vsmock", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	var url string
	flags.StringVar(&url, "u", "", "url")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if url == "" {
		url = os.Getenv("GOVC_URL")
	}

	if url == "" {
		flags.Usage()
		return ExitCodeParseFlagError
	}

	return ExitCodeOK
}
