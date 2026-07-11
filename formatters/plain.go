package formatters

import (
	"fmt"
	"strings"

	"code/diff"
	"code/parsers"
)

type plainFormatter struct{}

func (p *plainFormatter) fmtValue(value any) string {
	switch value.(type) {
	case parsers.ParsedData:
		return "[complex value]"
	case nil:
		return "null"
	case string:
		return fmt.Sprintf("'%s'", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}

func (p *plainFormatter) Format(diffNodes []diff.Diff) string {
	return p.walk(diffNodes, "")
}

func (p *plainFormatter) walk(nodes []diff.Diff, path string) string {
	var b strings.Builder

	for _, d := range nodes {
		propPath := d.Key
		if path != "" {
			propPath = path + "." + d.Key
		}

		switch d.Type {
		case diff.Added:
			fmt.Fprintf(&b, "Property '%s' was added with value: %s\n", propPath, p.fmtValue(d.Value2))
		case diff.Deleted:
			fmt.Fprintf(&b, "Property '%s' was removed\n", propPath)
		case diff.Unchanged:
			continue
		case diff.Changed:
			fmt.Fprintf(&b, "Property '%s' was updated. From %s to %s\n", propPath, p.fmtValue(d.Value1), p.fmtValue(d.Value2))
		case diff.Nested:
			fmt.Fprintln(&b, p.walk(d.Children, propPath))
		}
	}

	return strings.TrimSpace(b.String())
}
