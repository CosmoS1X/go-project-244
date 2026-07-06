package code

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/parsers"
)

const (
	Added     = "added"
	Deleted   = "deleted"
	Unchanged = "unchanged"
	Changed   = "changed"
)

type Diff struct {
	key      string
	value    any
	newValue any
	status   string
}

func getCommonKeys(data1, data2 parsers.ParsedData) []string {
	uniqMap := make(map[string]struct{}, len(data1)+len(data2))

	for k := range data1 {
		uniqMap[k] = struct{}{}
	}
	for k := range data2 {
		uniqMap[k] = struct{}{}
	}

	return slices.Sorted(maps.Keys(uniqMap))
}

func buildDiff(data1, data2 parsers.ParsedData) []Diff {
	commonKeys := getCommonKeys(data1, data2)
	diff := make([]Diff, 0, len(commonKeys))

	for _, k := range commonKeys {
		value, hasKeyInData1 := data1[k]
		newValue, hasKeyInData2 := data2[k]

		switch {
		case value == newValue:
			diff = append(diff, Diff{key: k, value: value, status: Unchanged})
		case hasKeyInData1 && hasKeyInData2:
			diff = append(diff, Diff{key: k, value: value, newValue: newValue, status: Changed})
		case hasKeyInData1 && !hasKeyInData2:
			diff = append(diff, Diff{key: k, value: value, status: Deleted})
		case !hasKeyInData1 && hasKeyInData2:
			diff = append(diff, Diff{key: k, newValue: newValue, status: Added})
		}
	}

	return diff
}

func fmtDiffStr(key string, value any, sign string) string {
	return fmt.Sprintf("  %s %s: %v\n", sign, key, value)
}

func fmtDiff(diff []Diff) string {
	var b strings.Builder
	b.WriteString("{\n")

	for _, d := range diff {
		switch d.status {
		case Added:
			b.WriteString(fmtDiffStr(d.key, d.newValue, "+"))
		case Deleted:
			b.WriteString(fmtDiffStr(d.key, d.value, "-"))
		case Unchanged:
			b.WriteString(fmtDiffStr(d.key, d.value, " "))
		case Changed:
			b.WriteString(fmtDiffStr(d.key, d.value, "-"))
			b.WriteString(fmtDiffStr(d.key, d.newValue, "+"))
		}
	}

	b.WriteString("}")

	return b.String()
}

func GenDiff(path1, path2, format string) (string, error) {
	parsedData1, err := parsers.ParseFile(path1)
	if err != nil {
		return "", err
	}
	parsedData2, err := parsers.ParseFile(path2)
	if err != nil {
		return "", err
	}

	diff := buildDiff(parsedData1, parsedData2)

	return fmtDiff(diff), nil
}
