package utils

import (
	"code/internal/compare"
	"maps"
	"slices"
)

func SortKeys(list compare.DiffMap) []string {
	return slices.Sorted(maps.Keys(list))
}
