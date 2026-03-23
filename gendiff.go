package code

import (
	"code/internal/compare"
	"code/internal/formatters"
	"code/internal/parser"
	"path/filepath"
)

func GenDiff(path1, path2, format string) (string, error) {
	normalizedPath1, err := normalizePath(path1)
	if err != nil {
		return "", err
	}

	normalizedPath2, err := normalizePath(path2)
	if err != nil {
		return "", err
	}

	parsed1, err := parser.Parse(normalizedPath1)
	if err != nil {
		return "", err
	}

	parsed2, err := parser.Parse(normalizedPath2)
	if err != nil {
		return "", err
	}

	diff := compare.BuildDiff(parsed1, parsed2)
	result, err := formatters.GetFormattedDif(format, diff)

	return result, err
}

func normalizePath(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Clean(abs), nil
}
