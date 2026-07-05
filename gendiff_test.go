package code

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGendiff_Stylish(t *testing.T) {
	data, err := os.ReadFile("testdata/stylish.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := strings.TrimSpace(string(data))
	got, err := GenDiff("testdata/file1.json", "testdata/file2.json", "stylish")

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestGenDiff_UnsupportedExt(t *testing.T) {
	_, err := GenDiff("testdata/file1.json", "testdata/unsupported.txt", "stylish")
	assert.Error(t, err)

	_, err = GenDiff("testdata/unsupported.txt", "testdata/file2.json", "stylish")
	assert.Error(t, err)
}

func TestGenDiff_NonExistentFile(t *testing.T) {
	_, err := GenDiff("unknown.json", "testdata/file2.json", "stylish")
	assert.Error(t, err)

	_, err = GenDiff("testdata/file1.json", "unknown.json", "stylish")
	assert.Error(t, err)
}

func TestGendiff_ParseError(t *testing.T) {
	_, err := GenDiff("testdata/broken.json", "testdata/file2.json", "stylish")
	assert.Error(t, err)
}
