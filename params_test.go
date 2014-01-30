package pangaea

import (
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"testing"
)

var parseParamsTests = []struct {
	in  string
	out objx.Map
	err string
}{
	{
		in:  `name=Mat&age=30`,
		out: objx.MSI("name", "Mat", "age", "30"),
	},
}

func TestParseParams(t *testing.T) {

	for _, test := range parseParamsTests {

		params, err := parseParams(test.in)

		if len(test.err) > 0 {
			assert.Equal(t, test.err, err.Error())
		} else {
			for k, v := range test.out {
				assert.Equal(t, v, params[k], "Key %s should be %#v, but was %#v", k, v, params[k])
			}
		}

	}

}
