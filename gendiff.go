package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func validateSupportedFile(path string) error {
	supportedExtensions := []string{".json", ".yml", ".yaml"}

	base := filepath.Base(path)
	ext := filepath.Ext(path)

	if !slices.Contains(supportedExtensions, ext) {
		return fmt.Errorf("extension '%s' of file '%s' is not supported", ext, base)
	}

	return nil
}

type ParsedData map[string]any

func parse(data []byte) (ParsedData, error) {
	parsedData := make(ParsedData)
	err := json.Unmarshal(data, &parsedData)
	if err != nil {
		return nil, err
	}

	return parsedData, nil
}

func readAndParseFile(path string) (ParsedData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parse(data)
}

func GenDiff(path1, path2, format string) (string, error) {
	if err := validateSupportedFile(path1); err != nil {
		return "", err
	}
	if err := validateSupportedFile(path2); err != nil {
		return "", err
	}

	parsedData1, err := readAndParseFile(path1)
	if err != nil {
		return "", err
	}

	parsedData2, err := readAndParseFile((path2))
	if err != nil {
		return "", err
	}

	fmt.Println("Parsed data 1:", parsedData1)
	fmt.Println("Parsed data 2:", parsedData2)
	fmt.Println("Selected format:", format)

	return "Data parsed successfully", nil
}
