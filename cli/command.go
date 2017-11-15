package cli

import (
	"context"
	"flag"
	"io"
	"os"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

type HasFlags interface {
	Register(ctx context.Context, f *flag.FlagSet)
	Process(ctx context.Context) error
}

type Command interface {
	HasFlags

	Run(ctx context.Context, f *flag.FlagSet) error
}

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("vsmock", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)

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
