package assert

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
	return m.expectedValue == value
}

func (m *isMatcher) DescribeMismatch(value interface{}) error {
	return buildError("expected <%v> but got <%v>", m.message, m.expectedValue, value)
}
