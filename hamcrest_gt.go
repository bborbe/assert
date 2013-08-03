package assert

import (
	"reflect"
)

type gtMatcher struct {
	expectedValue interface{}
	message       string
}

func Gt(expectedValue interface{}) *gtMatcher {
	matcher := new(gtMatcher)
	matcher.expectedValue = expectedValue
	return matcher
}

func (m *gtMatcher) Message(message string) Matcher {
	m.message = message
	return m
}

func (m *gtMatcher) Matches(value interface{}) bool {
	if m.sameType(value) {
		switch value.(type) {
		case int:
			return value.(int) > m.expectedValue.(int)
		}
	}
	return false
}

func (m *gtMatcher) sameType(value interface{}) bool {
	expectedType := reflect.TypeOf(m.expectedValue)
	valueType := reflect.TypeOf(value)
	return expectedType == valueType
}

func (m *gtMatcher) DescribeMismatch(value interface{}) error {
	if m.sameType(value) {
		return buildError("expected <%v> is greater than <%v>", m.message, m.expectedValue, value)
	}
	expectedType := reflect.TypeOf(m.expectedValue)
	valueType := reflect.TypeOf(value)
	return buildError("expected type %v but got %v", m.message, expectedType, valueType)
}
