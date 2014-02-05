package pangaea

import (
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"testing"
)

var buildInMethodTests = []struct {
	code   string
	assert func(*testing.T, otto.Value, string)
}{
	{
		code: `$$contentsOf("test/test.txt")`,
		assert: func(t *testing.T, v otto.Value, msg string) {
			if valStr, err := v.ToString(); assert.NoError(t, err) {
				assert.Equal(t, `This is a test text file.`, valStr, msg)
			}
		},
	},
	{
		code: `$$run("echo", "-n", "Hello Pangaea.")`,
		assert: func(t *testing.T, v otto.Value, msg string) {
			if valStr, err := v.ToString(); assert.NoError(t, err) {
				assert.Equal(t, `Hello Pangaea.`, valStr, msg)
			}
		},
	},
}

func TestBuildInMethodTests(t *testing.T) {

	runtime := otto.New()
	assert.NoError(t, builtInMethods.Bind(runtime))

	for _, test := range buildInMethodTests {

		val, err := runtime.Run(test.code)
		if assert.NoError(t, err, "Expected no error when Codeing: "+test.code) {
			if test.assert == nil {
				assert.Fail(t, "Missing 'assert' func.", "Tests must provide an assert func: %s", test.code)
			} else {
				test.assert(t, val, "Assertions failed for: "+test.code)
			}
		}

	}

}
