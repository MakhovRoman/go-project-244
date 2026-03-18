package code

import (
	"code/internal/compare"
	"code/internal/parsers"
	"path/filepath"
)

func GenDiff(path1, path2 string) (string, error) {
	normalizedPath1, err := normalizePath(path1)
	if err != nil {
		return "", err
	}

	normalizedPath2, err := normalizePath(path2)
	if err != nil {
		return "", err
	}

	parsed1, err := parsers.Parse(normalizedPath1)
	if err != nil {
		return "", err
	}

	parsed2, err := parsers.Parse(normalizedPath2)
	if err != nil {
		return "", err
	}

	result := compare.BuildDiff(parsed1, parsed2)

	return result, nil
}

func normalizePath(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Clean(abs), nil
}
