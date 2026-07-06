package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	cases := []struct {
		name    string
		path    string
		want    ParsedData
		wantErr bool
	}{
		{
			name: "valid JSON",
			path: "../testdata/file1.json",
			want: ParsedData{
				"host":    "hexlet.io",
				"timeout": float64(50), // JSON numbers are unmarshaled as float64
				"proxy":   "123.234.53.22",
				"follow":  false,
			},
		},
		{
			name: "valid YAML",
			path: "../testdata/file2.yaml",
			want: ParsedData{
				"timeout": int(20), // YAML numbers are unmarshaled as int
				"verbose": true,
				"host":    "hexlet.io",
			},
		},
		{
			name:    "broken file",
			path:    "../testdata/broken.json",
			wantErr: true,
		},
		{
			name:    "non-existent path",
			path:    "../testdata/unknown.json",
			wantErr: true,
		},
		{
			name:    "unsupported file extension",
			path:    "../testdata/unsupported.txt",
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseFile(c.path)
			if c.wantErr {
				assert.Error(t, err)
				assert.Empty(t, got)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}
