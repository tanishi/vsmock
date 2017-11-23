package main

import (
	"os"

	"github.com/tanishi/vsmock/cli"
)

const Version string = "v0.1.0"

func main() {
	c := &cli.CLI{
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
	}
	os.Exit(c.Run(os.Args))
}
