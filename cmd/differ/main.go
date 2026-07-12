package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/urfave/cli/v3"

	"github.com/CosmoS1X/differ"
)

const (
	FormatStylish = "stylish"
	FormatPlain   = "plain"
	FormatJSON    = "json"
)

var allowedFormats = []string{FormatStylish, FormatPlain, FormatJSON}

func run(args []string, stdout, stderr io.Writer) int {
	formatsList := strings.Join(allowedFormats, ", ")

	cmd := &cli.Command{
		Name:                   "differ",
		Usage:                  "Compares two configuration files and shows a difference.",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"f"},
				Value:       FormatStylish,
				Usage:       fmt.Sprintf("output format (%s)", formatsList),
				DefaultText: FormatStylish,
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
			// Validate the number of path arguments
			argsLen := c.Args().Len()
			if argsLen != 2 {
				return ctx, fmt.Errorf("expected exactly 2 path arguments, got %d", argsLen)
			}

			// Validate the output format
			format := c.String("format")
			if !slices.Contains(allowedFormats, format) {
				return ctx, fmt.Errorf("unknown output format: %q (allowed: %s)", format, formatsList)
			}

			return ctx, nil
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			paths := c.StringArgs("paths")

			diff, err := differ.Gen(paths[0], paths[1], c.String("format"))
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
