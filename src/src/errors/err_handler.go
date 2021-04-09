package error

import (
	"fmt"

	object "github.com/Hydrogen/src/evaluator/objects"
)

func NewError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
