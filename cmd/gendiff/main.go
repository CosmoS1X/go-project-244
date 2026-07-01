package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
)

func run(args []string, stdout, stderr io.Writer) int {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Fprint(stdout, "I can't do anything yet, but I will be able to soon!\n")
			fmt.Fprint(stdout, "To display help, use the --help flag\n")
			return nil
		},
	}

	err := cmd.Run(context.Background(), args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %s\n", err.Error())
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}
