package compare

import (
	"maps"
	"slices"
)

// DiffStruct содержит информацию о различии одного поля.
type DiffStruct struct {
	Status   string
	OldValue any
	NewValue any
	Children DiffMap
}

// DiffMap — результат сравнения двух файлов: ключ — имя поля, значение — DiffStruct.
type DiffMap map[string]DiffStruct

const (
	CodeUnchanged = "unchanged"
	CodeAdded     = "added"
	CodeRemoved   = "removed"
	CodeChanged   = "changed"
)

// BuildDiff строит карту, содержащую информацию по каждому полю
// в сравниваемых файлах. Ключом является имя поля, значением — структура DiffStruct,
// содержащая статус сравнения, значения в обоих файлах и список потомков (если поле является объектом).
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
			buf.Status = CodeUnchanged
			buf.Children = BuildDiff(m1, m2)
		case mapOk1 && !ok2 || ok1 && !ok2:
			buf.Status = CodeRemoved
			buf.OldValue = v1
		case !ok1 && mapOk2 || !ok1 && ok2:
			buf.Status = CodeAdded
			buf.NewValue = v2
		case v1 != v2:
			buf.Status = CodeChanged
			buf.OldValue = v1
			buf.NewValue = v2
		default:
			buf.Status = CodeUnchanged
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
