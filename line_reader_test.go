package pangaea

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

var lineReaderTests = []struct {
	source string
	lines  []string
}{
	{
		source: `one
two
three`,
		lines: []string{"one", "two", "three"},
	},
	{
		source: `one`,
		lines:  []string{"one"},
	},
}

func TestNewLineReader(t *testing.T) {

	src := strings.NewReader("One\nTwo\nThree")
	r := NewLineReader(src)

	if assert.NotNil(t, r) {
		assert.Equal(t, r.source, src)
		assert.NotNil(t, r.scanner)
	}

}

func TestLineReader(t *testing.T) {

	for _, test := range lineReaderTests {
		lineReader := NewLineReader(strings.NewReader(test.source))
		for _, line := range test.lines {
			readLine, err := lineReader.ReadLine()
			if assert.NoError(t, err) {
				assert.Equal(t, line, string(readLine))
			}
		}
		// one last call should fail
		_, err := lineReader.ReadLine()
		assert.Equal(t, err, io.EOF, "Final call should return EOF")
	}

}
