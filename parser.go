package pangaea

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/robertkrimen/otto"
	"io"
)

var (
	emptyLine         []byte
	ScriptStartLine   []byte = []byte("<script>")
	ScriptEndLine     []byte = []byte("</script>")
	Linefeed          []byte = []byte("\n")
	InlineScriptStart []byte = []byte("<%=")
	InlineScriptEnd   []byte = []byte("%>")
)

type Parser struct {
	reader *LineReader
	writer io.Writer
	js     *otto.Otto
}

func New(reader io.Reader, writer io.Writer) *Parser {
	return &Parser{reader: NewLineReader(reader), writer: writer, js: otto.New()}
}

func (p *Parser) Parse() error {

	var line []byte
	var err error
	for err == nil {

		// read the line
		line, err = p.reader.ReadLine()

		if err == nil {

			// write the line out to the output
			parsedLine, err := p.parseLine(line)

			if err != nil {
				return err
			}
			if _, err = p.writer.Write(parsedLine); err != nil {
				return err
			}

		}

	}

	return nil
}

func (p *Parser) parseLine(line []byte) ([]byte, error) {

	if bytes.Contains(line, ScriptStartLine) {

		// read all script lines
		var block [][]byte = [][]byte{}
		var err error
		var nextline []byte

		for err == nil {
			nextline, err = p.reader.ReadLine()
			if err == io.EOF {
				// shouldn't get this while looking for end script tag
				return emptyLine, errors.New(fmt.Sprintf("pangaea: Missing end %s tag.", string(ScriptEndLine)))
			} else if err != nil {
				return emptyLine, err
			} else {
				if bytes.Contains(nextline, ScriptEndLine) {
					break
				} else {
					// collect the block
					block = append(block, nextline)
				}
			}
		}

		// execute the script block
		_, runErr := p.js.Run(string(bytes.Join(block, Linefeed)))

		if runErr != nil {
			// TODO: explain the block that caused this
			return nil, runErr
		}

		return emptyLine, nil

	}

	for bytes.Contains(line, InlineScriptStart) && bytes.Contains(line, InlineScriptEnd) {

		var newLine []byte
		is := bytes.Index(line, InlineScriptStart)
		ie := bytes.Index(line, InlineScriptEnd)

		newLine = append(newLine, line[:is]...)

		block := line[(is + len(InlineScriptStart)):ie]
		value, err := p.js.Run(string(block))

		if err != nil {
			return emptyLine, err
		}

		newLine = append(newLine, []byte(value.String())...)
		newLine = append(newLine, line[(ie+len(InlineScriptEnd)):]...)

		line = newLine

	}

	// put the linefeed back on
	line = append(line, Linefeed...)

	return line, nil

}
