package formatters

import (
	"encoding/json"

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

	data, _ := json.MarshalIndent(root, "", "  ")

	return string(data)
}
