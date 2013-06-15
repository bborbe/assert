package assert

import (
	"fmt"
	"reflect"
)

func NotNull(message string, value interface{}) error {
	if value == nil {
		return fmt.Errorf("%s, expect not null value", message)
	}
	return nil
}

func Null(message string, value interface{}) error {
	if value != nil {
		return fmt.Errorf("%s, expect null value", message)
	}
	return nil
}

func EqualsString(message string, expected string, value string) error {
	if expected != value {
		return fmt.Errorf("%s, expected '%s' but got '%s'", message, expected, value)
	}
	return nil
}

func EqualsBool(message string, expected bool, value bool) error {
	if expected != value {
		return fmt.Errorf("%s, expected %v but got %v", message, expected, value)
	}
	return nil
}

func EqualsInt(message string, expected int, value int) error {
	if expected != value {
		return fmt.Errorf("%s, expected %d but got %d", message, expected, value)
	}
	return nil
}

func True(message string, value bool) error {
	if value != true {
		return fmt.Errorf("%s, expected true but got false", message)
	}
	return nil
}

func False(message string, value bool) error {
	if value != false {
		return fmt.Errorf("%s, expected false but got true", message)
	}
	return nil
}

func Implements(message string, expected interface{}, value interface{}) error {
	expectedType := reflect.TypeOf(expected).Elem()
	valueType := reflect.TypeOf(value)
	if !valueType.Implements(expectedType) {
		return fmt.Errorf("%s, expected type '%s' but got '%s'", message, expectedType.Name(), valueType.Elem().Name())
	}
	return nil
}

func Length(message string, expectedLength int, arrayLength int) error {
	if expectedLength != arrayLength {
		return fmt.Errorf("%s, expected array length %d but got %d", message, expectedLength, arrayLength)
	}
	return nil
}
