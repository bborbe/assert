package assert

import "strings"

type startswithMatcher struct {
	expectedValue string
	message       string
}

func Startswith(expectedValue string) *startswithMatcher {
	matcher := new(startswithMatcher)
	matcher.expectedValue = expectedValue
	return matcher
}

func (m *startswithMatcher) Message(message string) Matcher {
	m.message = message
	return m
}

func (m *startswithMatcher) Matches(value interface{}) bool {
	text := value.(string)
	return strings.Index(text, m.expectedValue) == 0
}

func (m *startswithMatcher) DescribeMismatch(value interface{}) error {
	return buildError("expected <%v> starts with <%v>", m.message, value, m.expectedValue)
}
