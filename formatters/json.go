package formatters

import (
	"encoding/json"
	"strings"

	"code/diff"
)

type jsonFormatter struct{}

type rootNode struct {
	Key      string      `json:"key"`
	Type     string      `json:"type"`
	Children []diff.Diff `json:"children"`
}

func (j *jsonFormatter) Format(diffNodes []diff.Diff) string {
	root := rootNode{
		Key:      "",
		Type:     "root",
		Children: diffNodes,
	}

	indentSym := " "
	indentSize := 2

	data, _ := json.MarshalIndent(root, "", strings.Repeat(indentSym, indentSize))

	return string(data)
}
