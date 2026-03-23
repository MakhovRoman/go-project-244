package code

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Flat string
	Deep string
}

var stylishWant = TestCase{
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

var plainWant = TestCase{
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

var jsonWant = TestCase{
	Flat: `{
  "follow": {
    "status": "removed",
    "oldValue": false
  },
  "host": {
    "status": "unchanged",
    "value": "hexlet.io"
  },
  "proxy": {
    "status": "removed",
    "oldValue": "123.234.53.22"
  },
  "timeout": {
    "status": "changed",
    "oldValue": 50,
    "newValue": 20
  },
  "verbose": {
    "status": "added",
    "newValue": true
  }
}`,
	Deep: `{
  "common": {
    "status": "unchanged",
    "children": {
      "follow": {"status": "added", "newValue": false},
      "setting1": {"status": "unchanged", "value": "Value 1"},
      "setting2": {"status": "removed", "oldValue": 200},
      "setting3": {"status": "changed", "oldValue": true, "newValue": null},
      "setting4": {"status": "added", "newValue": "blah blah"},
      "setting5": {"status": "added", "newValue": {"key5": "value5"}},
      "setting6": {
        "status": "unchanged",
        "children": {
          "doge": {
            "status": "unchanged",
            "children": {
              "wow": {"status": "changed", "oldValue": "", "newValue": "so much"}
            }
          },
          "key": {"status": "unchanged", "value": "value"},
          "ops": {"status": "added", "newValue": "vops"}
        }
      }
    }
  },
  "group1": {
    "status": "unchanged",
    "children": {
      "baz": {"status": "changed", "oldValue": "bas", "newValue": "bars"},
      "foo": {"status": "unchanged", "value": "bar"},
      "nest": {"status": "changed", "oldValue": {"key": "value"}, "newValue": "str"}
    }
  },
  "group2": {
    "status": "removed",
    "oldValue": {"abc": 12345, "deep": {"id": 45}}
  },
  "group3": {
    "status": "added",
    "newValue": {"deep": {"id": {"number": 45}}, "fee": 100500}
  }
}`,
}

var filePath = struct {
	JSON []TestCase
	Yaml []TestCase
}{
	JSON: []TestCase{
		{
			Flat: "testdata/fixture/file1.json",
			Deep: "testdata/fixture/deep1.json",
		},
		{
			Flat: "testdata/fixture/file2.json",
			Deep: "testdata/fixture/deep2.json",
		},
	},
	Yaml: []TestCase{
		{
			Flat: "testdata/fixture/file1.yaml",
			Deep: "testdata/fixture/deep1.yaml",
		},
		{
			Flat: "testdata/fixture/file2.yaml",
			Deep: "testdata/fixture/deep2.yaml",
		},
	},
}

var formats = struct {
	Stylish string
	Plain   string
	JSON    string
}{
	Stylish: "stylish",
	Plain:   "plain",
	JSON:    "json",
}

func TestGenDiff(t *testing.T) {
	testData := []struct {
		name   string
		path1  string
		path2  string
		want   string
		format string
	}{
		//--------- Stylish -------------
		{
			name:   "flatJSON[Stylish]",
			path1:  filePath.JSON[0].Flat,
			path2:  filePath.JSON[1].Flat,
			want:   stylishWant.Flat,
			format: formats.Stylish,
		},
		{
			name:   "deepJSON[Stylish]",
			path1:  filePath.JSON[0].Deep,
			path2:  filePath.JSON[1].Deep,
			want:   stylishWant.Deep,
			format: formats.Stylish,
		},
		{
			name:   "flatYAML[Stylish]",
			path1:  filePath.Yaml[0].Flat,
			path2:  filePath.Yaml[1].Flat,
			want:   stylishWant.Flat,
			format: formats.Stylish,
		},
		{
			name:   "deepYAML[Stylish]",
			path1:  filePath.Yaml[0].Deep,
			path2:  filePath.Yaml[1].Deep,
			want:   stylishWant.Deep,
			format: formats.Stylish,
		},
		//--------- Plain -------------
		{
			name:   "flatJSON[Plain]",
			path1:  filePath.JSON[0].Flat,
			path2:  filePath.JSON[1].Flat,
			want:   plainWant.Flat,
			format: formats.Plain,
		},
		{
			name:   "deepJSON[Plain]",
			path1:  filePath.JSON[0].Deep,
			path2:  filePath.JSON[1].Deep,
			want:   plainWant.Deep,
			format: formats.Plain,
		},
		{
			name:   "flatYAML[Plain]",
			path1:  filePath.Yaml[0].Flat,
			path2:  filePath.Yaml[1].Flat,
			want:   plainWant.Flat,
			format: formats.Plain,
		},
		{
			name:   "deepYAML[Plain]",
			path1:  filePath.Yaml[0].Deep,
			path2:  filePath.Yaml[1].Deep,
			want:   plainWant.Deep,
			format: formats.Plain,
		},
		//--------- JSON -------------
		{
			name:   "flatJSON[JSON]",
			path1:  filePath.JSON[0].Flat,
			path2:  filePath.JSON[1].Flat,
			want:   jsonWant.Flat,
			format: formats.JSON,
		},
		{
			name:   "deepJSON[JSON]",
			path1:  filePath.JSON[0].Deep,
			path2:  filePath.JSON[1].Deep,
			want:   jsonWant.Deep,
			format: formats.JSON,
		},
		{
			name:   "flatYAML[JSON]",
			path1:  filePath.Yaml[0].Flat,
			path2:  filePath.Yaml[1].Flat,
			want:   jsonWant.Flat,
			format: formats.JSON,
		},
		{
			name:   "deepYAML[JSON]",
			path1:  filePath.Yaml[0].Deep,
			path2:  filePath.Yaml[1].Deep,
			want:   jsonWant.Deep,
			format: formats.JSON,
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

			if tc.format == formats.JSON {
				var gotMap, wantMap any
				require.NoError(t, json.Unmarshal([]byte(got), &gotMap))
				require.NoError(t, json.Unmarshal([]byte(tc.want), &wantMap))
				assert.Equal(t, wantMap, gotMap)
			} else {
				assert.Equal(t, tc.want, got)
			}

		})
	}
}
