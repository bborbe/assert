package assert

import "errors"

type notNilValueMatcher struct {
	message string
}

func NotNilValue() *notNilValueMatcher {
	matcher := new(notNilValueMatcher)
	return matcher
}

func (m *notNilValueMatcher) Message(message string) Matcher {
	m.message = message
	return m
}

func (m *notNilValueMatcher) Matches(value interface{}) bool {
	return value != nil
}

func (m *notNilValueMatcher) DescribeMismatch(value interface{}) error {
	return errors.New("Expected: is not nil but: was nil")
}
