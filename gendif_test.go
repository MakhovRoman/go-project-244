package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		file1 := "testdata/fixture/file1.json"
		file2 := "testdata/fixture/file2.json"

		got, err := GenDiff(file1, file2)
		if err != nil {
			t.Fatal(err)
		}

		want := "- follow: false\n" +
			"  host: hexlet.io\n" +
			"- proxy: 123.234.53.22\n" +
			"- timeout: 50\n" +
			"+ timeout: 20\n" +
			"+ verbose: true\n"

		assert.Equal(t, want, got)
	})
}
