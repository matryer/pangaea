package pangaea

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"os"
)

// builtInMethods represents the built in $$ methods available to
// pangaea.
var builtInMethods Methods = Methods{
	// $$contentsOf loads the contents of a file.
	//
	//     $$contentsOf(filename)
	//
	&Method{Name: "contentsOf", Func: func(call otto.FunctionCall) otto.Value {

		if filename, err := call.Argument(0).ToString(); err == nil {

			if file, err := os.Open(filename); err == nil {
				defer file.Close()

				if content, err := ioutil.ReadAll(file); err == nil {
					if val, err := otto.ToValue(string(content)); err == nil {
						return val
					} else {
						return stringToValue(err.Error())
					}
				} else {
					return stringToValue(err.Error())
				}

			} else {
				return stringToValue(err.Error())
			}

		} else {
			return stringToValue(err.Error())
		}

		return otto.UndefinedValue()
	}},
}

// stringToValue makes an otto.Value containing the specified string.
func stringToValue(message string) otto.Value {
	val, err := otto.ToValue(message)
	if err != nil {
		return otto.UndefinedValue()
	}
	return val
}
