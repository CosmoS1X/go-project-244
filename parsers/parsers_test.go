package parsers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile_JSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	json := `{"group1":{"baz":"bas","foo":"bar","nest":{"key": "value"}}}`

	err := os.WriteFile(path, []byte(json), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	want := ParsedData{
		"group1": ParsedData{
			"baz": "bas",
			"foo": "bar",
			"nest": ParsedData{
				"key": "value",
			},
		},
	}

	got, err := ParseFile(path)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestParseFile_YAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	yaml := "group2:\n  abc: 12345\n  deep:\n    id: 45\n"

	err := os.WriteFile(path, []byte(yaml), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	want := ParsedData{
		"group2": ParsedData{
			"abc": 12345,
			"deep": ParsedData{
				"id": 45,
			},
		},
	}

	got, err := ParseFile(path)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestParseFile_BrokenFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "broken.json")
	brokenJSON := `{"group1":{"baz":"bas","foo":"bar","nest":{"key": "value"}}` // Missing closing brace

	err := os.WriteFile(path, []byte(brokenJSON), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	got, err := ParseFile(path)

	assert.Error(t, err)
	assert.Empty(t, got)
}

func TestParseFile_NonExistentPath(t *testing.T) {
	path := "nonexistent.json"

	got, err := ParseFile(path)
	assert.Error(t, err)
	assert.Empty(t, got)
}

func TestParseFile_UnsupportedExtension(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.txt")

	err := os.WriteFile(path, []byte("some text"), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	got, err := ParseFile(path)

	assert.Error(t, err)
	assert.Empty(t, got)
}
