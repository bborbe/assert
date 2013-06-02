package assert

import (
	"errors"
	"fmt"
	"reflect"
)

func NotNull(message string, value interface{}) error {
	if value == nil {
		return errors.New(fmt.Sprintf("%s, expect not null value", message))
	}
	return nil
}

func Null(message string, value interface{}) error {
	if value != nil {
		return errors.New(fmt.Sprintf("%s, expect null value", message))
	}
	return nil
}

func EqualsString(message string, expected string, value string) error {
	if expected != value {
		return errors.New(fmt.Sprintf("%s, expected '%s' but got '%s'", message, expected, value))
	}
	return nil
}

func EqualsInt(message string, expected int, value int) error {
	if expected != value {
		return errors.New(fmt.Sprintf("%s, expected %d but got %d", message, expected, value))
	}
	return nil
}

func True(message string, value bool) error {
	if value != true {
		return errors.New(fmt.Sprintf("%s, expected true but got false", message))
	}
	return nil
}

func False(message string, value bool) error {
	if value != false {
		return errors.New(fmt.Sprintf("%s, expected false but got true", message))
	}
	return nil
}

func Implements(message string, expected interface{}, value interface{}) error {
	expectedType := reflect.TypeOf(expected).Elem()
	valueType := reflect.TypeOf(value)
	if !valueType.Implements(expectedType) {
		return errors.New(fmt.Sprintf("%s, expected type '%s' but got '%s'", message, expectedType.Name(), valueType.Elem().Name()))
	}
	return nil
}

func Length(message string, expectedLength int, arrayLength int) error {
	if expectedLength != arrayLength {
		return errors.New(fmt.Sprintf("%s, expected array length %d but got %d", message, expectedLength, arrayLength))
	}
	return nil
}
