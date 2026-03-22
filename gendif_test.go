package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var want = struct {
	Flat string
	Deep string
}{
	Flat: "{\n" +
		"  - follow: false\n" +
		"    host: hexlet.io\n" +
		"  - proxy: 123.234.53.22\n" +
		"  - timeout: 50\n" +
		"  + timeout: 20\n" +
		"  + verbose: true\n" +
		"}",
	Deep: "{\n" +
		"    common: {\n" +
		"      + follow: false\n" +
		"        setting1: Value 1\n" +
		"      - setting2: 200\n" +
		"      - setting3: true\n" +
		"      + setting3: null\n" +
		"      + setting4: blah blah\n" +
		"      + setting5: {\n" +
		"            key5: value5\n" +
		"        }\n" +
		"        setting6: {\n" +
		"            doge: {\n" +
		"              - wow: \n" +
		"              + wow: so much\n" +
		"            }\n" +
		"            key: value\n" +
		"          + ops: vops\n" +
		"        }\n" +
		"    }\n" +
		"    group1: {\n" +
		"      - baz: bas\n" +
		"      + baz: bars\n" +
		"        foo: bar\n" +
		"      - nest: {\n" +
		"            key: value\n" +
		"        }\n" +
		"      + nest: str\n" +
		"    }\n" +
		"  - group2: {\n" +
		"        abc: 12345\n" +
		"        deep: {\n" +
		"            id: 45\n" +
		"        }\n" +
		"    }\n" +
		"  + group3: {\n" +
		"        deep: {\n" +
		"            id: {\n" +
		"                number: 45\n" +
		"            }\n" +
		"        }\n" +
		"        fee: 100500\n" +
		"    }\n" +
		"}",
}

func TestGenDiff(t *testing.T) {
	testData := []struct {
		name  string
		path1 string
		path2 string
		want  string
	}{
		{
			name:  "flatJSONCompare",
			path1: "testdata/fixture/file1.json",
			path2: "testdata/fixture/file2.json",
			want:  want.Flat,
		},
		{
			name:  "deepJSONCompare",
			path1: "testdata/fixture/deep1.json",
			path2: "testdata/fixture/deep2.json",
			want:  want.Deep,
		},
		{
			name:  "flatYAMLCompare",
			path1: "testdata/fixture/file1.yaml",
			path2: "testdata/fixture/file2.yaml",
			want:  want.Flat,
		},
		{
			name:  "deepYAMLCompare",
			path1: "testdata/fixture/deep1.yaml",
			path2: "testdata/fixture/deep2.yaml",
			want:  want.Deep,
		},
	}

	for _, tt := range testData {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := GenDiff(tc.path1, tc.path2, "stylish")
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}
