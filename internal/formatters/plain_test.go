package formatters

import (
	"code/internal/compare"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlain(t *testing.T) {
	testData := []struct {
		name string
		diff compare.DiffMap
		want string
	}{
		{
			name: "empty",
			diff: compare.DiffMap{},
			want: "",
		},
		{
			name: "added scalar",
			diff: compare.DiffMap{
				"timeout": {Status: compare.CodeAdded, NewValue: 20},
			},
			want: "Property 'timeout' was added with value: 20",
		},
		{
			name: "added string",
			diff: compare.DiffMap{
				"host": {Status: compare.CodeAdded, NewValue: "hexlet.io"},
			},
			want: "Property 'host' was added with value: 'hexlet.io'",
		},
		{
			name: "added null",
			diff: compare.DiffMap{
				"setting3": {Status: compare.CodeAdded, NewValue: nil},
			},
			want: "Property 'setting3' was added with value: null",
		},
		{
			name: "added complex",
			diff: compare.DiffMap{
				"setting5": {Status: compare.CodeAdded, NewValue: map[string]any{"key5": "value5"}},
			},
			want: "Property 'setting5' was added with value: [complex value]",
		},
		{
			name: "removed",
			diff: compare.DiffMap{
				"proxy": {Status: compare.CodeRemoved, OldValue: "123.234.53.22"},
			},
			want: "Property 'proxy' was removed",
		},
		{
			name: "changed",
			diff: compare.DiffMap{
				"timeout": {Status: compare.CodeChanged, OldValue: 50, NewValue: 20},
			},
			want: "Property 'timeout' was updated. From 50 to 20",
		},
		{
			name: "changed complex to scalar",
			diff: compare.DiffMap{
				"nest": {Status: compare.CodeChanged, OldValue: map[string]any{"key": "value"}, NewValue: "str"},
			},
			want: "Property 'nest' was updated. From [complex value] to 'str'",
		},
		{
			name: "unchanged skipped",
			diff: compare.DiffMap{
				"host": {Status: compare.CodeUnchanged, OldValue: "hexlet.io"},
			},
			want: "",
		},
		{
			name: "nested path",
			diff: compare.DiffMap{
				"common": {
					Status: compare.CodeUnchanged,
					Children: compare.DiffMap{
						"follow":   {Status: compare.CodeAdded, NewValue: false},
						"setting2": {Status: compare.CodeRemoved, OldValue: 200},
					},
				},
			},
			want: "Property 'common.follow' was added with value: false\n" +
				"Property 'common.setting2' was removed",
		},
		{
			name: "deeply nested path",
			diff: compare.DiffMap{
				"common": {
					Status: compare.CodeUnchanged,
					Children: compare.DiffMap{
						"setting6": {
							Status: compare.CodeUnchanged,
							Children: compare.DiffMap{
								"ops": {Status: compare.CodeAdded, NewValue: "vops"},
							},
						},
					},
				},
			},
			want: "Property 'common.setting6.ops' was added with value: 'vops'",
		},
	}

	for _, tt := range testData {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Plain(tc.diff)
			assert.Equal(t, tc.want, got)
		})
	}
}
