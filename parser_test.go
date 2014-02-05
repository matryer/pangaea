package pangaea

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var parserTests = []struct {
	src    string
	out    string
	info   string
	err    string
	params string
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
		src: `<pangaea type="text/javascript">
    function name(){
      return "Pangaea";
    }
</pangaea>`,
		out: ``,
	},
	{
		info: "<%= %> are executed",
		src: `<pangaea type="text/javascript">
    function name(){
      return "Pangaea";
    }
</pangaea>
My name is <%= name() %>, and <%= name() %> is my name.`,
		out: `My name is Pangaea, and Pangaea is my name.
`,
	},
	{
		info: "Missing end script tag",
		src: `<pangaea type="text/javascript">
    function name(){
      return "Pangaea";
    }
My name is <%= name() %>, and <%= name() %> is my name.`,
		out: ``,
		err: "pangaea: Missing end </pangaea> tag.",
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
		info:   "Parameters are available as variables",
		src:    `My name is <%= $$params["name"] %>.`,
		params: "name=Mat",
		out: `My name is Mat.
`,
	},
	{
		info: "Invalid JavaScript yields error",
		src:  `10 + 5 is <%= 10 +  %>!`,
		out:  ``,
		err:  "SyntaxError: Unexpected end of input (line 1)",
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

		if len(test.params) > 0 {
			assert.NoError(t, parser.SetParamsFromURLStr(test.params))
		}

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

func TestSetParamsFromURLStr(t *testing.T) {

	parser := New(nil, nil)

	args, _ := parser.js.Get("$$params")
	assert.True(t, args.IsObject(), "$$params should be an empty object even if no params have been set")

	parser.SetParamsFromURLStr("name=Mat&age=30")

	args, _ = parser.js.Get("$$params")
	if assert.True(t, args.IsObject(), "$$params should be an object") {

		name, _ := args.Object().Get("name")
		age, _ := args.Object().Get("age")

		assert.Equal(t, name.String(), "Mat")
		assert.Equal(t, age.String(), "30")

	}

}
