package pangaea

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var parserTests = []struct {
	src  string
	out  string
	info string
	err  string
}{
	{
		info: "Nothing begets nothing",
		src:  ``,
		out:  ``,
	},
	{
		info: "Literal content is copied",
		src:  `Hello world`,
		out: `Hello world
`,
	},
	{
		info: "code blocks are removed",
		src: `<script>
    function name(){
      return "Pangaea";
    }
</script>`,
		out: ``,
	},
	{
		info: "<%= %> are executed",
		src: `<script>
    function name(){
      return "Pangaea";
    }
</script>
My name is <%= name() %>, and <%= name() %> is my name.`,
		out: `My name is Pangaea, and Pangaea is my name.
`,
	},
	{
		info: "Missing end script tag",
		src: `<script>
    function name(){
      return "Pangaea";
    }
My name is <%= name() %>, and <%= name() %> is my name.`,
		out: ``,
		err: "pangaea: Missing end </script> tag.",
	},
	{
		info: "Line feeds are respected",
		src: `one
two
three`,
		out: `one
two
three
`,
	},
	{
		info: "Other JavaScript works too",
		src:  `10 + 5 is <%= 10 + 5 %>!`,
		out: `10 + 5 is 15!
`,
	},
	{
		info: "Empty line feeds are respected",
		src: `





`,
		out: `





`,
	},
}

func TestParserTests(t *testing.T) {

	for _, test := range parserTests {

		var buf bytes.Buffer

		parser := New(strings.NewReader(test.src), &buf)
		err := parser.Parse()
		if len(test.err) > 0 {
			assert.Equal(t, err.Error(), test.err)
		} else {
			assert.NoError(t, err)
		}

		assert.Equal(t, buf.String(), test.out, test.info)

	}

}

func TestNew(t *testing.T) {

	var buf bytes.Buffer
	reader := strings.NewReader("")
	writer := &buf
	parser := New(reader, writer)

	if assert.NotNil(t, parser) {
		assert.Equal(t, reader, parser.reader.source, "Reader")
		assert.Equal(t, writer, parser.writer, "Writer")
		assert.NotNil(t, parser.js, "js")
	}

}
