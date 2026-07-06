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

	got, err = GenDiff("testdata/file1.yml", "testdata/file2.yaml", "stylish")
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
