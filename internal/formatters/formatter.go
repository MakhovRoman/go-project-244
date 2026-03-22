package formatters

import "code/internal/compare"

func GetFormattedDif(format string, diff compare.DiffMap) string {
	switch format {
	case "stylish":
		return Stylish(diff)
	case "plain":
		return Plain(diff)
	default:
		return ""
	}
}
