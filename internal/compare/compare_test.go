package compare

import (
	"code/internal/parsers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDiff(t *testing.T) {
	want := "- follow: false\n" +
		"  host: hexlet.io\n" +
		"- proxy: 123.234.53.22\n" +
		"- timeout: 50\n" +
		"+ timeout: 20\n" +
		"+ verbose: true\n"

	file1, file2 := "../../testdata/fixture/file1.json", "../../testdata/fixture/file2.json"

	t.Run("base", func(t *testing.T) {
		t.Parallel()

		parsed1, err := parsers.Parse(file1)
		if err != nil {
			t.Error(err)
		}

		parsed2, err := parsers.Parse(file2)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, want, BuildDiff(parsed1, parsed2))
	})
}
