package formatters

import (
	"fmt"
	"strings"

	"code/diff"
	"code/parsers"
)

func fmtPlainVal(value any) string {
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

func FmtPlain(diffNodes []diff.Diff) string {
	var iter func(nodes []diff.Diff, path string) string

	iter = func(nodes []diff.Diff, path string) string {
		var b strings.Builder

		for _, d := range nodes {
			propPath := strings.TrimLeft(path+"."+d.Key, ".")

			switch d.Status {
			case diff.Added:
				fmt.Fprintf(&b, "Property '%s' was added with value: %s\n", propPath, fmtPlainVal(d.NewValue))
			case diff.Deleted:
				fmt.Fprintf(&b, "Property '%s' was removed\n", propPath)
			case diff.Unchanged:
				continue
			case diff.Changed:
				fmt.Fprintf(&b, "Property '%s' was updated. From %s to %s\n", propPath, fmtPlainVal(d.Value), fmtPlainVal(d.NewValue))
			case diff.Nested:
				fmt.Fprintln(&b, iter(d.Children, propPath))
			}
		}

		return strings.TrimSpace(b.String())
	}

	return iter(diffNodes, "")
}
