package sfmt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zbindenren/sfmt"
)

func TestParseFormat(t *testing.T) {
	var tt = []struct {
		input    string
		expected sfmt.Format
	}{
		{
			"json",
			sfmt.JSON,
		},
		{
			"yaml",
			sfmt.YAML,
		},
		{
			"csv",
			sfmt.CSV,
		},
		{
			"table",
			sfmt.Table,
		},
		{
			"not-exist",
			sfmt.Unknown,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.input, func(t *testing.T) {
			f := sfmt.ParseFormat(tc.input)
			assert.Equal(t, tc.expected, f)
		})
	}
}
