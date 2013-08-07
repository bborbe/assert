package assert

import "reflect"

type isMatcher struct {
	expectedValue interface{}
	message       string
}

func Is(expectedValue interface{}) *isMatcher {
	matcher := new(isMatcher)
	matcher.expectedValue = expectedValue
	return matcher
}

func (m *isMatcher) Message(message string) Matcher {
	m.message = message
	return m
}

func (m *isMatcher) Matches(value interface{}) bool {
	if sameType(value, m.expectedValue) {
		return m.expectedValue == value
	}
	return false
}

func (m *isMatcher) DescribeMismatch(value interface{}) error {
	if sameType(value, m.expectedValue) {
		return buildError("expected <%v> but got <%v>", m.message, m.expectedValue, value)
	}
	expectedType := reflect.TypeOf(m.expectedValue)
	valueType := reflect.TypeOf(value)
	return buildError("expected type %v but got %v", m.message, expectedType, valueType)

}
