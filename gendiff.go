package code

import (
	"encoding/json"
	"fmt"
	"maps"
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

func getCommonKeys(data1, data2 ParsedData) []string {
	uniqMap := make(map[string]struct{}, len(data1)+len(data2))

	for k := range data1 {
		uniqMap[k] = struct{}{}
	}
	for k := range data2 {
		uniqMap[k] = struct{}{}
	}

	return slices.Sorted(maps.Keys(uniqMap))
}

type Diff struct {
	key      string
	value    any
	newValue any
	status   string
}

func makeDiff(data1, data2 ParsedData) []Diff {
	commonKeys := getCommonKeys(data1, data2)
	diff := make([]Diff, 0, len(commonKeys))

	for _, k := range commonKeys {
		value, hasKeyInData1 := data1[k]
		newValue, hasKeyInData2 := data2[k]

		switch {
		case value == newValue:
			diff = append(diff, Diff{key: k, value: value, status: "unchanged"})
		case hasKeyInData1 && hasKeyInData2:
			diff = append(diff, Diff{key: k, value: value, newValue: newValue, status: "changed"})
		case hasKeyInData1 && !hasKeyInData2:
			diff = append(diff, Diff{key: k, value: value, status: "deleted"})
		case !hasKeyInData1 && hasKeyInData2:
			diff = append(diff, Diff{key: k, newValue: newValue, status: "added"})
		}
	}

	return diff
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

	diff := makeDiff(parsedData1, parsedData2)

	return fmt.Sprintf("Diff created: %v", diff), nil
}
