package pangaea

import (
	"bufio"
	"io"
)

// LineReader reads from an io.Reader one line at a time.
type LineReader struct {
	source  io.Reader
	scanner *bufio.Scanner
}

// NewLineReader creates a new LineReader capable of reading lines from
// the specified source io.Reader.
func NewLineReader(source io.Reader) *LineReader {
	return &LineReader{
		source:  source,
		scanner: bufio.NewScanner(source),
	}
}

// ReadLine reads the next line from the io.Reader that was specified
// in NewLineReader.
func (r *LineReader) ReadLine() ([]byte, error) {
	if !r.scanner.Scan() {
		return nil, io.EOF
	}
	return []byte(r.scanner.Text()), nil
}
