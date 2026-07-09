package formatters

import (
	"fmt"

	"code/diff"
)

type Formatter interface {
	Format(diffNodes []diff.Diff) string
}

var registry = map[string]Formatter{
	"stylish": &stylishFormatter{},
	"plain":   &plainFormatter{},
}

func Format(diffNodes []diff.Diff, format string) (string, error) {
	formatter, ok := registry[format]
	if !ok {
		return "", fmt.Errorf("unsupported format name: %q", format)
	}

	return formatter.Format(diffNodes), nil
}
