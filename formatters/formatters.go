package formatters

import (
	"fmt"

	"github.com/CosmoS1X/differ/diff"
)

type Formatter interface {
	Format(diffNodes []diff.Diff) (string, error)
}

var registry = map[string]Formatter{
	"stylish": &stylishFormatter{},
	"plain":   &plainFormatter{},
	"json":    &jsonFormatter{},
}

func Format(diffNodes []diff.Diff, format string) (string, error) {
	formatter, ok := registry[format]
	if !ok {
		return "", fmt.Errorf("unsupported format name: %q", format)
	}

	return formatter.Format(diffNodes)
}
