package parsers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tempDir := t.TempDir()
	fakePath := "./fake_file.json"

	notJsonFile := filepath.Join(tempDir, "not_json.txt")
	if err := os.WriteFile(notJsonFile, []byte("{}"), 0644); err != nil {
		t.Fatal(err)
	}
	invalidJson := filepath.Join(tempDir, "invalid_json.json")
	if err := os.WriteFile(invalidJson, []byte("{"), 0644); err != nil {
		t.Fatal(err)
	}
	validJson := filepath.Join(tempDir, "valid_json.json")
	if err := os.WriteFile(validJson, []byte("{}"), 0644); err != nil {
		t.Fatal(err)
	}

	validYaml := filepath.Join(tempDir, "valid.yaml")
	if err := os.WriteFile(validYaml, []byte("key: value\n"), 0644); err != nil {
		t.Fatal(err)
	}
	validYml := filepath.Join(tempDir, "valid.yml")
	if err := os.WriteFile(validYml, []byte("key: value\n"), 0644); err != nil {
		t.Fatal(err)
	}
	invalidYaml := filepath.Join(tempDir, "invalid.yaml")
	if err := os.WriteFile(invalidYaml, []byte("key: :\n"), 0644); err != nil {
		t.Fatal(err)
	}

	testData := []struct {
		name  string
		path  string
		want  map[string]any
		check func(t *testing.T, err error)
	}{
		{
			name: "File not found",
			path: fakePath,
			want: map[string]any(nil),
			check: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, os.ErrNotExist)
			},
		},
		{
			name: "Not JSON file",
			path: notJsonFile,
			want: map[string]any(nil),
			check: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, ErrUnsupportedFormat)
			},
		},
		{
			name: "Invalid JSON file",
			path: invalidJson,
			want: map[string]any(nil),
			check: func(t *testing.T, err error) {
				var syntaxErr *json.SyntaxError
				assert.ErrorAs(t, err, &syntaxErr)
			},
		},
		{
			name: "Valid JSON file",
			path: validJson,
			want: map[string]any{},
			check: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "Valid YAML file (.yaml)",
			path: validYaml,
			want: map[string]any{"key": "value"},
			check: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "Valid YAML file (.yml)",
			path: validYml,
			want: map[string]any{"key": "value"},
			check: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "Invalid YAML file",
			path: invalidYaml,
			want: map[string]any(nil),
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tt := range testData {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := Parse(tc.path)

			if tc.check != nil {
				tc.check(t, err)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}
