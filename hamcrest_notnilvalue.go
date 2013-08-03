package assert

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
	return buildError("expected not nil value", m.message)
}
