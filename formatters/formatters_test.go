package formatters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CosmoS1X/differ/diff"
	"github.com/CosmoS1X/differ/parsers"
)

var testDiffNodes = []diff.Diff{
	{
		Key: "common",
		Children: []diff.Diff{
			{Key: "setting1", Value1: "Value 1", Type: diff.Unchanged},
			{Key: "setting2", Value1: 200, Type: diff.Deleted},
			{Key: "setting3", Value1: true, Value2: nil, Type: diff.Changed},
			{Key: "setting4", Value2: "blah blah", Type: diff.Added},
			{Key: "setting5", Value2: parsers.ParsedData{"key5": "value5"}, Type: diff.Added},
		},
		Type: diff.Nested,
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

func TestFormat_JSONFormatter(t *testing.T) {
	got, err := Format(testDiffNodes, "json")

	assert.NoError(t, err)
	assert.Contains(t, got, `"key": ""`)
	assert.Contains(t, got, `"type": "root"`)
	assert.Contains(t, got, `"children": [`)
	assert.Contains(t, got, `"key": "common"`)
	assert.Contains(t, got, `"type": "nested"`)
	assert.Contains(t, got, `"value1": "Value 1"`)
	assert.Contains(t, got, `"value2": "blah blah"`)
}

func TestJSONFormatter_ReturnsErrorForUnsupportedValue(t *testing.T) {
	formatter := &jsonFormatter{}
	_, err := formatter.Format([]diff.Diff{{Key: "broken", Value1: func() {}, Type: diff.Added}})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "marshal diff nodes")
}
