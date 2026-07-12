package differ

import (
	"github.com/CosmoS1X/differ/diff"
	"github.com/CosmoS1X/differ/formatters"
	"github.com/CosmoS1X/differ/parsers"
)

// Gen compares two configuration files and returns the formatted diff.
//
// It parses both input files, builds a diff tree, and formats the result
// according to the provided output format.
func Gen(path1, path2, format string) (string, error) {
	parsedData1, err := parsers.ParseFile(path1)
	if err != nil {
		return "", err
	}
	parsedData2, err := parsers.ParseFile(path2)
	if err != nil {
		return "", err
	}

	diffNodes := diff.Build(parsedData1, parsedData2)

	return formatters.Format(diffNodes, format)
}
