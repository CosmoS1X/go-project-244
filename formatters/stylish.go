package formatters

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/diff"
	"code/parsers"
)

func makeIndent(depth, offset int) string {
	indentSym := " "
	indentSize := 4

	return strings.Repeat(indentSym, depth*indentSize-offset)
}

func fmtValue(value any, depth int) string {
	if value == nil {
		return "null"
	}

	m, ok := value.(parsers.ParsedData)
	if !ok {
		return fmt.Sprintf("%v", value)
	}

	var b strings.Builder
	b.WriteString("{\n")

	keys := slices.Sorted(maps.Keys(m))
	for _, k := range keys {
		fmt.Fprintf(&b, "%s%s: %s\n", makeIndent(depth+1, 0), k, fmtValue(m[k], depth+1))
	}

	fmt.Fprintf(&b, "%s}", makeIndent(depth, 0))

	return b.String()
}

func FmtStylish(diffNodes []diff.Diff) string {
	var iter func(nodes []diff.Diff, depth int) string

	iter = func(nodes []diff.Diff, depth int) string {
		var b strings.Builder
		b.WriteString("{\n")
		indent := makeIndent(depth, 2)

		for _, d := range nodes {
			switch d.Status {
			case diff.Added:
				fmt.Fprintf(&b, "%s+ %s: %s\n", indent, d.Key, fmtValue(d.NewValue, depth))
			case diff.Deleted:
				fmt.Fprintf(&b, "%s- %s: %s\n", indent, d.Key, fmtValue(d.Value, depth))
			case diff.Unchanged:
				fmt.Fprintf(&b, "%s  %s: %s\n", indent, d.Key, fmtValue(d.Value, depth))
			case diff.Changed:
				fmt.Fprintf(&b, "%s- %s: %s\n", indent, d.Key, fmtValue(d.Value, depth))
				fmt.Fprintf(&b, "%s+ %s: %s\n", indent, d.Key, fmtValue(d.NewValue, depth))
			case diff.Nested:
				fmt.Fprintf(&b, "%s  %s: %v\n", indent, d.Key, iter(d.Children, depth+1))
			}
		}

		fmt.Fprintf(&b, "%s}", makeIndent(depth, 4))

		return b.String()
	}

	return iter(diffNodes, 1)
}
