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
		Name:                   "gendiff",
		Usage:                  "Compares two configuration files and shows a difference.",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"f"},
				Value:       "stylish",
				Usage:       "output format",
				DefaultText: "stylish",
			},
		},
		ArgsUsage: "<file1> <file2>",
		Arguments: []cli.Argument{
			&cli.StringArgs{
				Name: "paths",
				Min:  2,
				Max:  2,
			},
		},
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			argsLen := c.Args().Len()
			if argsLen != 2 {
				return ctx, fmt.Errorf("expected exactly 2 path arguments, got %d", argsLen)
			}
			return ctx, nil
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			format := c.String("format")
			fmt.Fprintf(stdout, "Format: %s\n", format)

			paths := c.StringArgs("paths")

			diff, err := code.GenDiff(paths[0], paths[1], c.String("format"))
			if err != nil {
				return err
			}

			fmt.Fprintln(stdout, diff)
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
