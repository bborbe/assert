package assert

import (
	"errors"
	"fmt"
)

func buildError(format string, message string, args ...interface{}) error {
	if len(message) == 0 {
		return fmt.Errorf(format, args...)
	}
	var a []interface{}
	b := append(a, message)
	c := append(b, args...)
	return errors.New(fmt.Sprintf("%s, "+format, c...))
}
