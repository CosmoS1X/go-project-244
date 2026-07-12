package formatters

import (
	"encoding/json"
	"fmt"

	"github.com/CosmoS1X/differ/diff"
)

type jsonFormatter struct{}

type rootNode struct {
	Key      string      `json:"key"`
	Type     string      `json:"type"`
	Children []diff.Diff `json:"children"`
}

func (j *jsonFormatter) Format(diffNodes []diff.Diff) (string, error) {
	root := rootNode{
		Key:      "",
		Type:     "root",
		Children: diffNodes,
	}

	data, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal diff nodes: %w", err)
	}

	return string(data), nil
}
