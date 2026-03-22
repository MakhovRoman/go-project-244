package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Want struct {
	Flat string
	Deep string
}

var stylishWant = Want{
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

var plainWant = Want{
	Flat: "Property 'follow' was removed\n" +
		"Property 'proxy' was removed\n" +
		"Property 'timeout' was updated. From 50 to 20\n" +
		"Property 'verbose' was added with value: true\n",
	Deep: "Property 'common.follow' was added with value: false\n" +
		"Property 'common.setting2' was removed\n" +
		"Property 'common.setting3' was updated. From true to null\n" +
		"Property 'common.setting4' was added with value: 'blah blah'\n" +
		"Property 'common.setting5' was added with value: [complex value]\n" +
		"Property 'common.setting6.doge.wow' was updated. From '' to 'so much'\n" +
		"Property 'common.setting6.ops' was added with value: 'vops'\n" +
		"Property 'group1.baz' was updated. From 'bas' to 'bars'\n" +
		"Property 'group1.nest' was updated. From [complex value] to 'str'\n" +
		"Property 'group2' was removed\n" +
		"Property 'group3' was added with value: [complex value]\n",
}

func TestGenDiff(t *testing.T) {
	testData := []struct {
		name   string
		path1  string
		path2  string
		want   string
		format string
	}{
		{
			name:   "flatJSONStylish",
			path1:  "testdata/fixture/file1.json",
			path2:  "testdata/fixture/file2.json",
			want:   stylishWant.Flat,
			format: "stylish",
		},
		{
			name:   "deepJSONStylish",
			path1:  "testdata/fixture/deep1.json",
			path2:  "testdata/fixture/deep2.json",
			want:   stylishWant.Deep,
			format: "stylish",
		},
		{
			name:   "flatYAMLStylish",
			path1:  "testdata/fixture/file1.yaml",
			path2:  "testdata/fixture/file2.yaml",
			want:   stylishWant.Flat,
			format: "stylish",
		},
		{
			name:   "deepYAMLStylish",
			path1:  "testdata/fixture/deep1.yaml",
			path2:  "testdata/fixture/deep2.yaml",
			want:   stylishWant.Deep,
			format: "stylish",
		},
		{
			name:   "flatJSONPlain",
			path1:  "testdata/fixture/file1.json",
			path2:  "testdata/fixture/file2.json",
			want:   plainWant.Flat,
			format: "plain",
		},
		{
			name:   "deepJSONPlain",
			path1:  "testdata/fixture/deep1.json",
			path2:  "testdata/fixture/deep2.json",
			want:   plainWant.Deep,
			format: "plain",
		},
		{
			name:   "flatYAMLPlain",
			path1:  "testdata/fixture/file1.yaml",
			path2:  "testdata/fixture/file2.yaml",
			want:   plainWant.Flat,
			format: "plain",
		},
		{
			name:   "deepYAMLPlain",
			path1:  "testdata/fixture/deep1.yaml",
			path2:  "testdata/fixture/deep2.yaml",
			want:   plainWant.Deep,
			format: "plain",
		},
	}

	for _, tt := range testData {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := GenDiff(tc.path1, tc.path2, tc.format)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}
