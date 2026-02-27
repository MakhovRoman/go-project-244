package compare

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func BuildDiff(d1, d2 map[string]any) string {
	keysMap := make(map[string]struct{})

	for key := range d1 {
		keysMap[key] = struct{}{}
	}
	for key := range d2 {
		keysMap[key] = struct{}{}
	}

	sortedKeysList := slices.Collect(maps.Keys(keysMap))
	slices.Sort(sortedKeysList)

	var result strings.Builder

	for _, k := range sortedKeysList {
		v1, ok1 := d1[k]
		v2, ok2 := d2[k]

		switch {
		case !ok1 && ok2:
			result.WriteString(fmt.Sprintf("+ %s: %v\n", k, v2))
		case ok1 && !ok2:
			result.WriteString(fmt.Sprintf("- %s: %v\n", k, v1))
		case v1 != v2:
			result.WriteString(fmt.Sprintf("- %s: %v\n", k, v1))
			result.WriteString(fmt.Sprintf("+ %s: %v\n", k, v2))
		default:
			result.WriteString(fmt.Sprintf("  %s: %v\n", k, v1))
		}
	}

	return result.String()
}
