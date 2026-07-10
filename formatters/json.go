package formatters

import (
	"encoding/json"
	"strings"

	"code/diff"
)

type jsonFormatter struct{}

func (j *jsonFormatter) Format(diffNodes []diff.Diff) string {
	indentSym := " "
	indentSize := 4

	data, _ := json.MarshalIndent(diffNodes, "", strings.Repeat(indentSym, indentSize))

	return string(data)
}
