package formatters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code/diff"
	"code/parsers"
)

var testDiffNodes = []diff.Diff{
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

func TestFormat_ReturnsErrorForUnsupportedFormat(t *testing.T) {
	_, err := Format(testDiffNodes, "unknown")
	assert.EqualError(t, err, `unsupported format name: "unknown"`)
}

func TestFormat_StylishFormatter(t *testing.T) {
	got, err := Format(testDiffNodes, "stylish")

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

func TestFormat_PlainFormatter(t *testing.T) {
	got, err := Format(testDiffNodes, "plain")

	assert.NoError(t, err)
	assert.NotContains(t, got, "Property 'common.setting1'")
	assert.Contains(t, got, "Property 'common.setting2' was removed")
	assert.Contains(t, got, "Property 'common.setting3' was updated. From true to null")
	assert.Contains(t, got, "Property 'common.setting4' was added with value: 'blah blah'")
	assert.Contains(t, got, "Property 'common.setting5' was added with value: [complex value]")
}
