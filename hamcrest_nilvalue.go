package assert

import "reflect"

type nilValueMatcher struct {
	message string
}

func NilValue() *nilValueMatcher {
	matcher := new(nilValueMatcher)
	return matcher
}

func (m *nilValueMatcher) Message(message string) Matcher {
	m.message = message
	return m
}

func (m *nilValueMatcher) Matches(value interface{}) bool {
	if value == nil {
		return true
	}
	r := reflect.ValueOf(value)
	if r.Kind() == reflect.String {
		return false
	}
	return r.IsNil()
}

func (m *nilValueMatcher) DescribeMismatch(value interface{}) error {
	return buildError("expected nil but: was <%v>", m.message, value)
}
