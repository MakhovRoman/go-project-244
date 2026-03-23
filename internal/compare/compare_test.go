package compare

import (
	"code/internal/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDiff(t *testing.T) {
	wantDiff := DiffMap{
		"follow":  {Status: CodeRemoved},
		"host":    {Status: CodeUnchanged},
		"proxy":   {Status: CodeRemoved},
		"timeout": {Status: CodeChanged},
		"verbose": {Status: CodeAdded},
	}

	file1, file2 := "../../testdata/fixture/file1.json", "../../testdata/fixture/file2.json"

	t.Run("base", func(t *testing.T) {
		t.Parallel()

		parsed1, err := parser.Parse(file1)
		if err != nil {
			t.Error(err)
		}

		parsed2, err := parser.Parse(file2)
		if err != nil {
			t.Error(err)
		}
		gotDiff := BuildDiff(parsed1, parsed2)
		for k, v := range wantDiff {
			assert.Equal(t, v.Status, gotDiff[k].Status)
		}
	})
}
