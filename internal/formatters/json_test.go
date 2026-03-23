package formatters

import (
	"code/internal/compare"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {
	testData := []struct {
		name string
		diff compare.DiffMap
		want map[string]any
	}{
		{
			name: "empty",
			diff: compare.DiffMap{},
			want: map[string]any{},
		},
		{
			name: "added scalar",
			diff: compare.DiffMap{
				"timeout": {Status: compare.CodeAdded, NewValue: 20},
			},
			want: map[string]any{
				"timeout": map[string]any{
					"status":   "added",
					"newValue": float64(20),
				},
			},
		},
		{
			name: "added null",
			diff: compare.DiffMap{
				"setting3": {Status: compare.CodeAdded, NewValue: nil},
			},
			want: map[string]any{
				"setting3": map[string]any{
					"status":   "added",
					"newValue": nil,
				},
			},
		},
		{
			name: "removed",
			diff: compare.DiffMap{
				"proxy": {Status: compare.CodeRemoved, OldValue: "123.234.53.22"},
			},
			want: map[string]any{
				"proxy": map[string]any{
					"status":   "removed",
					"oldValue": "123.234.53.22",
				},
			},
		},
		{
			name: "changed",
			diff: compare.DiffMap{
				"timeout": {Status: compare.CodeChanged, OldValue: 50, NewValue: 20},
			},
			want: map[string]any{
				"timeout": map[string]any{
					"status":   "changed",
					"oldValue": float64(50),
					"newValue": float64(20),
				},
			},
		},
		{
			name: "unchanged",
			diff: compare.DiffMap{
				"host": {Status: compare.CodeUnchanged, OldValue: "hexlet.io"},
			},
			want: map[string]any{
				"host": map[string]any{
					"status": "unchanged",
					"value":  "hexlet.io",
				},
			},
		},
		{
			name: "nested",
			diff: compare.DiffMap{
				"common": {
					Status: compare.CodeUnchanged,
					Children: compare.DiffMap{
						"follow":  {Status: compare.CodeAdded, NewValue: false},
						"setting": {Status: compare.CodeRemoved, OldValue: 200},
					},
				},
			},
			want: map[string]any{
				"common": map[string]any{
					"status": "unchanged",
					"children": map[string]any{
						"follow": map[string]any{
							"status":   "added",
							"newValue": false,
						},
						"setting": map[string]any{
							"status":   "removed",
							"oldValue": float64(200),
						},
					},
				},
			},
		},
	}

	for _, tt := range testData {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := JSON(tc.diff)
			require.NoError(t, err)

			var gotMap map[string]any
			err = json.Unmarshal([]byte(got), &gotMap)
			require.NoError(t, err)

			assert.Equal(t, tc.want, gotMap)
		})
	}
}
