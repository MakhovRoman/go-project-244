package parser

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type FuncParser func(data []byte) (map[string]any, error)

var ErrUnsupportedFormat = errors.New("unsupported format")
var parsersMap = map[string]FuncParser{
	".json": jsonParser,
	".yaml": yamlParser,
	".yml":  yamlParser,
}

func Parse(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)
	if p, ok := parsersMap[ext]; ok {
		return p(data)
	}

	return nil, ErrUnsupportedFormat
}

func jsonParser(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func yamlParser(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
