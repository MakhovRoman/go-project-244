package formatters

import (
	"code/internal/compare"
	"code/internal/utils"
	"encoding/json"
	"fmt"
)

type DiffNode struct {
	Status   string              `json:"status"`
	Value    *any                `json:"value,omitempty"`
	OldValue *any                `json:"oldValue,omitempty"`
	NewValue *any                `json:"newValue,omitempty"`
	Children map[string]DiffNode `json:"children,omitempty"`
}

func JSON(diff compare.DiffMap) (string, error) {
	node := buildJSON(diff)

	b, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return string(b), nil
}

func buildJSON(diff compare.DiffMap) map[string]DiffNode {
	result := make(map[string]DiffNode)

	for _, k := range utils.SortKeys(diff) {
		v := diff[k]
		node := DiffNode{Status: v.Status}

		if len(v.Children) > 0 {
			node.Children = buildJSON(v.Children)
		} else {
			switch v.Status {
			case compare.CodeUnchanged:
				if len(v.Children) == 0 {
					node.Value = Ptr(v.OldValue)
				}
			case compare.CodeAdded:
				node.NewValue = Ptr(v.NewValue)
			case compare.CodeChanged:
				node.OldValue = Ptr(v.OldValue)
				node.NewValue = Ptr(v.NewValue)
			case compare.CodeRemoved:
				node.OldValue = Ptr(v.OldValue)
			}
		}

		result[k] = node
	}

	return result
}

func Ptr(v any) *any {
	return &v
}
