package pangaea

import (
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"testing"
)

var buildInMethodTests = []struct {
	Call   string
	Assert func(*testing.T, otto.Value, string)
}{
	{
		Call: `$$contentsOf("test/test.txt")`,
		Assert: func(t *testing.T, v otto.Value, msg string) {
			if valStr, err := v.ToString(); assert.NoError(t, err) {
				assert.Equal(t, `This is a test text file.`, valStr, msg)
			}
		},
	},
}

func TestBuildInMethodTests(t *testing.T) {

	runtime := otto.New()
	assert.NoError(t, builtInMethods.Bind(runtime))

	for _, test := range buildInMethodTests {

		val, err := runtime.Run(test.Call)
		if assert.NoError(t, err, "Expected no error when calling: "+test.Call) {
			test.Assert(t, val, "Assertions failed for: "+test.Call)
		}

	}

}
