package formatters

import (
	"code/internal/compare"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStylish(t *testing.T) {
	testData := []struct {
		name string
		diff compare.DiffMap
		want string
	}{
		{
			name: "empty",
			diff: compare.DiffMap{},
			want: "{\n}",
		},
		{
			name: "flat unchanged",
			diff: compare.DiffMap{
				"host": {Code: compare.CodeUnchanged, OldValue: "hexlet.io"},
			},
			want: "{\n" +
				"    host: hexlet.io\n" +
				"}",
		},
		{
			name: "flat added",
			diff: compare.DiffMap{
				"timeout": {Code: compare.CodeAdded, NewValue: 20},
			},
			want: "{\n" +
				"  + timeout: 20\n" +
				"}",
		},
		{
			name: "flat removed",
			diff: compare.DiffMap{
				"proxy": {Code: compare.CodeRemoved, OldValue: "123.234.53.22"},
			},
			want: "{\n" +
				"  - proxy: 123.234.53.22\n" +
				"}",
		},
		{
			name: "flat changed",
			diff: compare.DiffMap{
				"timeout": {Code: compare.CodeChanged, OldValue: 50, NewValue: 20},
			},
			want: "{\n" +
				"  - timeout: 50\n" +
				"  + timeout: 20\n" +
				"}",
		},
		{
			name: "null value",
			diff: compare.DiffMap{
				"setting3": {Code: compare.CodeAdded, NewValue: nil},
			},
			want: "{\n" +
				"  + setting3: null\n" +
				"}",
		},
		{
			name: "nested unchanged",
			diff: compare.DiffMap{
				"group": {
					Code: compare.CodeUnchanged,
					Children: compare.DiffMap{
						"key": {Code: compare.CodeUnchanged, OldValue: "value"},
					},
				},
			},
			want: "{\n" +
				"    group: {\n" +
				"        key: value\n" +
				"    }\n" +
				"}",
		},
		{
			name: "nested with changes",
			diff: compare.DiffMap{
				"group": {
					Code: compare.CodeUnchanged,
					Children: compare.DiffMap{
						"foo": {Code: compare.CodeAdded, NewValue: "bar"},
						"baz": {Code: compare.CodeRemoved, OldValue: "bas"},
					},
				},
			},
			want: "{\n" +
				"    group: {\n" +
				"      - baz: bas\n" +
				"      + foo: bar\n" +
				"    }\n" +
				"}",
		},
		{
			name: "raw map value",
			diff: compare.DiffMap{
				"group": {Code: compare.CodeRemoved, OldValue: map[string]any{"key": "value"}},
			},
			want: "{\n" +
				"  - group: {\n" +
				"        key: value\n" +
				"    }\n" +
				"}",
		},
	}

	for _, tt := range testData {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Stylish(tc.diff)
			assert.Equal(t, tc.want, got)
		})
	}
}
