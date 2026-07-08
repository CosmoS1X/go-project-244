package code

import (
	"code/diff"
	"code/formatters"
	"code/parsers"
)

func GenDiff(path1, path2, format string) (string, error) {
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
