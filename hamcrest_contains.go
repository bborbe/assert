package assert

import "strings"

type containsMatcher struct {
	expectedValue string
	message       string
}

func Contains(expectedValue string) *containsMatcher {
	matcher := new(containsMatcher)
	matcher.expectedValue = expectedValue
	return matcher
}

func (m *containsMatcher) Message(message string) Matcher {
	m.message = message
	return m
}

func (m *containsMatcher) Matches(value interface{}) bool {
	text := value.(string)
	return strings.Contains(text, m.expectedValue)
}

func (m *containsMatcher) DescribeMismatch(value interface{}) error {
	return buildError("expected <%v> contains <%v>", m.message, value, m.expectedValue)
}
