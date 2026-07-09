package formatters

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/diff"
	"code/parsers"
)

type stylishFormatter struct{}

func makeIndent(depth, offset int) string {
	indentSym := " "
	indentSize := 4

	return strings.Repeat(indentSym, depth*indentSize-offset)
}

func (s *stylishFormatter) fmtValue(value any, depth int) string {
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
		fmt.Fprintf(&b, "%s%s: %s\n", makeIndent(depth+1, 0), k, s.fmtValue(m[k], depth+1))
	}

	fmt.Fprintf(&b, "%s}", makeIndent(depth, 0))

	return b.String()
}

func (s *stylishFormatter) Format(diffNodes []diff.Diff) string {
	return s.walk(diffNodes, 1)
}

func (s *stylishFormatter) walk(nodes []diff.Diff, depth int) string {
	var b strings.Builder
	b.WriteString("{\n")
	indent := makeIndent(depth, 2)

	for _, d := range nodes {
		switch d.Status {
		case diff.Added:
			fmt.Fprintf(&b, "%s+ %s: %s\n", indent, d.Key, s.fmtValue(d.NewValue, depth))
		case diff.Deleted:
			fmt.Fprintf(&b, "%s- %s: %s\n", indent, d.Key, s.fmtValue(d.Value, depth))
		case diff.Unchanged:
			fmt.Fprintf(&b, "%s  %s: %s\n", indent, d.Key, s.fmtValue(d.Value, depth))
		case diff.Changed:
			fmt.Fprintf(&b, "%s- %s: %s\n", indent, d.Key, s.fmtValue(d.Value, depth))
			fmt.Fprintf(&b, "%s+ %s: %s\n", indent, d.Key, s.fmtValue(d.NewValue, depth))
		case diff.Nested:
			fmt.Fprintf(&b, "%s  %s: %v\n", indent, d.Key, s.walk(d.Children, depth+1))
		}
	}

	fmt.Fprintf(&b, "%s}", makeIndent(depth, 4))

	return b.String()
}
