package formatters

import (
	"code/internal/compare"
	"code/internal/utils"
	"fmt"
	"strings"
)

func Plain(diff compare.DiffMap) string {
	p := plain(diff, "")
	return strings.TrimRight(p, "\n")
}

func plain(diff compare.DiffMap, path string) string {
	var result strings.Builder
	sortedKeys := utils.SortKeys(diff)

	for _, k := range sortedKeys {
		v := diff[k]

		if len(v.Children) > 0 {
			result.WriteString(plain(v.Children, formatPath(path, k)))
			continue
		}

		switch v.Status {
		case compare.CodeAdded:
			fmt.Fprintf(&result, "Property '%s' was added with value: %v\n", formatPath(path, k), plainFormatter(v.NewValue))
		case compare.CodeRemoved:
			fmt.Fprintf(&result, "Property '%s' was removed\n", formatPath(path, k))
		case compare.CodeChanged:
			fmt.Fprintf(&result, "Property '%s' was updated. From %v to %v\n", formatPath(path, k), plainFormatter(v.OldValue), plainFormatter(v.NewValue))
		}
	}

	return result.String()
}

func plainFormatter(p any) string {
	if _, ok := p.(map[string]any); ok {
		return "[complex value]"
	}

	if s, ok := p.(string); ok {
		return fmt.Sprintf("'%s'", s)
	}

	if p == nil {
		return "null"
	}

	return fmt.Sprintf("%v", p)
}

func formatPath(path, key string) string {
	if len(path) == 0 {
		return key
	}
	return fmt.Sprintf("%s.%s", path, key)
}
