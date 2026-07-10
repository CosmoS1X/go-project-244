package diff

import (
	"maps"
	"slices"

	"code/parsers"
)

type Diff struct {
	Key      string
	Value    any
	NewValue any
	Status   string
	Children []Diff
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
		value, hasKeyInData1 := data1[k]
		newValue, hasKeyInData2 := data2[k]
		hasKeyInBothData := hasKeyInData1 && hasKeyInData2
		map1, isMap1 := asParsedData(value)
		map2, isMap2 := asParsedData(newValue)
		hasChildren := isMap1 && isMap2

		switch {
		case hasChildren:
			diffNodes = append(diffNodes, Diff{Key: k, Children: Build(map1, map2), Status: Nested})
		case isEqual(value, newValue):
			diffNodes = append(diffNodes, Diff{Key: k, Value: value, Status: Unchanged})
		case hasKeyInBothData:
			diffNodes = append(diffNodes, Diff{Key: k, Value: value, NewValue: newValue, Status: Changed})
		case hasKeyInData1:
			diffNodes = append(diffNodes, Diff{Key: k, Value: value, Status: Deleted})
		case hasKeyInData2:
			diffNodes = append(diffNodes, Diff{Key: k, NewValue: newValue, Status: Added})
		}
	}

	return diffNodes
}
