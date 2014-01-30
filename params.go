package pangaea

import (
	"github.com/stretchr/objx"
)

// parseParams processes the parameter string and returns a
// map of the keys and values.
func parseParams(s string) (objx.Map, error) {
	return objx.FromURLQuery(s)
}
