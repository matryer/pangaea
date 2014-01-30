package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/pangaea"
	"os"
)

var (
	params = flag.String("params", "", "A URL encoded string containing parameters that are made available to scripts.")
)

func main() {

	var err error
	flag.Parse()

	// make a parser
	parser := pangaea.New(os.Stdin, os.Stdout)

	// any params to set?
	if len(*params) > 0 {
		if err := parser.SetParamsFromURLStr(*params); err != nil {
			assertNoErr(err, "Failed to parse parameters.")
		}
	}

	// parse the input and write the output
	err = parser.Parse()

	// make sure there were no errors
	assertNoErr(err, "")

	// OK
	os.Exit(0)

}

func assertNoErr(e error, msg string) {
	if e != nil {
		fatal(e.Error())
	}
}

func fatal(s string) {
	fmt.Printf("%s\n", s)
	flag.Usage()
	os.Exit(1)
}
