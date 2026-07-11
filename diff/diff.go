package diff

import (
	"maps"
	"slices"

	"code/parsers"
)

type Diff struct {
	Key      string `json:"key"`
	Type     string `json:"type"`
	Value1   any    `json:"value1,omitempty"`
	Value2   any    `json:"value2,omitempty"`
	Children []Diff `json:"children,omitempty"`
}

const (
	Added     = "added"
	Deleted   = "deleted"
	Unchanged = "unchanged"
	Changed   = "changed"
	Nested    = "nested"
)

func getCommonKeys(data1, data2 parsers.ParsedData) []string {
	uniqMap := make(map[string]struct{}, len(data1)+len(data2))

	for k := range data1 {
		uniqMap[k] = struct{}{}
	}
	for k := range data2 {
		uniqMap[k] = struct{}{}
	}

	return slices.Sorted(maps.Keys(uniqMap))
}

func asParsedData(v any) (parsers.ParsedData, bool) {
	if m, ok := v.(parsers.ParsedData); ok {
		return m, true
	}

	return nil, false
}

func isEqual(val1, val2 any) bool {
	switch v1 := val1.(type) {
	case []any:
		v2, ok := val2.([]any)
		if !ok {
			return false
		}
		return slices.EqualFunc(v1, v2, isEqual)
	case map[string]any:
		v2, ok := val2.(map[string]any)
		if !ok {
			return false
		}
		return maps.EqualFunc(v1, v2, isEqual)
	default:
		return val1 == val2
	}
}

func Build(data1, data2 parsers.ParsedData) []Diff {
	commonKeys := getCommonKeys(data1, data2)
	diffNodes := make([]Diff, 0, len(commonKeys))

	for _, k := range commonKeys {
		value1, hasKeyInData1 := data1[k]
		value2, hasKeyInData2 := data2[k]
		hasKeyInBothData := hasKeyInData1 && hasKeyInData2
		map1, isMap1 := asParsedData(value1)
		map2, isMap2 := asParsedData(value2)
		hasChildren := isMap1 && isMap2

		switch {
		case hasChildren:
			diffNodes = append(diffNodes, Diff{Key: k, Children: Build(map1, map2), Type: Nested})
		case isEqual(value1, value2):
			diffNodes = append(diffNodes, Diff{Key: k, Value1: value1, Type: Unchanged})
		case hasKeyInBothData:
			diffNodes = append(diffNodes, Diff{Key: k, Value1: value1, Value2: value2, Type: Changed})
		case hasKeyInData1:
			diffNodes = append(diffNodes, Diff{Key: k, Value1: value1, Type: Deleted})
		case hasKeyInData2:
			diffNodes = append(diffNodes, Diff{Key: k, Value2: value2, Type: Added})
		}
	}

	return diffNodes
}
