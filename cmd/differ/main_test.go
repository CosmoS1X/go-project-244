package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	appname := "differ"
	root := filepath.Join("..", "..")
	testDirPath := filepath.Join(root, "testdata")
	filePath1 := filepath.Join(testDirPath, "file1.json")
	filePath2 := filepath.Join(testDirPath, "file2.json")

	cases := []struct {
		name     string
		args     []string
		expected string
		wantCode int
		wantErr  bool
	}{
		{
			name:     "compare files with stylish format by default",
			args:     []string{appname, filePath1, filePath2},
			expected: "result_stylish.txt",
		},
		{
			name:     "compare files with stylish format",
			args:     []string{appname, "-f", "stylish", filePath1, filePath2},
			expected: "result_stylish.txt",
		},
		{
			name:     "compare files with plain format using long flag",
			args:     []string{appname, "--format", "plain", filePath1, filePath2},
			expected: "result_plain.txt",
		},
		{
			name:     "compare files with json format",
			args:     []string{appname, "-f", "json", filePath1, filePath2},
			expected: "result_json.json",
		},
		{
			name:     "try to compare with broken file",
			args:     []string{appname, filePath1, filepath.Join(testDirPath, "broken.json")},
			wantCode: 1,
			wantErr:  true,
		},
		{
			name:     "try to compare files with unknown format",
			args:     []string{appname, "-f", "unknown", filePath1, filePath2},
			wantCode: 1,
			wantErr:  true,
		},
		{
			name:     "try to provide only one file path",
			args:     []string{appname, filePath1},
			wantCode: 1,
			wantErr:  true,
		},
		{
			name:     "try to provide more than two file paths",
			args:     []string{appname, filePath1, filePath2, filepath.Join(testDirPath, "file3.json")},
			wantCode: 1,
			wantErr:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			var stdout, stderr bytes.Buffer

			code := run(c.args, &stdout, &stderr)

			require.Equal(t, c.wantCode, code, "expected exit code")

			if c.wantErr {
				require.Empty(t, stdout.String(), "expected empty stdout")
				require.Contains(t, stderr.String(), "Error:", "expected error message in stderr")
				return
			}

			wantData, err := os.ReadFile(filepath.Join(testDirPath, c.expected))
			if err != nil {
				t.Fatal(err)
			}

			require.Empty(t, stderr.String(), "expected empty stderr")
			require.Equal(t, string(wantData), stdout.String(), "expected stdout")
		})
	}
}
