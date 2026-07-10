package code

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testdataDir   = "testdata"
	jsonFile1     = "file1.json"
	jsonFile2     = "file2.json"
	yamlFile1     = "file1.yml"
	yamlFile2     = "file2.yaml"
	stylishFile   = "stylish.txt"
	plainFile     = "plain.txt"
	jsonFile      = "json.txt"
	formatStylish = "stylish"
	formatPlain   = "plain"
	formatJSON    = "json"
)

func readExpectedOutput(t *testing.T, path string) string {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	return strings.TrimSpace(string(data))
}

func TestGendiff(t *testing.T) {
	cases := []struct {
		name         string
		path1, path2 string
		format       string
		expected     string
		wantErr      bool
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
		{
			name:     "compare json files with json format",
			path1:    filepath.Join(testdataDir, jsonFile1),
			path2:    filepath.Join(testdataDir, jsonFile2),
			format:   formatJSON,
			expected: filepath.Join(testdataDir, jsonFile),
		},
		{
			name:     "compare yaml files with json format",
			path1:    filepath.Join(testdataDir, yamlFile1),
			path2:    filepath.Join(testdataDir, yamlFile2),
			format:   formatJSON,
			expected: filepath.Join(testdataDir, jsonFile),
		},
		{
			name:     "compare json and yaml files with stylish format",
			path1:    filepath.Join(testdataDir, jsonFile1),
			path2:    filepath.Join(testdataDir, yamlFile2),
			format:   formatStylish,
			expected: filepath.Join(testdataDir, stylishFile),
		},
		{
			name:    "attempt to compare files with unsupported format",
			path1:   filepath.Join(testdataDir, jsonFile1),
			path2:   filepath.Join(testdataDir, jsonFile2),
			format:  "unsupported",
			wantErr: true,
		},
		{
			name:    "attempt to compare non-existent file",
			path1:   filepath.Join(testdataDir, "nonexistent.json"),
			path2:   filepath.Join(testdataDir, jsonFile2),
			format:  formatStylish,
			wantErr: true,
		},
		{
			name:    "attempt to compare files with unsupported extension",
			path1:   filepath.Join(testdataDir, jsonFile1),
			path2:   filepath.Join(testdataDir, "unsupported.txt"),
			format:  formatStylish,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			if c.wantErr {
				_, err := GenDiff(c.path1, c.path2, c.format)
				require.Error(t, err)
				return
			}

			want := readExpectedOutput(t, c.expected)
			got, err := GenDiff(c.path1, c.path2, c.format)

			require.NoError(t, err)
			require.Equal(t, want, got)
		})
	}
}
