package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/pangaea"
	"os"
)

func main() {

	flag.Parse()

	// make a parser
	parser := pangaea.New(os.Stdin, os.Stdout)
	parser.Parse()

	// OK
	os.Exit(0)

}

func assertNoErr(e error, msg string) {
	if e != nil {
		if pathErr, ok := e.(*os.PathError); ok {
			fatal(fmt.Sprintf("%s: %s %s", pathErr.Err.Error(), pathErr.Path, msg))
		} else {
			fatal(e.Error())
		}
	}
}

func fatal(s string) {
	fmt.Printf("%s\n", s)
	flag.Usage()
	os.Exit(1)
}
