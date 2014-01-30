package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/pangaea"
	"log"
	"os"
	"path/filepath"
)

/*

  pangaea - command line

  pangaea ./

*/

// flags
var (
	source = flag.String("s", "", "File to process")
	output = flag.String("o", "pangaea-out.txt", "File to output to")
)

func main() {

	flag.Parse()

	// check inputs
	wd, err := os.Getwd()
	assertNoErr(err, "Failed to get working directory.")
	sourceFullPath := filepath.Join(wd, *source)
	log.Printf("%s", sourceFullPath)
	sourcePath, err := os.Stat(sourceFullPath)
	assertNoErr(err, "Invalid source.")

	if sourcePath.IsDir() {
		fatal("Source must be a file, cannot be a directory.")
	}

	// open the source file
	sourceFile, err := os.Open(sourceFullPath)
	assertNoErr(err, "Couldn't read source.")
	defer sourceFile.Close()

	// open the output file
	outputFile, err := os.Create(*output)
	assertNoErr(err, "Couldn't open output file for writing.")
	defer outputFile.Close()

	// make a parser
	parser := pangaea.New(sourceFile, outputFile)
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
