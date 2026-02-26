package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Parse(path string) (map[string]any, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	separated := strings.Split(info.Name(), ".")
	extension := separated[len(separated)-1]

	if extension != "json" {
		return nil, errors.New("unsupported format")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid json: %w", err)
	}
	return result, nil
}
