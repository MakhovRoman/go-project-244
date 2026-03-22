package formatters

import (
	"code/internal/compare"
	"code/internal/utils"
	"fmt"
	"maps"
	"slices"
	"strings"
)

var indentIdx = 4

func Stylish(diff compare.DiffMap) string {
	return "{\n" + stylish(diff, 1) + "}"
}

func stylish(diff compare.DiffMap, depth int) string {
	var result strings.Builder
	sortedKeys := utils.SortKeys(diff)

	for _, k := range sortedKeys {
		v := diff[k]
		indent := strings.Repeat(" ", indentIdx*depth-2)

		if len(v.Children) > 0 {
			fmt.Fprintf(&result, indent+"  %s: {\n", k)
			result.WriteString(stylish(v.Children, depth+1))
			fmt.Fprint(&result, indent+"  }\n")
			continue
		}
		switch v.Code {
		case compare.CodeAdded:
			fmt.Fprintf(&result, indent+"+ %s: %s\n", k, stylishFormatter(v.NewValue, depth))
		case compare.CodeRemoved:
			fmt.Fprintf(&result, indent+"- %s: %s\n", k, stylishFormatter(v.OldValue, depth))
		case compare.CodeChanged:
			fmt.Fprintf(&result, indent+"- %s: %s\n", k, stylishFormatter(v.OldValue, depth))
			fmt.Fprintf(&result, indent+"+ %s: %s\n", k, stylishFormatter(v.NewValue, depth))
		default:
			fmt.Fprintf(&result, indent+"  %s: %s\n", k, stylishFormatter(v.OldValue, depth))
		}
	}

	return result.String()
}

func stylishFormatter(v any, depth int) string {
	if v == nil {
		return "null"
	}
	if m, ok := v.(map[string]any); ok {
		return formatRawMap(m, depth)
	}
	return fmt.Sprintf("%v", v)
}

func formatRawMap(m map[string]any, depth int) string {
	var str strings.Builder
	innerIndent := strings.Repeat(" ", indentIdx*(depth+1)-2)
	closingIndent := strings.Repeat(" ", indentIdx*depth-2)
	sortedKeys := slices.Sorted(maps.Keys(m))

	str.WriteString("{\n")
	for _, k := range sortedKeys {
		v := m[k]
		if dm, ok := v.(map[string]any); ok {
			fmt.Fprintf(&str, innerIndent+"  %s: %s\n", k, formatRawMap(dm, depth+1))
		} else {
			fmt.Fprintf(&str, innerIndent+"  %s: %s\n", k, fmt.Sprintf("%v", v))
		}
	}
	fmt.Fprint(&str, closingIndent+"  }")
	return str.String()
}
