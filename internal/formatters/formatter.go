package formatters

import (
	"code/internal/compare"
	"fmt"
)

// GetFormattedDif возвращает результат сравнения файлов в указанном формате.
func GetFormattedDif(format string, diff compare.DiffMap) (string, error) {
	switch format {
	case "stylish":
		return Stylish(diff), nil
	case "plain":
		return Plain(diff), nil
	case "json":
		return JSON(diff)

	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}
