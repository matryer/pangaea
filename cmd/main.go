package main

import (
	"fmt"
	"github.com/stretchr/pangaea"
	"os"
)

func main() {

	// make a parser
	parser := pangaea.New(os.Stdin, os.Stdout)
	err := parser.Parse()
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
	os.Exit(1)
}
