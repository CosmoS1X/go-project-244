//nolint:goconst // repeated literals in test fixtures are intentional for readability
package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code/parsers"
)

func buildTestData() (data1, data2 parsers.ParsedData) {
	data1 = parsers.ParsedData{
		"common": parsers.ParsedData{
			"setting1": "Value 1",
			"setting2": 200,
			"setting3": true,
			"setting6": parsers.ParsedData{
				"key": "value",
				"deep": parsers.ParsedData{
					"wow": "",
				},
			},
		},
		"slice1": []any{parsers.ParsedData{"id": 1}, parsers.ParsedData{"id": 2}},
		"slice2": []any{[]any{1, 2}},
	}
	data2 = parsers.ParsedData{
		"common": parsers.ParsedData{
			"follow":   false,
			"setting1": "Value 1",
			"setting3": nil,
			"setting4": "blah blah",
			"setting5": parsers.ParsedData{
				"key5": "value5",
			},
			"setting6": parsers.ParsedData{
				"key": "value",
				"ops": "vops",
				"deep": parsers.ParsedData{
					"wow": "so much",
				},
			},
		},
		"slice1": []any{parsers.ParsedData{"id": 1}, parsers.ParsedData{"id": 3}},
		"slice2": []any{[]any{1, 2}, 3},
	}

	return data1, data2
}

func expectedDiff() []Diff {
	return []Diff{
		{
			Key:    "common",
			Status: Nested,
			Children: []Diff{
				{Key: "follow", Status: Added, NewValue: false},
				{Key: "setting1", Status: Unchanged, Value: "Value 1"},
				{Key: "setting2", Status: Deleted, Value: 200},
				{Key: "setting3", Status: Changed, Value: true, NewValue: nil},
				{Key: "setting4", Status: Added, NewValue: "blah blah"},
				{Key: "setting5", Status: Added, NewValue: parsers.ParsedData{"key5": "value5"}},
				{
					Key:    "setting6",
					Status: Nested,
					Children: []Diff{
						{Key: "deep", Status: Nested, Children: []Diff{{Key: "wow", Status: Changed, Value: "", NewValue: "so much"}}},
						{Key: "key", Status: Unchanged, Value: "value"},
						{Key: "ops", Status: Added, NewValue: "vops"},
					},
				},
			},
		},
		{
			Key:      "slice1",
			Status:   Changed,
			Value:    []any{parsers.ParsedData{"id": 1}, parsers.ParsedData{"id": 2}},
			NewValue: []any{parsers.ParsedData{"id": 1}, parsers.ParsedData{"id": 3}},
		},
		{
			Key:      "slice2",
			Status:   Changed,
			Value:    []any{[]any{1, 2}},
			NewValue: []any{[]any{1, 2}, 3},
		},
	}
}

func TestBuild_ProducesExpectedDiff(t *testing.T) {
	data1, data2 := buildTestData()
	got := Build(data1, data2)

	assert.Equal(t, expectedDiff(), got)
}

func TestIsEqual_ReturnsFalseForMismatchedContainerTypes(t *testing.T) {
	cases := []struct {
		name  string
		left  any
		right any
	}{
		{name: "slice type mismatch", left: []any{1, 2}, right: []string{"1", "2"}},
		{name: "map type mismatch", left: map[string]any{"a": 1}, right: map[string]string{"a": "1"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.False(t, isEqual(c.left, c.right))
		})
	}
}
