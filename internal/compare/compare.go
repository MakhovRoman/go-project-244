package compare

import (
	"maps"
	"slices"
)

type DiffStruct struct {
	Code     int
	OldValue any
	NewValue any
	Children DiffMap
}
type DiffMap map[string]DiffStruct

const (
	CodeUnchanged = 0
	CodeAdded     = 1
	CodeRemoved   = 2
	CodeChanged   = 3
)

func BuildDiff(d1, d2 map[string]any) DiffMap {
	keysMap := make(map[string]struct{})

	for key := range d1 {
		keysMap[key] = struct{}{}
	}
	for key := range d2 {
		keysMap[key] = struct{}{}
	}

	sortedKeysList := slices.Collect(maps.Keys(keysMap))
	slices.Sort(sortedKeysList)

	diffMap := make(DiffMap)

	for _, k := range sortedKeysList {
		v1, ok1 := d1[k]
		v2, ok2 := d2[k]

		m1, mapOk1 := isMap(d1[k])
		m2, mapOk2 := isMap(d2[k])

		buf := DiffStruct{}

		switch {
		case mapOk1 && mapOk2:
			buf.Code = CodeUnchanged
			buf.Children = BuildDiff(m1, m2)
		case mapOk1 && !ok2 || ok1 && !ok2:
			buf.OldValue = v1
			buf.Code = CodeRemoved
		case !ok1 && mapOk2 || !ok1 && ok2:
			buf.NewValue = v2
			buf.Code = CodeAdded
		case v1 != v2:
			buf.OldValue = v1
			buf.NewValue = v2
			buf.Code = CodeChanged
		default:
			buf.OldValue = v1
		}

		diffMap[k] = buf
	}

	return diffMap
}

func isMap(value any) (map[string]any, bool) {
	if value == nil {
		return nil, false
	}

	obj, ok := value.(map[string]any)
	return obj, ok
}

//func isArray(value any) ([]any, bool) {
//	if value == nil {
//		return nil, false
//	}
//
//	arr, ok := value.([]any)
//	return arr, ok
//}
