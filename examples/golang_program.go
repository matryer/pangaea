package main

// Run this from within cmd directory:
//
//   cat ../examples/golang_program.go | ./pangaea -params="code=10*5"
//

import (
  "log"
)

func main() {

  log.Printf("Running code: %s", `<%= $params["code"] %>`)
  log.Printf("%v", <%= $params["code"] %>)

}
