package formatters

import (
	"fmt"

	"code/diff"
)

type formatFn func([]diff.Diff) string

var formatters = map[string]formatFn{
	"stylish": FmtStylish,
}

func Format(diffNodes []diff.Diff, format string) (string, error) {
	formatter, ok := formatters[format]
	if !ok {
		return "", fmt.Errorf("unsupported format name: %q", format)
	}

	return formatter(diffNodes), nil
}
