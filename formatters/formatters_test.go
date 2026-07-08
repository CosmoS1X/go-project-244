package formatters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code/diff"
	"code/parsers"
)

func TestFormat_ReturnsErrorForUnsupportedFormat(t *testing.T) {
	_, err := Format(nil, "unknown")
	assert.EqualError(t, err, `unsupported format name: "unknown"`)
}

func TestFormat_StylishFormatter(t *testing.T) {
	diffNodes := []diff.Diff{
		{
			Key: "common",
			Children: []diff.Diff{
				{Key: "setting1", Value: "Value 1", Status: diff.Unchanged},
				{Key: "setting2", Value: 200, Status: diff.Deleted},
				{Key: "setting3", Value: true, NewValue: nil, Status: diff.Changed},
				{Key: "setting4", NewValue: "blah blah", Status: diff.Added},
				{Key: "setting5", NewValue: parsers.ParsedData{"key5": "value5"}, Status: diff.Added},
			},
			Status: diff.Nested,
		},
	}

	got, err := Format(diffNodes, "stylish")

	assert.NoError(t, err)
	assert.Contains(t, got, "  common: {")
	assert.Contains(t, got, "  setting1: Value 1")
	assert.Contains(t, got, "- setting2: 200")
	assert.Contains(t, got, "- setting3: true")
	assert.Contains(t, got, "+ setting3: null")
	assert.Contains(t, got, "+ setting4: blah blah")
	assert.Contains(t, got, "+ setting5: {")
	assert.Contains(t, got, "  key5: value5")
}
