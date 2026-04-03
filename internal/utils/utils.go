package utils

import (
	"code/internal/compare"
	"maps"
	"slices"
)

// SortKeys возвращает отсортированный список ключей из DiffMap.
func SortKeys(list compare.DiffMap) []string {
	return slices.Sorted(maps.Keys(list))
}
