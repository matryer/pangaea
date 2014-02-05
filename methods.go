package pangaea

import (
	"github.com/robertkrimen/otto"
)

// MethodFunc represents the Go func that does the work for
// a Method.
type MethodFunc func(call otto.FunctionCall) otto.Value

// Method represents a Pangaea method that is available to
// pangaea scripts.
type Method struct {
	// Name is the name of the method.
	Name string
	// Func is the MethodFunc that will be executed if this method
	// is called in the script.
	Func MethodFunc
}

// Bind binds the method to the specified runtime.
func (m *Method) Bind(runtime *otto.Otto) error {
	castFunc := (func(call otto.FunctionCall) otto.Value)(m.Func)
	if err := runtime.Set("$$"+m.Name, castFunc); err != nil {
		return err
	}
	return nil
}

// Methods represents a slice of *Method objects.
type Methods []*Method

// Bind binds all methods in ths slice to the specified runtime.
func (m Methods) Bind(runtime *otto.Otto) error {

	for _, method := range m {
		if err := method.Bind(runtime); err != nil {
			return err
		}
	}

	// all ok
	return nil
}
