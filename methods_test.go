package pangaea

import (
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMethodBind(t *testing.T) {

	method := &Method{Name: "sayHello", Func: func(call otto.FunctionCall) otto.Value {
		v, _ := otto.ToValue("Hello Pangaea")
		return v
	}}

	runtime := otto.New()
	value, err := runtime.Run("$$sayHello()")
	assert.Error(t, err, "Should be an error calling sayHello before the methods have been bound.")

	assert.NoError(t, method.Bind(runtime))

	value, err = runtime.Run("$$sayHello()")
	if assert.NoError(t, err) {
		if valueStr, err := value.ToString(); assert.NoError(t, err) {
			assert.Equal(t, valueStr, "Hello Pangaea")
		}
	}

}

func TestMethodsBind(t *testing.T) {

	methods := Methods{&Method{Name: "sayHello", Func: func(call otto.FunctionCall) otto.Value {
		v, _ := otto.ToValue("Hello Pangaea")
		return v
	}}}

	runtime := otto.New()
	value, err := runtime.Run("$$sayHello()")
	assert.Error(t, err, "Should be an error calling sayHello before the methods have been bound.")

	assert.NoError(t, methods.Bind(runtime))

	value, err = runtime.Run("$$sayHello()")
	if assert.NoError(t, err) {
		if valueStr, err := value.ToString(); assert.NoError(t, err) {
			assert.Equal(t, valueStr, "Hello Pangaea")
		}
	}

}
