package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Parse(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)
	if ext != ".json" {
		return nil, errors.New("unsupported format")
	}

	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid json: %w", err)
	}
	return result, nil
}
