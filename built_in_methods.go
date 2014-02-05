package pangaea

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"os"
	"os/exec"
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

	// $$run runs a shell command and returns the result.
	//
	//     $$run(command, args...)
	//
	&Method{Name: "run", Func: func(call otto.FunctionCall) otto.Value {

		if command, err := call.Argument(0).ToString(); err == nil {

			// collect all other arguments
			args := make([]string, len(call.ArgumentList)-1)
			for i := 1; i < len(call.ArgumentList); i++ {
				if argStr, err := call.Argument(i).ToString(); err == nil {
					args[i-1] = argStr
				} else {
					args[i-1] = ""
				}
			}

			cmd := exec.Command(command, args...)
			if output, err := cmd.Output(); err == nil {
				return stringToValue(string(output))
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
