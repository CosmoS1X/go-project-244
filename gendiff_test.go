package code

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testdataDir   = "testdata"
	jsonFile1     = "file1.json"
	jsonFile2     = "file2.json"
	yamlFile1     = "file1.yml"
	yamlFile2     = "file2.yaml"
	stylishFile   = "stylish.txt"
	plainFile     = "plain.txt"
	formatStylish = "stylish"
	formatPlain   = "plain"
)

func TestGendiff(t *testing.T) {
	cases := []struct {
		name         string
		path1, path2 string
		format       string
		expected     string
	}{
		{
			name:     "compare json files with stylish format",
			path1:    filepath.Join(testdataDir, jsonFile1),
			path2:    filepath.Join(testdataDir, jsonFile2),
			format:   formatStylish,
			expected: filepath.Join(testdataDir, stylishFile),
		},
		{
			name:     "compare yaml files with stylish format",
			path1:    filepath.Join(testdataDir, yamlFile1),
			path2:    filepath.Join(testdataDir, yamlFile2),
			format:   formatStylish,
			expected: filepath.Join(testdataDir, stylishFile),
		},
		{
			name:     "compare json files with plain format",
			path1:    filepath.Join(testdataDir, jsonFile1),
			path2:    filepath.Join(testdataDir, jsonFile2),
			format:   formatPlain,
			expected: filepath.Join(testdataDir, plainFile),
		},
		{
			name:     "compare yaml files with plain format",
			path1:    filepath.Join(testdataDir, yamlFile1),
			path2:    filepath.Join(testdataDir, yamlFile2),
			format:   formatPlain,
			expected: filepath.Join(testdataDir, plainFile),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			data, err := os.ReadFile(c.expected)
			if err != nil {
				t.Fatal(err)
			}

			want := strings.TrimSpace(string(data))
			got, err := GenDiff(c.path1, c.path2, c.format)

			assert.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}
