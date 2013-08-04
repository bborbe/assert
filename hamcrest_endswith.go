package assert

import "strings"

type endswithMatcher struct {
	expectedValue string
	message       string
}

func Endswith(expectedValue string) *endswithMatcher {
	matcher := new(endswithMatcher)
	matcher.expectedValue = expectedValue
	return matcher
}

func (m *endswithMatcher) Message(message string) Matcher {
	m.message = message
	return m
}

func (m *endswithMatcher) Matches(value interface{}) bool {
	text := value.(string)
	return strings.LastIndex(text, m.expectedValue) == len(text)-len(m.expectedValue)
}

func (m *endswithMatcher) DescribeMismatch(value interface{}) error {
	return buildError("expected <%v> ends with <%v>", m.message, value, m.expectedValue)
}
