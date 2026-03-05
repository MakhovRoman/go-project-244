package parser

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var ErrUnsupportedFormat = errors.New("unsupported format")

func Parse(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)
	if ext != ".json" {
		return nil, ErrUnsupportedFormat
	}

	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
