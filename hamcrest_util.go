package assert

import (
	"errors"
	"fmt"
)

func buildError(format string, message string, expectedValue interface{}, value interface{}) error {
	if len(message) == 0 {
		return fmt.Errorf(format, expectedValue, value)
	}
	return errors.New(fmt.Sprintf("%s, "+format, message, expectedValue, value))
}
