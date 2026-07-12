package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type (
	parseFn    func(in []byte, out any) error
	ParsedData = map[string]any
)

var registry = map[string]parseFn{
	".json": json.Unmarshal,
	".yml":  yaml.Unmarshal,
	".yaml": yaml.Unmarshal,
}

func parseRaw(data []byte, parser parseFn) (ParsedData, error) {
	var parsedData ParsedData
	if err := parser(data, &parsedData); err != nil {
		return nil, fmt.Errorf("failed to parse data: %w", err)
	}

	return parsedData, nil
}

func ParseFile(path string) (ParsedData, error) {
	ext := filepath.Ext(path)

	parser, ok := registry[ext]
	if !ok {
		return nil, fmt.Errorf("unsupported file extension %q", ext)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file error: %w", err)
	}

	return parseRaw(data, parser)
}
